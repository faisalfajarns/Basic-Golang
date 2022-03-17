package helper

import "fmt"

func Array() {
	var (
		names = [3]string{
			"faisal",
			"fajar",
			"nursaid",
		
		}
		values = []int{
			1,2,3,
		}
	)
	
	fmt.Println(names)
	fmt.Println(len(names)) //panjang data array
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	fmt.Println(values)

	values[1] = 4 //mengubah isian array
	fmt.Println(values)
	
	
	
}