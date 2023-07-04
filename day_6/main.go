package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func findDuplicates(str string) bool {
	for i := 0; i < len(str); i++ {
		numOccurences := strings.Count(str, string(str[i]))
		if numOccurences >= 2 {
			fmt.Println("dup")
			return true
		}
	}
	return false
}

func main() {

	fileData, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(fileData)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
		for i, _ := range scanner.Text() {

			if i <= len(scanner.Text())-14 {
				markers := scanner.Text()[i : i+14]
				duplicateFound := findDuplicates(markers)

				if !duplicateFound {
					fmt.Println("MArker found")
					fmt.Println(i + 14)
					return
				}
			}

		}
	}
}
