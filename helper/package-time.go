package helper

import (
	"fmt"
	"time"
)

func Time() {
	now := time.Now()

	fmt.Println(now.Local())
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())


	utc := time.Date(2022, 10, 10, 10, 10, 10, 10, time.UTC )

	fmt.Println(utc)

	layout := time.RFC1123
	
	parse, _ := time.Parse(layout, "2022-05-19")

	fmt.Println(parse)
}