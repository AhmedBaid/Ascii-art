package main

import (
	"bufio"
	"log"
	"os"

	ascii "ascii/functions"
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

	ctr := 0

	for Myscanner.Scan() {

		text := Myscanner.Text()
		ctr++
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
	splitSlice := ascii.Split(word)
	if ctr != 854 {
		log.Fatal("there is some issues with your file standard.txt")
	}
	ascii.PrintAscii(splitSlice, Maprunes)
}
