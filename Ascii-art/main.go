package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

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
	argument := args[1]
	if len(argument) == 0 {
		return
	}
	splitSlice := strings.Split(argument, "\\n")
	if ctr != 854 {
		log.Fatal("there is some issues with your file standard.txt")
	}
	if strings.Replace(argument, "\\n", "", -1) == "" {
		for i := 0; i < strings.Count(argument, "\\n"); i++ {
			fmt.Println()
		}
		return
	}
	ascii.PrintAscii(splitSlice, Maprunes)
}
