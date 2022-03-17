package database

import "fmt"

var connection string

func init() {
	fmt.Println("Init dipanggil")
	connection = "MySQL"
}

func GetDataBase() string {
	return connection
}