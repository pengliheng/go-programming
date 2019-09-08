package exercise

import (
	"fmt"
)

type hotdog = int

var x hotdog
var y int

func exercise(str string) string {
	fmt.Println(x)
	fmt.Printf("%T", x)
	x = 23
	fmt.Println(x)
	y = x
	fmt.Println(y)
	fmt.Printf("%T", y)
	return str
}
