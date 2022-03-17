package helper

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye" + name
}

func Valuefuncvalue() {
	sayGoodBye := getGoodBye

	result := sayGoodBye("Faisal")

	fmt.Println(result)
}