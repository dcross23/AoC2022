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

	var sections = make([]string, 0)
	for fileScanner.Scan() {
		var text = fileScanner.Text()
		sections = append(sections, text)
	}

	readFile.Close()
	return sections
}

func reverseArray(arr []rune) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func parseInput(input []string) ([][]rune, [][]int) {
	var firstCrateIdx = 1
	var parsingStacks = true
	var crates = make([]string, 0)
	var moves = make([][]int, 0)

	for l, line := range input {
		if line == "" {
			continue
		}

		if parsingStacks {
			if input[l+1] == "" {
				parsingStacks = false
				continue
			}

			var crate string
			for i := firstCrateIdx; i < len(line); i += 4 {
				crate += string(line[i])
			}

			crates = append(crates, crate)

		} else {
			var move = make([]int, 0)
			var s = strings.Split(line, " ")
			var numToMove, _ = strconv.Atoi(s[1])
			var from, _ = strconv.Atoi(s[3])
			var to, _ = strconv.Atoi(s[5])
			move = append(move, numToMove, from-1, to-1)
			moves = append(moves, move)
		}
	}

	var stacks = make([][]rune, len(crates[0]))
	for _, crate := range crates {
		for j, c := range crate {
			if c != ' ' {
				stacks[j] = append(stacks[j], c)
			}
		}
	}

	for _, s := range stacks {
		reverseArray(s)
	}

	return stacks, moves
}

func part1(input []string) string {
	var stacks, moves = parseInput(input)
	for _, move := range moves {
		var numToMove = move[0]
		for i := 0; i < numToMove; i++ {
			var from = move[1]
			var to = move[2]

			var crate = stacks[from][len(stacks[from])-1]
			stacks[from] = stacks[from][:len(stacks[from])-1]
			stacks[to] = append(stacks[to], crate)
		}
	}

	var result string
	for _, s := range stacks {
		result += string(s[len(s)-1])
	}
	return result
}

func part2(input []string) string {
	var stacks, moves = parseInput(input)
	for _, move := range moves {
		var numToMove = move[0]
		var from = move[1]
		var to = move[2]

		var lenFrom = len(stacks[from])

		for i := lenFrom - numToMove; i < lenFrom; i++ {
			stacks[to] = append(stacks[to], stacks[from][i])
		}
		stacks[from] = stacks[from][:len(stacks[from])-numToMove]

	}

	var result string
	for _, s := range stacks {
		result += string(s[len(s)-1])
	}
	return result
}

func main() {
	var input = readInput("day5.input")

	fmt.Println("Part1: ", part1(input))
	fmt.Println("Part2: ", part2(input))
}
