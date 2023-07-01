package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	totalEnclosures := 0
	for scanner.Scan() {
		str := scanner.Text()

		pair := strings.Split(str, ",")

		assignmentOne := pair[0]
		assignmentTwo := pair[1]

		assignmentOneSplit := strings.Split(assignmentOne, "-")
		assignmentTwoSplit := strings.Split(assignmentTwo, "-")

		asgnOneStart, _ := strconv.Atoi(assignmentOneSplit[0])
		asgnOneEnd, _ := strconv.Atoi(assignmentOneSplit[1])
		asgnTwoStart, _ := strconv.Atoi(assignmentTwoSplit[0])
		asgnTwoEnd, _ := strconv.Atoi(assignmentTwoSplit[1])

		if asgnOneStart <= asgnTwoStart && asgnOneEnd >= asgnTwoEnd {
			totalEnclosures++

		} else {
			if asgnTwoStart <= asgnOneStart && asgnTwoEnd >= asgnOneEnd {
				totalEnclosures++
			}
		}
	}
	fmt.Println(totalEnclosures)
}
