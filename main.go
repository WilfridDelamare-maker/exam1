package main

import (
	"fmt"
)

func main() {
	fmt.Println(CountChar("Hello World", 'l'))
	fmt.Println(CountChar("5  balloons", 5))
	fmt.Println(CountChar("   ", ' '))
	fmt.Println(CountChar("The 7 deadly sins", '7'))
}

func CountChar(str string, c rune) int {
	placechar := []int(str)
	count := 0
	if len(str) == 0 {
		return 0
	}
	for i := 0 ; i <= len(str) ; i++ {
		return placechar[i] 
			if placechar[i] == c {
				count ++
			}
		}
	return count
}
