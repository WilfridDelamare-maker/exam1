package main

import (
	"fmt"
)

func main() {
	fmt.Println(CheckNumber("Hello"))
	fmt.Println(CheckNumber("Hello1"))
}

func CheckNumber(arg string) bool {
	count := 0
	l := len(arg)
	for i := 0 ; i <= l ; i++ {
		b := arg[i]
		if b >= '1' && b <= '9'{
			count ++
		}
	} 
	if count >= 1 {
		return false
	}
	return true
}