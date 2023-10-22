package query_builder

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"

	"github.com/samber/lo"
	repository "github.com/sigmaott/gest/package/technique/database/base"
	"go.mongodb.org/mongo-driver/bson"
)

var operators = map[string]string{
	"$eq":     "$eq",
	"$ne":     "$ne",
	"$gt":     "$gt",
	"$gte":    "$gte",
	"$lt":     "$lt",
	"$lte":    "$lte",
	"$in":     "$in",
	"$nin":    "$nin",
	"$and":    "$and",
	"$or":     "$or",
	"$not":    "$not",
	"$exists": "$exists",
}

var operatorSorts = map[string]string{
	"asc":  "asc",
	"ASC":  "asc",
	"DESC": "desc",
	"desc": "desc",
}

const FILTER = "filter"

func MongoParserQuery[T any](query map[string][]string) (bson.M, map[string]string, *repository.Paginate, error) {

	queryDb := map[string][]string{}

	for key, value := range query {
		if strings.HasPrefix(key, FILTER) {
			queryDb[strings.Replace(key, fmt.Sprintf("%s.", FILTER), "", 1)] = value
		}

	}
	sort := map[string]string{}
	objectModel := *new(T)
	configValue := reflect.ValueOf(objectModel)

	// sort query
	if val, ok := query["sort"]; ok {
		//log.Print(val)
		for _, item := range val {

			key, operator, err := parseSortExpression(item)

			pathStruct, err := getPathByTag(key, "bson", configValue, "")
			if err != nil {
				return nil, nil, nil, err
			}
			err = validate(objectModel, pathStruct, "sortable", key)
			if err != nil {
				return nil, nil, nil, err
			}
			if err != nil {
				return nil, nil, nil, err
			}
			sort[key] = operator

		}

	}

	// query prefix filter
	filter := bson.M{}
	for key, val := range queryDb {

		pathStruct, err := getPathByTag(key, "bson", configValue, "")
		if err != nil {
			return nil, nil, nil, err
		}

		err = validate(objectModel, pathStruct, "filterable", key)
		if err != nil {
			return nil, nil, nil, err
		}

		if len(val) == 1 {
			filter[key] = val[0]
		} else if len(val) > 0 {
			queryInFields := lo.Map(val, func(item string, index int) bson.M {
				filter, err := parseFilterExpression(item)
				if err != nil {
					log.Print(err)
					return nil
				}
				return filter
			})
			filter[key] = queryInFields
		}

	}
	// query q
	if q, ok := query["q"]; ok && len(q) > 0 {

		qQuery := map[string][]map[string]any{}
		err := json.Unmarshal([]byte(q[0]), &qQuery)
		if err != nil {
			return nil, nil, nil, err

		}
		// TODO mongo inject
		// Iterate through the query data
		for _, v := range qQuery {
			for _, value := range v {
				for key, _ := range value {
					pathStruct, err := getPathByTag(key, "bson", configValue, "")
					if err != nil {
						return nil, nil, nil, err
					}

					err = validate(objectModel, pathStruct, "filterable", key)
					if err != nil {
						return nil, nil, nil, err
					}
				}

			}

		}

		query, err := mapToBsonA(qQuery)

		if err != nil {
			return nil, nil, nil, err
		}

		for k, v := range query {

			filter[k] = v
		}

	}

	paginate, err := parsePaginate(query)
	if err != nil {
		return nil, nil, nil, err
	}

	return filter, sort, paginate, nil

}

func parseFilterExpression(expression string) (filter bson.M, err error) {
	// filter = bson.M{}
	expression = strings.TrimSpace(expression)
	var value any
	var operator string
	for keyOperator, valueOperator := range operators {
		prefix := fmt.Sprintf("%s:", keyOperator)

		if strings.HasPrefix(expression, prefix) {
			operator = valueOperator
			valueStr := strings.TrimPrefix(expression, prefix)
			if operator == "$in" || operator == "$nin" {

				return bson.M{operator: strings.Split(valueStr, ",")}, nil
			}
			value, err = parseFilterExpression(valueStr)
			if err != nil {
				return nil, err
			}
			break

		}

	}

	//if strings.HasPrefix(expression, "$in:") {
	//	operator = "$in"
	//	value = strings.TrimPrefix(expression, "$in:")
	//}

	// Create the filter object with the specified operator
	if operator == "" {
		return bson.M{"$eq": expression}, nil
	}

	return bson.M{operator: value}, nil

}

func parseSortExpression(expression string) (key string, value string, err error) {

	expression = strings.TrimSpace(expression)
	var operator string
	for keyOperator, valueOperator := range operatorSorts {
		suffix := fmt.Sprintf(":%s", keyOperator)

		if strings.HasSuffix(expression, suffix) {
			operator = valueOperator
			valueStr := strings.TrimSuffix(expression, suffix)
			return valueStr, operator, nil

		}
		suffixTwo := fmt.Sprintf("|%s", keyOperator)

		if strings.HasSuffix(expression, suffixTwo) {
			operator = valueOperator
			valueStr := strings.TrimSuffix(expression, suffixTwo)
			return valueStr, operator, nil

		}

	}
	return "", "", NewValidateError(fmt.Errorf(fmt.Sprintf("sort %s is invaldate", expression)))
}

func parsePaginate(query map[string][]string) (*repository.Paginate, error) {

	paginate := new(repository.Paginate)
	perPageQueries, ok := query["perPage"]
	if ok && len(perPageQueries) > 0 {
		perPageQueryStr := perPageQueries[0]
		perPage, err := strconv.ParseInt(perPageQueryStr, 10, 0)
		if err != nil {
			return nil, err
		}
		paginate.Limit = perPage

	}
	pageQueries, ok := query["page"]
	if ok && len(pageQueries) > 0 {
		pageQueryStr := pageQueries[0]
		page, err := strconv.ParseInt(pageQueryStr, 10, 0)
		if err != nil {
			return nil, err
		}
		if page <= 0 {
			page = 1
		}
		paginate.Offset = (page - 1) * paginate.Limit

	}
	return paginate, nil

}

func mapToBsonA(data map[string][]map[string]any) (bson.M, error) {
	dataBytes, err := bson.Marshal(data)
	if err != nil {
		return nil, err
	}

	var bsonDoc bson.M
	err = bson.Unmarshal(dataBytes, &bsonDoc)
	if err != nil {
		return nil, err
	}

	return bsonDoc, nil
}
