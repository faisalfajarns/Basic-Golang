package helper

import "fmt"



func If_Expression() {

	var name = "Fajar"


	if name == "Faisal" {
		fmt.Println("Faisal nih")
	}else if name == "Joko"{
		fmt.Println("Joko nih")
	}else {
		fmt.Println("Salah")
	}


	//short statement if

	if length := len(name); length >5 {
		fmt.Println("Panjang")
	}else {
		fmt.Println("Pendek")
	}
}