package helper

import "fmt"

// import "fmt"

func returnFunc(name string) string {
	if name == "" {
		return "Hellow"
	}else{
		return "Hellow" + name
	}
}

func ReturnValue(){
result := returnFunc("Faisal")
fmt.Println(result)
}