package exercise_01

import "fmt"

func exercise_01(str string) string {
	switch {
	case 4 == 3:
		fmt.Println("111")
	case 4 == 4:
		fmt.Println("222")
		fallthrough			// control flow go through
	case 5 == 5:
		fmt.Println("333")
	}
	return str
}
