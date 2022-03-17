package helper

import "fmt"

func Break() {
	for i := 0; i < 10; i++ {
		fmt.Println("break", i)
		if i == 5{
			break
		}
	}

	//continue

	for c := 0; c < 10; c++ {
		if c%2 == 0{
			continue
		}
		fmt.Println("continue", c)
	}

}