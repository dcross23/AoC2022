package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
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

func parseInput(input []string) ([][]int, Point, Point) {
	var heightMap = make([][]int, 0)
	var start, end Point

	for i := 0; i < len(input); i++ {
		var row = make([]int, 0)
		for j, c := range input[i] {
			if c == 'S' {
				row = append(row, int('a'))
				start = Point{i, j}
			} else if c == 'E' {
				row = append(row, int('z'))
				end = Point{i, j}
			} else {
				row = append(row, int(c))
			}
		}
		heightMap = append(heightMap, row)
	}

	return heightMap, start, end
}

func part1(heightMap [][]int, start Point, end Point) int {

	var toVisitQueue = make([]Point, 0)
	var visited = make(map[Point]bool, 0)
	var distance = make(map[Point]int, 0)

	toVisitQueue = append(toVisitQueue, start)
	for {
		if len(toVisitQueue) == 0 {
			return -1
		}

		var current = toVisitQueue[0]
		toVisitQueue = toVisitQueue[1:]
		visited[current] = true

		if current.x == end.x && current.y == end.y {
			return distance[end]
		}

		var neighbours = []Point{
			Point{current.x - 1, current.y},
			Point{current.x + 1, current.y},
			Point{current.x, current.y - 1},
			Point{current.x, current.y + 1},
		}

		for _, neighbour := range neighbours {
			// Check if neighbour is out of bounds
			if neighbour.x < 0 || neighbour.x >= len(heightMap) || neighbour.y < 0 || neighbour.y >= len(heightMap[0]) {
				continue
			}

			// Check if neighbour is already visited
			if _, ok := visited[neighbour]; ok {
				continue
			}

			// Check if neighbour is too high
			if (heightMap[neighbour.x][neighbour.y] - heightMap[current.x][current.y]) > 1 {
				continue
			}

			if distance[neighbour] == 0 {
				toVisitQueue = append(toVisitQueue, neighbour)
				distance[neighbour] = distance[current] + 1

			} else if distance[neighbour] > distance[current] {
				distance[neighbour] = distance[current] + 1
			}
		}
	}
}

func part2(heightMap [][]int, end Point) int {

	var distances = make([]int, 0)

	for i := 0; i < len(heightMap); i++ {
		for j := 0; j < len(heightMap[0]); j++ {
			if heightMap[i][j] == int('a') {
				var distance = part1(heightMap, Point{i, j}, end)

				if distance != -1 {
					distances = append(distances, distance)
				}

			}
		}
	}

	var min = distances[0]
	for _, distance := range distances {
		if distance < min {
			min = distance
		}
	}
	return min
}

func main() {
	var input = readInput("day12.input")
	var heightMap, start, end = parseInput(input)

	fmt.Println("Part1: ", part1(heightMap, start, end))
	fmt.Println("Part2: ", part2(heightMap, end))
}
