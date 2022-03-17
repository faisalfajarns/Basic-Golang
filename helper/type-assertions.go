package helper

import (
	"fmt"
)

func random() interface{} {
	return 100
}

func TypeAssertions() {
	var result interface{} = random()


	switch value := result.(type){
	case string: fmt.Println("Value", value, "is string")
	case int : fmt.Println("Value", value, "is int")
	default: fmt.Println("Unknown Type")
}
	// var resultString = result.(string)

	// fmt.Println(resultString)
}