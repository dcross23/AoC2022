package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput(file string) []string {
	readFile, err := os.Open(file)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var input = make([]string, 0)
	for fileScanner.Scan() {
		var text = fileScanner.Text()
		input = append(input, text)
	}

	readFile.Close()
	return input
}

func part1(input []string) int {

	var numberOfInspectedItems = make([]int, 0)
	var allItems = make(map[int][]int, 0)

	//Load initial items for each monkey
	for i := 0; i < len(input); i++ {
		if strings.HasPrefix(input[i], "Monkey") {
			var items = strings.Split(input[i+1], ": ")[1]
			var itemsList = strings.Split(items, ", ")

			var monkey = strings.Split(input[i], " ")[1]
			monkey = strings.TrimRight(monkey, ":")
			var monkeyInt, _ = strconv.Atoi(monkey)

			allItems[monkeyInt] = make([]int, 0)
			for _, item := range itemsList {
				var itemInt, _ = strconv.Atoi(item)
				allItems[monkeyInt] = append(allItems[monkeyInt], itemInt)
			}

			numberOfInspectedItems = append(numberOfInspectedItems, 0)
		}
	}

	//Rounds
	for round := 0; round < 20; round++ {
		for i := 0; i < len(input); i++ {
			if strings.HasPrefix(input[i], "Monkey") {
				var monkey = strings.Split(input[i], " ")[1]

				monkey = strings.TrimRight(monkey, ":")
				var monkeyInt, _ = strconv.Atoi(monkey)

				for _, worryLevel := range allItems[monkeyInt] {
					//Inspect
					numberOfInspectedItems[monkeyInt]++

					var operationWithValue = strings.Split(input[i+2], "= old ")[1]

					var operation = strings.Split(operationWithValue, " ")[0]
					var value = strings.Split(operationWithValue, " ")[1]
					var intValue int

					//Operation
					if value == "old" {
						intValue = worryLevel
					} else {
						intValue, _ = strconv.Atoi(value)
					}

					switch operation {
					case "+":
						worryLevel += intValue
						break
					case "*":
						worryLevel *= intValue
						break
					}

					//Monkey gets bored
					worryLevel = worryLevel / 3

					//Test
					var testValue = strings.Split(input[i+3], "by ")[1]
					var testIntValue, _ = strconv.Atoi(testValue)

					var monkeyToThrow string

					if worryLevel%testIntValue == 0 {
						monkeyToThrow = strings.Split(input[i+4], "monkey ")[1]
					} else {
						monkeyToThrow = strings.Split(input[i+5], "monkey ")[1]
					}

					var monkeyToThrowInt, _ = strconv.Atoi(monkeyToThrow)
					allItems[monkeyToThrowInt] = append(allItems[monkeyToThrowInt], worryLevel)

					//Remove item from monkey
					allItems[monkeyInt] = allItems[monkeyInt][1:]
				}
			}
		}
	}

	var maxItem1 = 0
	var maxItem2 = 0

	for _, items := range numberOfInspectedItems {
		if items > maxItem1 {
			maxItem2 = maxItem1
			maxItem1 = items
		} else if items > maxItem2 {
			maxItem2 = items
		}
	}

	return maxItem1 * maxItem2
}

func part2(input []string) int {

	var numberOfInspectedItems = make([]int, 0)
	var allItems = make(map[int][]int, 0)
	var modulousCommonLimit = 1

	//Load initial items for each monkey
	for i := 0; i < len(input); i++ {
		if strings.HasPrefix(input[i], "Monkey") {
			var items = strings.Split(input[i+1], ": ")[1]
			var itemsList = strings.Split(items, ", ")

			var monkey = strings.Split(input[i], " ")[1]
			monkey = strings.TrimRight(monkey, ":")
			var monkeyInt, _ = strconv.Atoi(monkey)

			allItems[monkeyInt] = make([]int, 0)
			for _, item := range itemsList {
				var itemInt, _ = strconv.Atoi(item)
				allItems[monkeyInt] = append(allItems[monkeyInt], itemInt)
			}

			numberOfInspectedItems = append(numberOfInspectedItems, 0)

			var testValue = strings.Split(input[i+3], "by ")[1]
			var testIntValue, _ = strconv.Atoi(testValue)
			modulousCommonLimit *= testIntValue
		}
	}

	//Rounds
	for round := 0; round < 10000; round++ {
		for i := 0; i < len(input); i++ {
			if strings.HasPrefix(input[i], "Monkey") {
				var monkey = strings.Split(input[i], " ")[1]

				monkey = strings.TrimRight(monkey, ":")
				var monkeyInt, _ = strconv.Atoi(monkey)

				for _, worryLevel := range allItems[monkeyInt] {
					//Inspect
					numberOfInspectedItems[monkeyInt]++

					var operationWithValue = strings.Split(input[i+2], "= old ")[1]

					var operation = strings.Split(operationWithValue, " ")[0]
					var value = strings.Split(operationWithValue, " ")[1]
					var intValue int

					//Operation
					if value == "old" {
						intValue = worryLevel
					} else {
						intValue, _ = strconv.Atoi(value)
					}

					switch operation {
					case "+":
						worryLevel = (worryLevel + intValue) % modulousCommonLimit
						break
					case "*":
						worryLevel = (worryLevel * intValue) % modulousCommonLimit
						break
					}

					//Test
					var testValue = strings.Split(input[i+3], "by ")[1]
					var testIntValue, _ = strconv.Atoi(testValue)

					var monkeyToThrow string

					if worryLevel%testIntValue == 0 {
						monkeyToThrow = strings.Split(input[i+4], "monkey ")[1]
					} else {
						monkeyToThrow = strings.Split(input[i+5], "monkey ")[1]
					}

					var monkeyToThrowInt, _ = strconv.Atoi(monkeyToThrow)
					allItems[monkeyToThrowInt] = append(allItems[monkeyToThrowInt], worryLevel)

					//Remove item from monkey
					allItems[monkeyInt] = allItems[monkeyInt][1:]
				}
			}
		}
	}

	var maxItem1 = 0
	var maxItem2 = 0

	for _, items := range numberOfInspectedItems {
		if items > maxItem1 {
			maxItem2 = maxItem1
			maxItem1 = items
		} else if items > maxItem2 {
			maxItem2 = items
		}
	}

	return maxItem1 * maxItem2
}

func main() {
	var input = readInput("day11.input")

	fmt.Println("Part1: ", part1(input))
	fmt.Println("Part2: ", part2(input))
}
