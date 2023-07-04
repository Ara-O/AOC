package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func reverseFunc(s *Stack) {
	for i, j := 0, len((*s))-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}

type Stack []string

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) push(str string) {
	*s = append(*s, str)
}

func (s *Stack) display() {
	for _, stack := range *s {
		fmt.Println(stack)
	}
}

func (s *Stack) pop() (string, bool) {
	if (*s).isEmpty() {
		return "", false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func main() {

	var stacksTxtData []string
	var instructions []string
	var allStacks []Stack
	var columnsCalculated bool = false

	fileData, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(fileData)
	lengthOfRows := 0
	for scanner.Scan() {
		str := scanner.Text()

		if !columnsCalculated {
			if strings.Trim(str, " ") == "" {
				columnsCalculated = true
				continue
			}

			stacksTxtData = append(stacksTxtData, str)
		} else {
			lengthOfRows = len(strings.Split(stacksTxtData[len(stacksTxtData)-1], "   "))

			if len(allStacks) != lengthOfRows {
				for i := 0; i < lengthOfRows; i++ {
					var newStack Stack
					allStacks = append(allStacks, newStack)
				}
			}

			instructions = append(instructions, scanner.Text())
		}
	}

	stacksTxtData = stacksTxtData[:len(stacksTxtData)-1]
	for _, data := range stacksTxtData {
		for i, row := range data {
			if (i+1)%4 == 2 {
				if string(row) != " " {
					allStacks[(i-1)/4].push(string(row))
				}

			}
		}

	}

	for _, stacks := range allStacks {
		reverseFunc(&stacks)
	}

	for _, stacks := range allStacks {
		fmt.Println(stacks)
	}

	for _, instruction := range instructions {
		instructionsParsed := string(instruction)
		parsedInstr := strings.Split(instructionsParsed, " ")

		moveVal, err := strconv.Atoi(parsedInstr[1])
		if err != nil {
			log.Fatal(err)
		}

		moveFrom, err := strconv.Atoi(parsedInstr[3])
		moveTo, err := strconv.Atoi(parsedInstr[5])

		fmt.Printf("%d->%d\n", moveFrom, moveTo)

		valsToBePopped := []string{}
		for i := 0; i < moveVal; i++ {
			poppedVal, success := allStacks[moveFrom-1].pop()

			if success == false {
				log.Fatal(err)
			}

			valsToBePopped = append(valsToBePopped, poppedVal)
		}

		fmt.Println(valsToBePopped)
		for i := 0; i < len(valsToBePopped); i++ {
			allStacks[moveTo-1] = append(allStacks[moveTo-1], valsToBePopped[len(valsToBePopped)-(i+1)])
		}

		// fmt.Println(valsToBePopped)

		fmt.Println("-------  ---------- --------- --------")
		for _, stacks := range allStacks {
			fmt.Println(stacks)
		}
		// }
		// for _, crate := range stacksTxtData {
		// 	trimmedCrate := strings.Split(crate, "   ")
		// 	currentCrate := ""
		// 	column := 1
		// 	for _, crateT := range trimmedCrate {
		// 		if string(crateT) != "[" && string(crateT) != "]" {
		// 			if strings.Trim(string(crateT), " ") == "" {
		// 				column++
		// 			} else {
		// 				currentCrate = string(crateT)
		// 				fmt.Printf("Crate: %s, Column %d\n", currentCrate, column)
		// 			}
		// 		}
		// 	}
		// 	fmt.Println("-----------")
		// }

	}
}
