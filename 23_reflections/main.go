package main

import (
	"fmt"
	"reflect"
)

// Cat struct
type Cat struct {
	Name string
	Age  int
}

func example(cat Cat) {
	fmt.Println(reflect.TypeOf(cat))             // main.Cat
	fmt.Println(reflect.TypeOf(cat).Name())      // Cat
	fmt.Println(reflect.ValueOf(cat))            //{Nyiung 1}
	fmt.Println(reflect.ValueOf(cat).Kind())     //struct
	fmt.Println(reflect.ValueOf(cat).NumField()) // 2 --> field Name, Age

	if reflect.ValueOf(cat).Kind() == reflect.Struct {
		value := reflect.ValueOf(cat)
		countField := value.NumField()
		for i := 0; i < countField; i++ {
			fieldValue := value.Field(i).Kind()
			var result string

			switch fieldValue {
			case reflect.Int:
				result = "Integer"
			case reflect.String:
				result = "String"
			default:
				result = "Nil"
			}

			fmt.Printf("Field: %d type: %T value: %v --> %s\n", i, value.Field(i), value.Field(i), result)
		}
	}
}

func main() {
	cat1 := Cat{
		Name: "Nyiung",
		Age:  1,
	}
	example(cat1)

	// a := 56
	// x := reflect.ValueOf(a).Int()
	// fmt.Printf("type:%T value:%v\n", x, x)
	// b := "Naveen"
	// y := reflect.ValueOf(b).String()
	// fmt.Printf("type:%T value:%v\n", y, y)
}
