package helper

import "fmt"


type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)

	fmt.Println("Hellow", nameFiltered)
}

func spamFilter(name string) string{
	if name =="anjing" {
		return "..."
	}else {
		return name
	}
}


func ResultFuncAsParam(){
	sayHelloWithFilter("Faisal", spamFilter)
	sayHelloWithFilter("anjing", spamFilter)
}