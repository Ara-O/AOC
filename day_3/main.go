package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Group struct {
	group1 string
	group2 string
	group3 string
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

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

	// Part 1

	// totalPoints := 0
	// for scanner.Scan() {
	// 	str := scanner.Text()
	// 	itemsLength := len(str)
	// 	firstHalf := str[:itemsLength/2]
	// 	secondHalf := str[itemsLength/2:]

	// 	for _, letter := range firstHalf {
	// 		if strings.Contains(secondHalf, string(letter)) {
	// 			totalPoints += strToPoints[string(letter)]
	// 			break
	// 		}
	// 	}
	// }

	// fmt.Println(totalPoints)

	//Part 2
	group := 0
	totalPoints := 0
	var rucksackGroup Group
	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		group++
		switch group {
		case 1:
			rucksackGroup.group1 = scanner.Text()
		case 2:
			rucksackGroup.group2 = scanner.Text()
		case 3:
			rucksackGroup.group3 = scanner.Text()
			var sharedLettersG1G2 []string
			for _, letter := range rucksackGroup.group1 {
				if strings.Contains(rucksackGroup.group2, string(letter)) && !contains(sharedLettersG1G2, string(letter)) {
					sharedLettersG1G2 = append(sharedLettersG1G2, string(letter))
				}
			}

			for _, letter := range sharedLettersG1G2 {
				if strings.Contains(rucksackGroup.group3, string(letter)) {
					totalPoints += strToPoints[string(letter)]
				}
			}

			// fmt.Printf("Shared letters - %s\n", sharedLettersG1G2)

			group = 0
		}

	}
	fmt.Printf("Total points %d\n", totalPoints)
}
