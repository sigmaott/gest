package main

import (
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"query-builder/mongo"
)

type Test struct {
	RecordType string `json:"name"  bson:"recordType,omitempty" sortable:"true" filterable:"true"`
	Items      []Item `json:"name"  bson:"Items" sortable:"true" filterable:"true"`
}
type Item struct {
	Address string `json:"address" bson:"address "sortable:"true" filterable:"true"`
	Name    string `json:"name" bson:"name" sortable:"true"`
}

func main() {
	s := `{
	 "condition": "and",
	 "rules": [{
	   "field": "RecordType",
	   "operator": "=",
	   "value": "Item"
	 }]
	}`
	var result bson.M
	p := mongo.NewQueryMongoBuilderParser[Test]()
	err := p.Parser(s, &result)
	log.Print(result, err)
	//getTagFromStruct("Items.Name", Test{}, "filterable")

}
