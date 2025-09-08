package main

import "github.com/01-edu/z01"

func main() {
	STARTINGNUMBERASCII := 122
	for i := 0 ; i < 26 ; i ++ {
		z01.PrintRune(rune(STARTINGNUMBERASCII - i))
	}
	z01.PrintRune('\n')
}
