package exercise

import "fmt"

func exercise(str string) string {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	y := append(x[:3], x[6:]...)
	fmt.Println(y)

	return str
}
