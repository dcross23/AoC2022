// hello world
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readInput(file string) []int {
	readFile, err := os.Open(file)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var calories []int
	var count int = 0
	for fileScanner.Scan() {
		var text = fileScanner.Text()
		if text == "" {
			calories = append(calories, count)
			count = 0
			continue
		}

		var value, _ = strconv.Atoi(text)
		count += value
	}

	if count != 0 {
		calories = append(calories, count)
	}

	readFile.Close()
	return calories
}

func getTop3MaxCalories(calories []int) []int {
	var maxCalories = make([]int, 3)
	for i := 0; i < 3; i++ {
		var max int = 0
		var maxIdx int = 0
		for i := 0; i < len(calories); i++ {
			if calories[i] > max {
				max = calories[i]
				maxIdx = i
			}
		}
		maxCalories[i] = max
		calories[maxIdx] = -1
	}
	return maxCalories
}

func main() {
	var calories = readInput("day1.input")
	var maxCalories = getTop3MaxCalories(calories)

	fmt.Println("Part1: " + strconv.Itoa(maxCalories[0]))
	fmt.Println("Part2: " + strconv.Itoa(maxCalories[0]+maxCalories[1]+maxCalories[2]))
}
