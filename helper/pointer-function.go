package helper

import "fmt"

func changeCountryToIndonesia(alamat *Address) {
	alamat.Country = "Indonesia"
}

func PointerFunction() {

	var alamat = Address{
		City:     "Tangerang",
		Province: "Banten",
		Country:  "",
	}
	changeCountryToIndonesia(&alamat)
	fmt.Println(alamat)
}