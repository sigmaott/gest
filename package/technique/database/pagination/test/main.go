package main

import (
	"log"

	pagination "github.com/sigmaott/gest/package/technique/database/pagination"
)

type User struct {
	Id   string `bson:"id"`
	Name string `bson:"name" filterable:"true"`
}

func main() {
	query := map[string][]string{}

	query["q"] = []string{
		`{"$and":[{"name": {"$regex": "^test"}}]}`,
	}
	filter, _, _, _ := pagination.MongoParserQuery[User](query)
	log.Print(filter)

}
