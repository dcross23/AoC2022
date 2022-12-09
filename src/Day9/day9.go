package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	x, y int
}

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

func calculateNewPosition(H Position, T Position) Position {
	var newPos = Position{T.x, T.y}

	if H.y > T.y+1 {
		if H.x > T.x {
			newPos = Position{T.x + 1, T.y + 1}

		} else if H.x < T.x {
			newPos = Position{T.x - 1, T.y + 1}

		} else {
			newPos = Position{T.x, T.y + 1}
		}
		return newPos
	}

	if H.y < T.y-1 {
		if H.x > T.x {
			newPos = Position{T.x + 1, T.y - 1}

		} else if H.x < T.x {
			newPos = Position{T.x - 1, T.y - 1}

		} else {
			newPos = Position{T.x, T.y - 1}
		}
		return newPos
	}

	if H.x > T.x+1 {
		if H.y > T.y {
			newPos = Position{T.x + 1, T.y + 1}

		} else if H.y < T.y {
			newPos = Position{T.x + 1, T.y - 1}

		} else {
			newPos = Position{T.x + 1, T.y}
		}
		return newPos
	}

	if H.x < T.x-1 {
		if H.y > T.y {
			newPos = Position{T.x - 1, T.y + 1}

		} else if H.y < T.y {
			newPos = Position{T.x - 1, T.y - 1}

		} else {
			newPos = Position{T.x - 1, T.y}
		}
		return newPos
	}

	return newPos
}

func part1(input []string) int {
	var visited = make(map[Position]int, 0)
	var currentH = Position{0, 0}
	var currentT = Position{0, 0}

	visited[currentT] = 1

	for _, movement := range input {
		var direction = strings.Split(movement, " ")[0]
		var distance, _ = strconv.Atoi(strings.Split(movement, " ")[1])

		for i := 0; i < distance; i++ {
			switch direction {
			case "U":
				currentH.y++
			case "D":
				currentH.y--
			case "L":
				currentH.x--
			case "R":
				currentH.x++
			}

			currentT = calculateNewPosition(currentH, currentT)
			visited[currentT] += 1
		}
	}

	return len(visited)
}

func part2(input []string) int {
	var visited = make(map[Position]int, 0)
	var currentSnake = make([]Position, 10)

	visited[currentSnake[9]] = 1

	for _, movement := range input {
		var direction = strings.Split(movement, " ")[0]
		var distance, _ = strconv.Atoi(strings.Split(movement, " ")[1])

		for i := 0; i < distance; i++ {
			switch direction {
			case "U":
				currentSnake[0].y++

				break

			case "D":
				currentSnake[0].y--
				break

			case "L":
				currentSnake[0].x--
				break

			case "R":
				currentSnake[0].x++
				break

			}

			for i := 1; i < 10; i++ {
				currentSnake[i] = calculateNewPosition(currentSnake[i-1], currentSnake[i])
			}

			visited[currentSnake[9]] += 1
		}
	}

	return len(visited)
}

func main() {
	var input = readInput("day9.input")

	fmt.Println("Part1: ", part1(input))
	fmt.Println("Part2: ", part2(input))
}
