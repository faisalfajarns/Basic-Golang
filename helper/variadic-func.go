package helper

import "fmt"

func variadic(numbers ...int) int {
	total := 0

	for _, value := range numbers {
		total = total + value
	}

	return total
}

func ValueVariadic() {
	total := variadic(10, 20, 30, 40, 60)

	fmt.Println(total)

	slice := []int{1,2,3,4,5}

	totals := variadic(slice...)

	fmt.Println(totals)
}