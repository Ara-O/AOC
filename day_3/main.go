package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	strToPoints := make(map[string]int)

	for i, letter := range "abcdefghijklmnopqrstuvwxyz" {
		strToPoints[string(letter)] = i + 1
	}

	for i, letter := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		strToPoints[string(letter)] = i + 27
	}

	totalPoints := 0
	for scanner.Scan() {
		str := scanner.Text()
		itemsLength := len(str)
		firstHalf := str[:itemsLength/2]
		secondHalf := str[itemsLength/2:]

		for _, letter := range firstHalf {
			if strings.Contains(secondHalf, string(letter)) {
				totalPoints += strToPoints[string(letter)]
				break
			}
		}
	}

	fmt.Println(totalPoints)
}
