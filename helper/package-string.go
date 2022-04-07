package helper

import (
	"fmt"
	"regexp"
	"strings"
)

func StringPackage() {
	name := "Faisal Fajar Nursaid.jpg" 
	address := "BEKASI"
	country := "indonesia"
	fmt.Println(strings.Trim(name, " ."))
	fmt.Println(strings.ToLower(address))
	fmt.Println(strings.ToUpper(country))
	
	

split := regexp.MustCompile("[. ]").Split(name, -1)
fmt.Println(strings.Join(split, "-"))

users := []string{
	"name",
	"age",
}
fmt.Println(strings.Join(users, ", "))
}