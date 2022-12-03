package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func readInput(file string) []string {
	readFile, err := os.Open(file)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var plays = make([]string, 0)
	for fileScanner.Scan() {
		var text = fileScanner.Text()
		plays = append(plays, text)
	}

	readFile.Close()
	return plays
}

func part1(plays []string) int {
	var prioritiesSum = 0

	for _, rucksack := range plays {
		var firstHalf = rucksack[:len(rucksack)/2]
		var secondHalf = rucksack[len(rucksack)/2:]

		for _, c := range firstHalf {
			var idx = strings.Index(secondHalf, string(c))
			if idx != -1 {
				var priority = 0
				if unicode.IsUpper(c) {
					priority = int(c-'A') + 27
				} else {
					priority = int(c-'a') + 1
				}
				prioritiesSum += priority
				break
			}
		}
	}

	return prioritiesSum
}

func part2(plays []string) int {
	var prioritiesSum = 0

	for i := 0; i < len(plays); i += 3 {
		var rucksack1 = plays[i]
		var rucksack2 = plays[i+1]
		var rucksack3 = plays[i+2]

		for _, c := range rucksack1 {
			var idx1 = strings.Index(rucksack2, string(c))
			var idx2 = strings.Index(rucksack3, string(c))
			if idx1 != -1 && idx2 != -1 {
				var priority = 0
				if unicode.IsUpper(c) {
					priority = int(c-'A') + 27
				} else {
					priority = int(c-'a') + 1
				}
				prioritiesSum += priority
				break
			}
		}
	}

	return prioritiesSum
}

func main() {
	var plays = readInput("day3.input")

	fmt.Println("Part1: " + strconv.Itoa(part1(plays)))
	fmt.Println("Part2: " + strconv.Itoa(part2(plays)))
}
