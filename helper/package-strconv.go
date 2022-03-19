package helper

import (
	"fmt"
	"strconv"
)

func StrConv() {
	boolean, err := strconv.ParseBool("true")

	if err == nil {
		fmt.Println(boolean)

	}else {
		fmt.Println("INI ERROR: ", err.Error())
	}

	number, err := strconv.ParseInt("1000000", 10, 64 )

	if err == nil {
		fmt.Println(number)
	}else{
		fmt.Println("ERROR", err.Error() )
	}

	value := strconv.FormatInt(1000000, 10)

	fmt.Println(value)

	value2 := strconv.Itoa(1000)
	fmt.Println(value2)
}