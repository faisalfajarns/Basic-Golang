package helper

import (
	"fmt"
	"os"
)

func Os(){
	args := os.Args
	fmt.Println(args)

	// fmt.Println(args[1])
	// fmt.Println(args[2])
	// fmt.Println(args[3])

name, err :=	os.Hostname()

if err == nil {
	fmt.Println("HostName:", name)
}else {
	fmt.Println("Err", err.Error())
}

username := os.Getenv("USERNAME")
password := os.Getenv("PASSWORD")

fmt.Println(username)
fmt.Println(password)
}