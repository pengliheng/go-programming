package exercise

import "fmt"

func exercise(str string) string {
	for i := 10; i <= 100; i++ {

		if i%4 == 0 {
			fmt.Println(i)
		}

	}
	return str
}
