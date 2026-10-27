package main

// import (
// 	"fmt"
// 	"reflect"
// 	"strings"
// )

// type Student struct {
// 	Name1 string `big:"-"`
// 	Name2 string `big:"-"`
// }

// func main() {
// 	s := Student{
// 		Name1: "lkl",
// 		Name2: "clx",
// 	}
// 	t := reflect.TypeOf(s)
// 	v := reflect.ValueOf(&s).Elem()

// 	for i := 0; i < t.NumField(); i++ {
// 		field := t.Field(i)
// 		bigField := field.Tag.Get("big")
// 		if field.Type.Kind() != reflect.String {
// 			continue
// 		}

// 		if bigField == "" {
// 			continue
// 		}

// 		valuefiled := v.Field(i)
// 		valuefiled.SetString(strings.ToTitle(valuefiled.String()))
// 	}
// 	fmt.Println(s)
// }
