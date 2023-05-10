package mongo

import (
	"encoding/json"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"query-builder"
	"strings"
	"time"
)

var conditions = map[string]string{"and": "$and", "or": "$or"}
var operators = map[string]string{
	"=":        "$eq",
	"!=":       "$ne",
	"<":        "$lt",
	"<=":       "$lte",
	">":        "$gt",
	">=":       "$gte",
	"in":       "$in",
	"not in":   "$nin",
	"contains": "$regex",
}

type QueryParser[T any] struct {
}

func (q QueryParser[T]) validate(object any, field string) (ok bool, error error) {
	okStr, err := query_builder.GetTagFromStruct(field, object, "filterable")
	if err != nil {
		return false, query_builder.NewValidateError(error.Error())
	}
	return okStr == "true", nil
}
func (q QueryParser[T]) mapRule(rule *query_builder.Rule) (query bson.M, err error) {

	field := rule.Field
	model := *new(T)
	bsonValue, err := query_builder.GetTagFromStruct(field, model, "bson")

	ok, err := q.validate(model, field)

	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, fmt.Errorf("can't fillter %s", field)
	}
	var value any
	okStr, err := query_builder.GetTypeFromStruct(field, *new(T))
	if err != nil {
		return nil, err
	}
	if okStr == "*time.Time" || okStr == "time.Time" {
		value, err = time.Parse(time.RFC3339, rule.Value.(string))
		if err != nil {
			return nil, err
		}
	} else {
		value = rule.Value
	}

	operator, ok := operators[*rule.Operator]
	if !ok {
		operator = "$eq"
	}
	//var query bson.M
	field = strings.TrimSpace(strings.Split(bsonValue.(string), ",")[0])
	if operator == "$regex" {
		query = bson.M{
			field: bson.M{operator: value, "$options": "i"},
		}
	} else {
		query = bson.M{
			field: bson.M{operator: value},
		}
	}
	return query, nil
}
func (q QueryParser[T]) mapRuleSet(ruleSet *query_builder.Rule) (query bson.M, err error) {
	if len(ruleSet.Rules) == 0 {
		return bson.M{}, nil
	}
	var queries []bson.M
	for _, rule := range ruleSet.Rules {
		if rule.Operator != nil {
			result, err := q.mapRule(rule)
			if err != nil {
				return nil, err

			}
			queries = append(queries, result)
		} else {

			if result, err := q.mapRuleSet(rule); err != nil {
				return nil, err
			} else {
				queries = append(queries, result)
			}

		}
	}
	if len(queries) == 0 {
		return nil, errors.New("query is empty")
	}
	query = bson.M{
		conditions[ruleSet.Condition]: queries,
	}
	return query, nil
}
func (q QueryParser[T]) Parser(queryString string, query any) (err error) {
	p := new(QueryParser[T])
	rule := new(query_builder.Rule)
	if err := json.Unmarshal([]byte(queryString), rule); err != nil {
		return err
	}
	queryParser, err := p.mapRuleSet(rule)

	if err != nil {
		return err
	}
	*query.(*bson.M) = queryParser
	return err
}

func NewQueryMongoBuilderParser[T any]() query_builder.IQueryBuilder {

	return &QueryParser[T]{}

}
