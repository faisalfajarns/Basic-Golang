package helper

import "fmt"

func Map() {
	var person = map[interface{}]interface{}{
		"name":    "Faisal",
		"address": "Bekasi",
		"age" : 20,
	}

	
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["age"])

	//menambah data ke map
	person["title"] = "programmer"

	fmt.Println(person)

	//merubah data map
	person["name"] = "Fajar"
	fmt.Println(person)

	//menghapus data map
	delete(person, "title")
	fmt.Println(person)

	//membuat map baru
	var newPerson = make(map[interface{}]interface{})

	newPerson["name"] = "Fahmi"
	newPerson["address"] = "Bogor"
	newPerson["age"] = 20

	fmt.Println(newPerson)
}