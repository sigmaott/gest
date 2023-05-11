package main

import (
	"fmt"
	"regexp"
)

//type Person struct {
//	Name    string `json:"name"`
//	Age     int    `json:"age"`
//	Address string `json:"address"`
//}
//
//func main() {
//	m := map[string]interface{}{
//		"name":    "Alice",
//		"age":     "33",
//		"address": "123 Main St",
//	}
//
//	var person Person
//	cfg := &mapstructure.DecoderConfig{
//		Metadata: nil,
//		Result:   &person,
//		TagName:  "json",
//	}
//	//err := mapstructure.Decode(m, &person)
//	decoder, err := mapstructure.NewDecoder(cfg)
//	if err != nil {
//		log.Print(err)
//		// Handle the error
//	}
//	log.Print(decoder)
//	err = decoder.Decode(m)
//	log.Print(err)
//	fmt.Println(person.Name)    // "Alice"
//	fmt.Println(person.Age)     // 30
//	fmt.Println(person.Address) // "123 Main St"
//}

func main() {
	errorMessage := fmt.Sprintf("'%s' expected type '%s', got unconvertible type '%s', value: '%v'", "q", "int", "string", "r")

	// Define a regular expression pattern to match the value part of the error message
	pattern := `value:\s*'([\w\s]*)'`
	regex := regexp.MustCompile(pattern)

	// Find the first submatch that matches the pattern in the error message
	match := regex.FindStringSubmatch(errorMessage)

	// Extract the captured value from the match
	if len(match) >= 2 {
		value := match[1]
		fmt.Println(value)
		fmt.Println(match)
	}
}
