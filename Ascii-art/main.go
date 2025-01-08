package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		log.Fatal("make sure that the arguments is two")
	}
	Filename := "standard.txt"
	file, err := os.Open(Filename)
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()
	run := []string{}
	Maprunes := map[rune][]string{}
	count := 0
	space := ' '
	Myscanner := bufio.NewScanner(file)
	for Myscanner.Scan() {
		text := Myscanner.Text()
		if text != "" {
			run = append(run, text)
			count++
		}
		if count == 8 {
			Maprunes[space] = run
			space++
			run = []string{}
			count = 0
		}
	}
	word := args[1]

	if len(word) == 0 {
		return
	}
	for _, char := range word {
		if char > 126 || char < 32 {
			log.Fatal("Error ajmii special charachtere")
		}
	}
	splitSlice := strings.Split(word, "\\n")

	for _, word := range splitSlice {
		for line := 0; line < 8; line++ {
			for _, char := range word {
				if string(char) != "" {
					if ascii, exist := Maprunes[char]; exist {
						fmt.Print(ascii[line])
					}
				}
			}
			fmt.Println()
		}
	}

	// fmt.Println(splitSlice)
}
