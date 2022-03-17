package helper

import "fmt"

type Blacklist func(string) bool

func resgisterUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("Blocked", name)
	}else {
		fmt.Println("Not Blocked")
	}
}

func AnonymousFunc(){
	blacklist := func(name string) bool {
		return name == "admin"
	}
	resgisterUser("admin", blacklist)
	resgisterUser("faisal", blacklist)
}