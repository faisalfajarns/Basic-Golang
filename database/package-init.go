package database

import "fmt"

func PackageInit(){
	result := GetDataBase()
	fmt.Println(result)
}