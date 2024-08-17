package test

import (
	"reflect"
	"testing"

	repository "github.com/sigmaott/gest/package/technique/database/base"
	pagination "github.com/sigmaott/gest/package/technique/database/pagination"
	"go.mongodb.org/mongo-driver/bson"
)

// Sample struct to use for testing purposes
type SampleStruct struct {
	Name  string `bson:"name" sortable:"true" filterable:"true"`
	Age   int    `bson:"age" sortable:"true" filterable:"true"`
	Email string `bson:"email" filterable:"true"`
}

func TestMongoParserQuery(t *testing.T) {
	tests := []struct {
		name         string
		query        map[string][]string
		wantFilter   bson.M
		wantSort     map[string]string
		wantPaginate *repository.Paginate
		wantErr      bool
	}{
		{
			name: "Simple filter query",
			query: map[string][]string{
				"filter.name": {"$eq:John"},
			},
			wantFilter: bson.M{"name": []bson.M{bson.M{"$eq": bson.M{
				"$eq": "John",
			}}}},
			wantSort:     map[string]string{},
			wantPaginate: &repository.Paginate{},
			wantErr:      false,
		},
		{
			name: "Complex filter query with $in operator",
			query: map[string][]string{
				"filter.age": {`$in:25,30,35`},
			},
			wantFilter:   bson.M{"age": []bson.M{bson.M{"$in": []string{"25", "30", "35"}}}},
			wantSort:     map[string]string{},
			wantPaginate: &repository.Paginate{},
			wantErr:      false,
		},
		{
			name: "Sort query",
			query: map[string][]string{
				"sort": {"name:asc", "age:desc"},
			},
			wantFilter:   bson.M{},
			wantSort:     map[string]string{"name": "asc", "age": "desc"},
			wantPaginate: &repository.Paginate{},
			wantErr:      false,
		},
		{
			name: "Pagination query",
			query: map[string][]string{
				"perPage": {"10"},
				"page":    {"2"},
			},
			wantFilter:   bson.M{},
			wantSort:     map[string]string{},
			wantPaginate: &repository.Paginate{Limit: 10, Offset: 10},
			wantErr:      false,
		},
		{
			name: "Invalid sort query",
			query: map[string][]string{
				"sort": {"invalidsort:wrong"},
			},
			wantFilter: nil,
			wantSort:   nil,
			wantErr:    true,
		},
		{
			name: "Simple q query",
			query: map[string][]string{
				"q": {`{"$or":[{"name":{"$eq":"John"}}]}`},
			},
			wantFilter: bson.M{
				"$or": []bson.M{
					{"name": bson.M{"$eq": "John"}},
					// {"age": bson.M{"$gte": "30"}},
				},
			},
			wantSort:     map[string]string{},
			wantPaginate: &repository.Paginate{},
			wantErr:      false,
		},
		// {
		// 	name: "Empty q query",
		// 	query: map[string][]string{
		// 		"q": {`[]`},
		// 	},
		// 	wantFilter:   bson.M{},
		// 	wantSort:     map[string]string{},
		// 	wantPaginate: &base.Paginate{},
		// 	wantErr:      false,
		// },
		// {
		// 	name: "Invalid q query",
		// 	query: map[string][]string{
		// 		"q": {`invalid json`},
		// 	},
		// 	wantFilter:   bson.M{},
		// 	wantSort:     map[string]string{},
		// 	wantPaginate: &base.Paginate{},
		// 	wantErr:      true,
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("Running test: %s", tt.name)
			t.Logf("Query params: %v", tt.query)

			gotFilter, gotSort, gotPaginate, err := pagination.MongoParserQuery[SampleStruct](tt.query)

			t.Logf("Resulting Filter: %v", gotFilter)
			t.Logf("Resulting Sort: %v", gotSort)
			t.Logf("Resulting Paginate: %v", gotPaginate)

			if (err != nil) != tt.wantErr {
				t.Errorf("MongoParserQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !bsonEqual(gotFilter, tt.wantFilter) {
				t.Errorf("MongoParserQuery() gotFilter = %v, want %v", gotFilter, tt.wantFilter)
			}

			if !reflect.DeepEqual(gotSort, tt.wantSort) {
				t.Errorf("MongoParserQuery() gotSort = %v, want %v", gotSort, tt.wantSort)
			}

			if !reflect.DeepEqual(gotPaginate, tt.wantPaginate) {
				t.Errorf("MongoParserQuery() gotPaginate = %v, want %v", gotPaginate, tt.wantPaginate)
			}
		})
	}
}
func bsonEqual(a, b bson.M) bool {
	aJSON, _ := bson.MarshalExtJSON(a, true, true)
	bJSON, _ := bson.MarshalExtJSON(b, true, true)
	return string(aJSON) == string(bJSON)
}
