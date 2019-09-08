package exercise

import "fmt"

func exercise(str string) string {
	s := [][]string{
		[]string{
			"James", "Bond", "Shaken, not stirred",
		},
		[]string{
			"Miss", "Moneypenny", "Helloooooo, James.",
		},
	}

	// fmt.Println(s)
	for j, e := range s {
		fmt.Println("record:", j)
		for i, ee := range e {
			fmt.Printf("\tindex: %v,\t value: %v\n", i, ee)
		}
	}
	return str
}
