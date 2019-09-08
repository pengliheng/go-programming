package exercise

import "fmt"

var a int

func exercise(str string) string {
	// assigns an int to a variable
	a = 42
	// prints that int in decimal, binary, and hex
	fmt.Printf("%d\t%b\t%#x\n", a, a, a)
	// shifts the bits of that int over 1 position to the left, and assigns that to a variable
	b := a << 1
	// prints that variable in decimal, binary, and hex
	fmt.Printf("%d\t%b\t%#x\n", b, b, b)
	return str
}
