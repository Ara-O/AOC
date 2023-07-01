package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var allTotalCalories []int64
	var calorie int64
	for scanner.Scan() {
		if strings.Trim(scanner.Text(), " ") == "" {
			allTotalCalories = append(allTotalCalories, calorie)
			calorie = 0
		} else {
			calorieInt, err := strconv.ParseInt(scanner.Text(), 10, 64)

			if err != nil {
				log.Fatal(err)
			}

			calorie += calorieInt
		}
	}

	slice := make([]*int64, len(allTotalCalories))
	for i := range allTotalCalories {
		slice[i] = &allTotalCalories[i]
	}

	// Use sort.Slice to sort the slice
	sort.Slice(slice, func(i, j int) bool {
		return *slice[i] < *slice[j]
	})

	// Create a new sorted int64 maxCalorieay from the sorted slice
	sortedmaxCalorie := make([]int64, len(slice))
	for i, ptr := range slice {
		sortedmaxCalorie[i] = *ptr
	}

	// Print the sorted maxCalorieay
	fmt.Println(sortedmaxCalorie)

}
