package main

import (
	"piscine"
	"fmt"
)

func main() {
	fmt.Print(student.PrintIf("abcdefz"))
	fmt.Print(student.PrintIf("abc"))
	fmt.Print(student.PrintIf(""))
	fmt.Print(student.PrintIf("14"))
}