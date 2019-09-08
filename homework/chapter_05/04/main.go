package exercise

import (
	"fmt"
)

type age = int

var x age

func exercise(str string) string {
	x = 28
	fmt.Printf("%v\t%T", x, x)
	return str
}
