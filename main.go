package main

import (
	"fmt"
	"reflect"
)

type MyStruct struct {
	Field1 string
	Field2 int
}

func (s MyStruct) Function1() {
	fmt.Println("Function1 called")
}

func (s MyStruct) Function2() {
	fmt.Println("Function2 called")
}
func main() {
	myStruct := MyStruct{
		Field1: "Hello",
		Field2: 42,
	}

	// Get the reflect.Value of the struct
	value := reflect.ValueOf(myStruct)

	// Iterate over the struct's methods
	for i := 0; i < value.NumMethod(); i++ {
		// Get the method value
		methodValue := value.Method(i)

		// Call the method dynamically
		methodValue.Call(nil)
	}
}
