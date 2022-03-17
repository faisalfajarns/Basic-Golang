package helper

import "fmt"

func For() {
	var counter int = 1

	for   counter <= 2 {
		fmt.Println("Perulangan Ke", counter)
		counter++
		
	}

	//for dengan statement
	for count := 1 ; count <=2 ; count++ {
		fmt.Println("Valuenya", count)
	}

	//Mengambil data slice dengan for
	slice := []string{"Faisal", "Fajar", "Nursaid",}

	for i := 0 ; i < len(slice) ; i++ {
		fmt.Println(slice[i])
	}

	//For Range slice 
	test1 := []string{"test1", "test2", "test3"}

	for index, test := range test1 {
		fmt.Println("index", index, "=", test)
	}
	//for Range array
	test2 := [3]string{"test2.1", "test2.2", "test2.3",}

	for index, test2 := range test2 {
		fmt.Println("index", index, "=", test2)
	}

	//for Range map
	var person = map[interface{}]interface{}{
		"name":    "Faisal",
		"address": "Bekasi",
		"age" : 20,
	}
for key, value := range person {
	fmt.Println(key , "=", value)
}

}