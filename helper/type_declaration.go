package helper

import "fmt"

func Type_Declaration() {
	type noKTP string
	type married bool

	var noKTPku noKTP = "12332423423423"
	var marriedStatus married = true
	fmt.Println(noKTPku)
	fmt.Println(marriedStatus)
}