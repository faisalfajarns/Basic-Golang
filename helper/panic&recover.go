package helper

import "fmt"

func endApp() {
	fmt.Println("App Done")
	msg := recover()
	fmt.Println("Err :", msg)
}

//defer func always on top
func RunApp(err bool){
	defer endApp()
	if err {
		panic("Apps Error")
	}
	fmt.Println("App Run")
}