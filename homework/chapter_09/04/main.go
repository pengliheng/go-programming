package exercise

import "fmt"

func exercise(str string) string {
	bd := 1992
	for {
		fmt.Printf("%d\n", bd)
		if bd < 2019 {
			bd++
		} else {
			break
		}
	}
	return str
}
