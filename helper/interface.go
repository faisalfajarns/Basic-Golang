package helper

import "fmt"


type Person struct{
	Name string
}

type Animal struct{
	Name string
}

type HasName interface {
	//kontrak dengan GetName yang mereturn string
	GetName() string
}

func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

func (animal Animal) GetName() string {
	return animal.Name
}

func (person Person) GetName() string {
	return person.Name
}
func Interface() {
	
	var faisal  Person

	faisal.Name = "Faisal"
	SayHello(faisal)
	
	cat := Animal{
		Name: "abu",
	}
SayHello(cat)
}