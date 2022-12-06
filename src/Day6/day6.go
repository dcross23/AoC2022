package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readInput(file string) string {
	readFile, err := os.Open(file)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var input string
	for fileScanner.Scan() {
		input = fileScanner.Text()
	}

	readFile.Close()
	return input
}

func part1(input string) int {
	var idx = 0
	for idx <= len(input)-4 {
		var valid = true
		var sop = input[idx : idx+4]
		for i := 0; i < len(sop); i++ {
			if strings.Count(sop, string(sop[i])) > 1 {
				idx += 1
				valid = false
				break
			}
		}

		if valid {
			return idx + 4
		}

	}

	return -1
}

func part2(input string) int {
	var idx = 0
	for idx <= len(input)-14 {
		var valid = true
		var sop = input[idx : idx+14]
		for i := 0; i < len(sop); i++ {
			if strings.Count(sop, string(sop[i])) > 1 {
				idx += 1
				valid = false
				break
			}
		}

		if valid {
			return idx + 14
		}

	}

	return -1
}

func main() {
	var input = readInput("day6.input")

	fmt.Println("Part1: ", part1(input))
	fmt.Println("Part2: ", part2(input))
}
