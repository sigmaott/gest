package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	// Read the JSON file
	file, err := ioutil.ReadFile("translations.json")
	if err != nil {
		panic(err)
	}

	// Parse the JSON data
	var data []map[string]string
	if err := json.Unmarshal(file, &data); err != nil {
		panic(err)
	}

	// Generate Go constants for each key
	var constants []string
	for _, item := range data {
		key := item["key"]
		constant := fmt.Sprintf("const %s = \"%s\"", strings.Title(key), key)
		constants = append(constants, constant)
	}

	// Print the Go code
	fmt.Println(strings.Join(constants, "\n"))
}
