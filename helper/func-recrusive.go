package helper

import "fmt"

func factorialLoop(value int) int {
	result := 1

	for i := value; i > 0; i-- {
		result = result * i
		fmt.Println("bb", i)
	}
	fmt.Println("aa",result)
	return result

}

func factorialRecursive(value int) int{
	if value == 1 {
		return 1

	}else{
	
		return value * factorialRecursive(value-1)
	}
}

func Recursive(){
	notRecursive := factorialLoop(5)
	fmt.Println(notRecursive)

	recursive := factorialRecursive(5)
	fmt.Println(recursive)
}
