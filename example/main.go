//package main
//
//import (
//	"fmt"
//	"github.com/go-playground/locales/en"
//	ut "github.com/go-playground/universal-translator"
//	"github.com/go-playground/validator/v10"
//	en_translations "github.com/go-playground/validator/v10/translations/en"
//	"regexp"
//)

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

package main

import (
	"fmt"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

type Address struct {
	Street string `validate:"required"`
	City   string `validate:"required"`
}

type User struct {
	Name    string  `validate:"required"`
	Age     int     `validate:"required,min=18"`
	Address Address `validate:"required,dive"`
}

func main() {
	// Create a new validator instance
	validate := validator.New()

	// Create a new English translator instance
	enTranslator := en.New()
	uniTranslator := ut.New(enTranslator, enTranslator)
	trans, _ := uniTranslator.GetTranslator("en")

	// Register the English translator for the validator
	en_translations.RegisterDefaultTranslations(validate, trans)

	// Validate a struct
	user := User{
		Name: "Alice",
		Age:  16,
		Address: Address{
			Street: "",
			City:   "Wonderland",
		},
	}
	err := validate.Struct(user)
	if err != nil {
		// Convert the errors to a map with the full path of each field
		errorsMap := make(map[string]string)
		errors := err.(validator.ValidationErrors)
		for _, e := range errors {
			// Get the field name and full path
			//fieldName := e.Field()
			fullPath := e.Namespace()

			// Get the translated error message
			message := e.Translate(trans)

			// Store the error message with the full path
			errorsMap[fullPath] = message
		}

		// Print the errors map
		fmt.Println(errorsMap)
	}
}
