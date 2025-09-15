package main

import (
	"fmt"
)

func main() {
	fmt.Println(CountAlpha("Hello world"))
	fmt.Println(CountAlpha("H e l l o"))
	fmt.Println(CountAlpha("H1e2l3l4o"))
}

func CountAlpha(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if (c >='a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
				count ++
		}		
	}
	return count
}
	
