package helper

import "fmt"

func Comparation_Operator() {

	var (
		lastName  string
		firstName string
		compare bool
		val1 int32
		val2 int32
		comp bool
	)

	firstName = "Faisal"
	lastName = "Fajar"

	compare = firstName == lastName

	fmt.Println(compare)

	val1 = 1000
	val2 = 2000

	comp = val1 < val2

	fmt.Println(comp)
}