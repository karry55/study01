package main

// import (
// 	"fmt"
// 	"reflect"
// )

// type Student struct {
// 	Name string
// 	Age  int `json:"age"`
// }

// func main() {
// 	s := Student{
// 		Name: "枫枫知道",
// 		Age:  24,
// 	}
// 	t := reflect.TypeOf(s)
// 	v := reflect.ValueOf(s)

// 	for i := 0; i < t.NumField(); i++ {
// 		field := t.Field(i)
// 		jsonField := field.Tag.Get("json")
// 		if jsonField == "" {
// 			// 说明json的tag是空的
// 			jsonField = field.Name
// 		}
// 		fmt.Printf("Name: %s, type: %s, json: %s, value: %v\n", field.Name, field.Type, jsonField, v.Field(i))
// 	}
// }
