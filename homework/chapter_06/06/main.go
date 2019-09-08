package exercise

import "fmt"

var a int

func exercise(str string) string {
	//  iota 4 year
	const (
		a = 2017 + iota
		b = 2017 + iota
		c = 2017 + iota
		d = 2017 + iota
	)
	fmt.Println(a, b, c, d)
	return str
}
