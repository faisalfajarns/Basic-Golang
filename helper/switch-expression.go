package helper

import "fmt"



func Switch_Expression() {

	name := "Faja"

switch name {
case "Faisal":
	fmt.Println("Hello Faisal")
case "Fajar" :
	fmt.Println("Hello Fajar")
case "Nursaid" :
	fmt.Println("Hello Nursaid")
default : 
	fmt.Println("Coba Lagi")
}

//switch short statement

switch length := len(name); length > 5{
case true : 
fmt.Println("Panjang")
case false :
	fmt.Println("Pendek")
}

//switch tanpa experssion
panjang := len(name)
	switch{
	case panjang > 10 :
		fmt.Println("Panjang Banget")
	case panjang > 5 :
		fmt.Println("Pendek")
		default :
		fmt.Println("Benar")
	}
	
}