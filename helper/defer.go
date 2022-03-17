package helper

import "fmt"

func logging() {
	fmt.Println("Selesai")
}

func RunApplication(val int){
	defer logging()
	fmt.Println("Run")
	res := 10/ val
	fmt.Println("RES::", res)
}