package helper

import (
	"fmt"
)

type Address struct {
	City, Province, Country string
}

func Pointer() {
	//pass by value
	// address1 := Address{"Bekasi", "Jawa Barat", "Indonesia"}
	// address2 := address1

	// address2.City = "Bandung"
	// fmt.Println(address1)
	// fmt.Println(address2)

	//pass by reference menggunakan &
	address1 := Address{"Bekasi", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1
	var address3 *Address = &address1
	
	address3.City = "Tandung"
	
	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	*address2 = Address{"Malang", "Jawa Timur", "Indonesia"}
	fmt.Println("NEW",*address2)

	// fmt.Println(address2)
	// fmt.Println(address3)
	a:= map[string]string{"a":"a",}
	b:= &a
	c := *b
	fmt.Println(a["a"], (*b)["a"], c["a"])

	//menggunakan new() untuk membuat pointer kosong
	var address4 *Address = new(Address)
	address4.City = "Bali"
	fmt.Println(address4)
}