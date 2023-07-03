package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func expandArray(num1 int, num2 int) []int {
	var generatedArr []int
	for i := num1; i <= num2; i++ {
		generatedArr = append(generatedArr, i)
	}
	return generatedArr
}

func contains(arr []int, num int) bool {
	for _, n := range arr {
		if n == num {
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

		genArrAsgnOne := expandArray(asgnOneStart, asgnOneEnd)
		genArrAsgnTwo := expandArray(asgnTwoStart, asgnTwoEnd)

		for _, num := range genArrAsgnOne {
			if contains(genArrAsgnTwo, num) {
				totalEnclosures++
				break
			}
		}

	}
	fmt.Println(totalEnclosures)

}
