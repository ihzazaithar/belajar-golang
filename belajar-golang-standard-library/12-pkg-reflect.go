package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	// Field TipeData StructTag:value StructTag:value
	Name string `required:"true" max:"10"`
}

type Person struct {
	// Field TipeData StructTag:value StructTag:value
	Name    string `required:"true" max:"10"`
	Address string `required:"true" max:"10"`
	Age     int    `required:"true" max:"10"`
}

func readFile(value any) {

	var valueType reflect.Type = reflect.TypeOf(value)
	fmt.Println("Type Name", valueType.Name())
	// Membaca struktur data Struct
	for i := 0; i < valueType.NumField(); i++ {
		var valueField reflect.StructField = valueType.Field(i)
		fmt.Println(valueField.Name, "with type", valueField.Type)
		// Membaca StructTag
		fmt.Println(valueField.Tag.Get("required"))
		fmt.Println(valueField.Tag.Get("max"))
	}

}

// Validasi
func IsValid(value any) (result bool) {
	result = true
	t := reflect.TypeOf(value)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Tag.Get("required") == "true" {
			data := reflect.ValueOf(value).Field(i).Interface()
			result = data != ""
			if result == false {
				return result
			}
		}
	}
	return result
}

func main() {

	readFile(Sample{"Eko"})
	readFile(Person{"Budi", "Depok", 30})

	person := Person{
		Name:    "Joko",
		Address: "Solo",
		Age:     50,
	}

	fmt.Println(IsValid(person))

}
