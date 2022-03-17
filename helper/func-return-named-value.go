package helper

import "fmt"

func getFullName() (firstName string, middleName string, lastName string) {
	firstName = "Faisal"
	middleName = "Fajar"
	lastName = "Nursaid"

	return
}

func ReturnFullName() {
	a, b, c := getFullName()

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
