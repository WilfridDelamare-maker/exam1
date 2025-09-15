package main

import (
	"fmt"
)

func main() {
	fmt.Print(PrintIf("abcdefz"))
	fmt.Print(PrintIf("abc"))
	fmt.Print(PrintIf(""))
	fmt.Print(PrintIf("14"))
}

func PrintIf(str string) string {
	if len(str) >= 3 {
		z01.PrintRune(rune(71))
	}
	if len(str) == 0 {
		z01.PrintRune(rune(71))
	} else {
		return "Invalid Input" 
	}
	z01.PrintRune('\n')
}