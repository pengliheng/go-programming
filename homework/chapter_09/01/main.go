package exercise

import "fmt"

func exercise(str string) string {
	for i := 0; i < 10000; i++ {
		fmt.Println(i)
	}
	return str
}
