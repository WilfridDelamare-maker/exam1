package student

import "github.com/01-edu/z01"

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

