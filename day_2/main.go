package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	f, err := os.Open("input.txt")
	points := make(map[string]int)
	mapWinningComboLTR := make(map[string]string)
	mapDrawComboLTR := make(map[string]string)
	mapLosingComboLTR := make(map[string]string)
	AWARDED_WINNING_POINTS := 6
	AWARDED_DRAW_POINTS := 3

	// A -> rock, B -> paper, C -> scissor
	// X -> rock, Y -> paper, Z -> scissor
	mapWinningComboLTR["A"] = "Z"
	mapWinningComboLTR["B"] = "X"
	mapWinningComboLTR["C"] = "Y"
	mapDrawComboLTR["A"] = "X"
	mapDrawComboLTR["B"] = "Y"
	mapDrawComboLTR["C"] = "Z"
	mapLosingComboLTR["A"] = "Y"
	mapLosingComboLTR["B"] = "Z"
	mapLosingComboLTR["C"] = "X"

	points["A"] = 1
	points["B"] = 2
	points["C"] = 3
	points["X"] = 1
	points["Y"] = 2
	points["Z"] = 3

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	totalScore := 0
	for scanner.Scan() {
		// str := scanner.Text()
		// scoreOfRound := 0
		// scoreOfPiece := 0

		picks := strings.Split(scanner.Text(), " ")
		fmt.Println(picks)

		if picks[1] == "X" {
			userPick := mapWinningComboLTR[picks[0]]
			totalScore += points[userPick]
		} else if picks[1] == "Y" {
			userPick := mapDrawComboLTR[picks[0]]
			totalScore += points[userPick]
			totalScore += AWARDED_DRAW_POINTS

		} else {
			userPick := mapLosingComboLTR[picks[0]]
			totalScore += points[userPick]
			totalScore += AWARDED_WINNING_POINTS
		}
	}

	fmt.Println(totalScore)

}
