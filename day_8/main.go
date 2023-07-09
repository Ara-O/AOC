package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var treeStructure [99][99]int

func main() {
	visibleTrees := 0
	fileData, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(fileData)
	lengthOfTreeColumns := 0
	lengthOfTreeRows := 0
	row := 0
	for scanner.Scan() {
		lengthOfTreeColumns = len(scanner.Text())
		for idx, i := range scanner.Text() {
			i_int, _ := strconv.Atoi(string(i))
			treeStructure[row][idx] = i_int
		}
		row++
	}

	fmt.Println(treeStructure)

	lengthOfTreeRows = row
	fmt.Println(lengthOfTreeColumns, lengthOfTreeRows)
	for i_idx, i := range treeStructure {
		for j_idx, _ := range i {
			if i_idx == 0 || j_idx == 0 || i_idx == (lengthOfTreeRows-1) || j_idx == (lengthOfTreeColumns-1) {
				visibleTrees++
				continue
			}

			highest0x := true
			highest0y := true
			highestCurx := true
			highestCury := true
			//Test for left
			for x := 0; x < i_idx; x++ {
				if treeStructure[x][j_idx] >= treeStructure[i_idx][j_idx] {
					// fmt.Println("Tree - ", treeStructure[i_idx][j_idx], " is not highest in 0 - curr row")
					highest0x = false
					break
				}
			}

			for x := i_idx + 1; x < lengthOfTreeRows; x++ {
				if treeStructure[x][j_idx] >= treeStructure[i_idx][j_idx] {
					highestCurx = false
					break
				}
			}

			for y := 0; y < j_idx; y++ {
				if treeStructure[i_idx][y] >= treeStructure[i_idx][j_idx] {
					highest0y = false
					break
				}
			}

			for y := j_idx + 1; y < row; y++ {
				if treeStructure[i_idx][y] >= treeStructure[i_idx][j_idx] {
					highestCury = false
					break
				}
			}

			if highest0x || highest0y || highestCurx || highestCury {
				fmt.Println("Tree", i_idx, " ", j_idx, "is highest", highest0x, " - ", highest0y, " - ", highestCurx, " - ", highestCury)
				visibleTrees++
			}

		}
		fmt.Println("")
	}

	// fmt.Println(treeStructure)
	fmt.Println("Visible trees", visibleTrees)
}
