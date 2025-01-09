package ascii

import (
	"fmt"
	"log"
)

func PrintAscii(splitSlice []string, Maprunes map[rune][]string) {
	for _, word := range splitSlice {
		if word != "" {
			for line := 0; line < 8; line++ {
				for _, char := range word {
					if len(string(char)) == 0 {
						fmt.Println()
					}
					if char > 126 || char < 32 {
						log.Fatal("Error ajmii special charachtere 😍")
					}
					if string(char) != "" {
						if ascii, exist := Maprunes[char]; exist {
							fmt.Print(ascii[line])
						}
					}
				}
				fmt.Println()
			}
		}else{
			fmt.Println()
		}
	}

	
}
