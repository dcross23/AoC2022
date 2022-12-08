package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func parseTreeMap(input []string) ([][]int, int, int) {
	var treeMap = make([][]int, 0)

	for i := 0; i < len(input); i++ {
		var line = input[i]

		var row = make([]int, 0)
		for j := 0; j < len(line); j++ {
			var value, _ = strconv.Atoi(string(line[j]))
			row = append(row, value)
		}

		treeMap = append(treeMap, row)
	}

	return treeMap, len(input), len(treeMap[0])
}

func part1(input []string) int {
	var treeMap, rows, columns = parseTreeMap(input)
	var edgesTrees = 2*columns + 2*(rows-2)
	var interiorTrees = 0

	for i := 1; i < rows-1; i++ {
		for j := 1; j < columns-1; j++ {
			var treeHeight = treeMap[i][j]

			//Check row left
			var isVisible = true
			for k := j - 1; k >= 0; k-- {
				if treeMap[i][k] >= treeHeight {
					isVisible = false
					break
				}
			}

			if isVisible == true {
				interiorTrees++
				continue
			}

			//Check row right
			isVisible = true
			for k := j + 1; k < columns; k++ {
				if treeMap[i][k] >= treeHeight {
					isVisible = false
					break
				}
			}

			if isVisible == true {
				interiorTrees++
				continue
			}

			//Check column up
			isVisible = true
			for k := i - 1; k >= 0; k-- {
				if treeMap[k][j] >= treeHeight {
					isVisible = false
					break
				}
			}

			if isVisible == true {
				interiorTrees++
				continue
			}

			//Check column down
			isVisible = true
			for k := i + 1; k < rows; k++ {
				if treeMap[k][j] >= treeHeight {
					isVisible = false
					break
				}
			}

			if isVisible == true {
				interiorTrees++
				continue
			}

		}
	}

	return interiorTrees + edgesTrees
}

func part2(input []string) int {
	var treeMap, rows, columns = parseTreeMap(input)

	var maxVisibleTrees = -1
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			var tree = treeMap[i][j]

			//Check row left
			var visibleLeft = 0
			for k := j - 1; k >= 0; k-- {
				if treeMap[i][k] >= tree {
					visibleLeft += 1
					break

				} else {
					visibleLeft += 1
				}
			}

			//Check row right
			var visibleRight = 0
			for k := j + 1; k < columns; k++ {
				if treeMap[i][k] >= tree {
					visibleRight += 1
					break

				} else {
					visibleRight += 1
				}
			}

			//Check column up
			var visibleUp = 0
			for k := i - 1; k >= 0; k-- {
				if treeMap[k][j] >= tree {
					visibleUp += 1
					break
				} else {
					visibleUp += 1
				}
			}

			//Check column down
			var visibleDown = 0
			for k := i + 1; k < rows; k++ {
				if treeMap[k][j] >= tree {
					visibleDown += 1
					break
				} else {
					visibleDown += 1
				}
			}

			var visibleTrees = visibleLeft * visibleRight * visibleUp * visibleDown
			if visibleTrees > maxVisibleTrees {
				maxVisibleTrees = visibleTrees
			}
		}
	}

	return maxVisibleTrees
}

func main() {
	var input = readInput("day8.input")

	fmt.Println("Part1: ", part1(input))
	fmt.Println("Part2: ", part2(input))
}
