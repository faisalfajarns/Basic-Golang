package helper

import (
	"errors"
	"fmt"
)

func pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("pembagi tidak boleh 0")
	}else {
		result := nilai/pembagi
		return result, nil
	}
}

func ErrorInterface() {
	
hasil, err :=pembagian(10,0)

if err == nil {
	fmt.Println("Hasil", hasil)
}else{
	fmt.Println("Error", err.Error())
}
}