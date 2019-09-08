package exercise_01

import "fmt"

func exercise_01(str string) string {
	x := 1023
	// print [decimal], [binary], and [hex]
	fmt.Printf("%d\t%b\t%#x\n", x, x, x)
	return str
}
