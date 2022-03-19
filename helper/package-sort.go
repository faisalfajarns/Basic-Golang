package helper

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (value UserSlice) Len() int {
	return len(value)
}

func (value UserSlice) Less(i, j int) bool {
	return value[i].Age < value[j].Age
}

func (value UserSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}

func Sort() {
	users := []User{
		{"Faisal", 25},
		{"Fahmi", 20},
		{"Tia", 25},
		{"Fajar", 21},
	}



	// user := UserSlice{}
	// user = append(user, users...)
	// user.Swap(0,3)



	fmt.Println(users)

	sort.Sort(UserSlice(users))

	fmt.Println(users)
}