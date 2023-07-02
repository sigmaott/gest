package query_builder

import (
	"fmt"
	"reflect"
	"strings"
)

func getPathByTag(tagValuePath string, tagName string, v reflect.Value, path string) (string, error) {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return "", fmt.Errorf("value is not a struct")
	}

	typ := v.Type()

	// Split the tagValuePath into its parts
	parts := strings.Split(tagValuePath, ".")
	if len(parts) == 0 {
		return "", fmt.Errorf("invalid path")
	}

	// Find the first field that matches the first part of the path
	fieldIndex := -1
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := typ.Field(i)

		if fieldType.Anonymous && field.Kind() == reflect.Struct {

			p, err := getPathByTag(tagName, tagValuePath, field, path)
			if err == nil {
				if path != "" {
					path = fieldType.Name + "." + p
				} else {
					path = p
				}
				return path, nil
			}
		}
		if fieldType.Tag.Get(tagName) == parts[0] {
			fieldIndex = i
			break
		}
	}

	if fieldIndex == -1 {
		return "", fmt.Errorf("field with tag %s:%s not found", tagName, tagValuePath)
	}

	// If this is the last part of the path, return the field name
	if len(parts) == 1 {
		if path != "" {
			path = path + "." + typ.Field(fieldIndex).Name
		} else {
			path = typ.Field(fieldIndex).Name
		}
		return path, nil
	}

	// Otherwise, recurse with the next part of the path
	nextValue := v.Field(fieldIndex)
	nextPath := typ.Field(fieldIndex).Name
	if path != "" {
		nextPath = path + "." + nextPath
	}

	return getPathByTag(strings.Join(parts[1:], "."), tagName, nextValue, nextPath)
}
func getTagFromStruct(keyWithDots string, object any, tagName string) (any, error) {

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

func validate(object any, field string, tagName string, key string) (error error) {
	//getPathByTag(field, "filterable",)
	okStr, err := getTagFromStruct(field, object, tagName)
	if err != nil {
		return NewValidateError(error)
	}
	if okStr == "true" {
		return nil
	}
	return NewValidateError(fmt.Errorf("%s is not %s", key, tagName))
}
