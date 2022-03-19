package helper

import (
	"fmt"
	"reflect"
)

//struct dengan tag
type Sample struct {
	Name string `required:"true" max:"10"` 
}


func isValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	fmt.Println("T::",t)
	fmt.Println(t.NumField())
	for i:=0; i <t.NumField(); i++{
			field := t.Field(i)
			fmt.Println("FIELD::", field)
		if field.Tag.Get("required") == "true" {
			value :=  reflect.ValueOf(data).Field(i).Interface()
			fmt.Println(reflect.TypeOf(value))
			if value == ""{

				return false
			}
		}
	}
	return true
}


func Reflection() {
	sample := Sample{"Faisal"}

	var sampleType reflect.Type = reflect.TypeOf(sample)
	
	fmt.Println(sampleType.NumField())

	fmt.Println(sampleType.Field(0).Name)
	fmt.Println(sampleType.Field(0).Tag.Get("required"))
	fmt.Println(sampleType.Field(0).Tag.Get("max"))
	// sample.Name = ""
	fmt.Println(isValid(sample))
}