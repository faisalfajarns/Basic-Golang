package helper

import "fmt"

// representasi membuat data lebih cocok menggunakan struct
type Customer struct {
	Name, Address string
	Age           int
}

//struct method
func(customer Customer) sayHi( name string){
	fmt.Println("Hellow", name, "My Name Is", customer.Name)
}

func Struct() {
	

	var data Customer 
	data.Name = "Faisal"
	data.Address = "Bekasi"
	data.Age = 25

	fmt.Println(data)

	//struct literal
	data2 := Customer{
		Name: "Fahmi",
		Address: "Bogor",
		Age: 20,
	}

	fmt.Println(data2)

	data.sayHi("Fahmi")

}

