package query_builder

import (
	"fmt"
	"reflect"
	"strings"
)

type Rule struct {
	Field     string  `json:"field,omitempty"`
	Operator  *string `json:"operator,omitempty"`
	Value     any     `json:"value,omitempty"`
	Condition string  `json:"condition,omitempty"`
	Type      string  `json:"type,omitempty"`
	Rules     []*Rule `json:"rules,omitempty"`
}

func GetTagFromStruct(keyWithDots string, object any, tagName string) (any, error) {

	keySlice := strings.Split(keyWithDots, ".")
	v := reflect.ValueOf(object)
	var tag any

	class, ok := reflect.TypeOf(object).FieldByName(keySlice[0])

	if ok {
		tag = class.Tag.Get(tagName)
	} else {
		return nil, fmt.Errorf("%s isn't path", keyWithDots)
	}
	for _, key := range keySlice[1:] {
		for v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		// we only accept structs
		if v.Kind() != reflect.Struct {
			return nil, fmt.Errorf("%s isn't path", keyWithDots)
		}

		class, ok = reflect.TypeOf(object).FieldByName(key)
		if ok {
			tag = class.Tag.Get(tagName)
		} else {
			return nil, fmt.Errorf("%s isn't path", keyWithDots)
		}

	}

	return tag, nil
}

func GetTypeFromStruct(keyWithDots string, object any) (any, error) {

	keySlice := strings.Split(keyWithDots, ".")
	v := reflect.ValueOf(object)

	class, ok := reflect.TypeOf(object).FieldByName(keySlice[0])

	if !ok {

		return nil, fmt.Errorf("%s isn't path", keyWithDots)
	}
	for _, key := range keySlice[1:] {
		for v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		// we only accept structs
		if v.Kind() != reflect.Struct {
			return nil, fmt.Errorf("%s isn't path", keyWithDots)
		}

		class, ok = reflect.TypeOf(object).FieldByName(key)
		if !ok {
			return nil, fmt.Errorf("%s isn't path", keyWithDots)
		}

	}

	return class.Type.String(), nil
}
