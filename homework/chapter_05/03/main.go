package exercise

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = false

func exercise(str string) string {
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	return s
}
