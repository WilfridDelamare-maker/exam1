package main

import (
	"piscine"

	"github.com/01-edu/z01"
)

func main() {
	z01.PrintRune(piscine.FirstRune("Hello!"))
	z01.PrintRune(piscine.FirstRune("Salut!"))
	z01.PrintRune(piscine.FirstRune("Ola!"))
	z01.PrintRune('\n')
}

func FirstRune(s string) rune {
	if len(s) == 0 {
		return 0
	}
	rune_0 := []rune(s)
	return rune_0[0]
}