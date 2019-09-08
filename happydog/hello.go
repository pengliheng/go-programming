package hello

import (
	"fmt"
)

var y = 42
type hotdog int

var b hotdog

func Hello() string {
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", b)
	return ""
}
