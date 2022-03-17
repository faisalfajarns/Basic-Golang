package helper

import "fmt"

type Man struct {
	Name string
}

func (man *Man) married() {
	man.Name = "Mr. " + man.Name
	// fmt.Println("Di Method", man.Name)
}

func PointerMethod() {
	faisal := Man{"Faisal"}
	faisal.married()

	fmt.Println(faisal.Name)
}
//pointer method tidak perlu menggunakan &