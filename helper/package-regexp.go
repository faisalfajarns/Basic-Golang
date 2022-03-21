package helper

import (
	"fmt"
	"regexp"
)

func Regexp() {
	var regex *regexp.Regexp = regexp.MustCompile("f([a-z])i")

	fmt.Println(regex.MatchString("fai"))
	fmt.Println(regex.MatchString("fai"))
	fmt.Println(regex.MatchString("fAi"))
	
	search := regex.FindAllString("faisal fajar nursaid", 1)

	fmt.Println(search)

}