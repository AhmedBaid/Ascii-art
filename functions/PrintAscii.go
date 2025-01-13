package ascii

import "fmt"

func Print(splitslice []string, Mapascii map[rune][]string) {
	for _, word := range splitslice {
		if word != "" {
			for line := 0; line < 8; line++ {
				for _, char := range word {
					if variable, exist := Mapascii[char]; exist {
						fmt.Print(variable[line])
					}
				}
				fmt.Println()
			}
		} else {
			fmt.Println()
		}
	}
}
