package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
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

func loadRocks(input []string) (map[Point]bool, int) {
	var rocks = make(map[Point]bool, 0)
	var maxY = math.MinInt64

	for _, rockrange := range input {
		var vertices = strings.Split(rockrange, " -> ")

		for i := 0; i < len(vertices)-1; i++ {
			var x1, _ = strconv.Atoi(strings.Split(vertices[i], ",")[0])
			var y1, _ = strconv.Atoi(strings.Split(vertices[i], ",")[1])
			var x2, _ = strconv.Atoi(strings.Split(vertices[i+1], ",")[0])
			var y2, _ = strconv.Atoi(strings.Split(vertices[i+1], ",")[1])

			if x1 == x2 {
				if y1 <= y2 {
					for j := y1; j <= y2; j++ {
						var newPoint = Point{x1, j}
						rocks[newPoint] = true
					}
				} else {
					for j := y2; j <= y1; j++ {
						var newPoint = Point{x1, j}
						rocks[newPoint] = true
					}
				}

			} else if y1 == y2 {
				if x1 <= x2 {
					for j := x1; j <= x2; j++ {
						var newPoint = Point{j, y1}
						rocks[newPoint] = true
					}
				} else {
					for j := x2; j <= x1; j++ {
						var newPoint = Point{j, y1}
						rocks[newPoint] = true
					}
				}
			}

			if y1 > maxY {
				maxY = y1
			}

			if y2 > maxY {
				maxY = y2
			}
		}
	}

	return rocks, maxY
}

func poseSand(currentPos Point, rocks, sands map[Point]bool, abysmStart int) *Point {
	//Drops to abysm
	if currentPos.y+1 > abysmStart {
		return nil
	}

	//Drop one down
	var newPoint = Point{currentPos.x, currentPos.y + 1}
	if rocks[newPoint] == false && sands[newPoint] == false {
		return poseSand(newPoint, rocks, sands, abysmStart)
	}

	//Drop down left
	newPoint = Point{currentPos.x - 1, currentPos.y + 1}
	if rocks[newPoint] == false && sands[newPoint] == false {
		return poseSand(newPoint, rocks, sands, abysmStart)
	}

	//Drop down right
	newPoint = Point{currentPos.x + 1, currentPos.y + 1}
	if rocks[newPoint] == false && sands[newPoint] == false {
		return poseSand(newPoint, rocks, sands, abysmStart)
	}

	//Posed
	return &currentPos
}

func poseSand2(currentPos Point, rocks, sands map[Point]bool, floor int) Point {
	//Posed in floor
	if currentPos.y+1 == floor {
		return currentPos
	}

	//Drop one down
	var newPoint = Point{currentPos.x, currentPos.y + 1}
	if rocks[newPoint] == false && sands[newPoint] == false {
		return poseSand2(newPoint, rocks, sands, floor)
	}

	//Drop down left
	newPoint = Point{currentPos.x - 1, currentPos.y + 1}
	if rocks[newPoint] == false && sands[newPoint] == false {
		return poseSand2(newPoint, rocks, sands, floor)
	}

	//Drop down right
	newPoint = Point{currentPos.x + 1, currentPos.y + 1}
	if rocks[newPoint] == false && sands[newPoint] == false {
		return poseSand2(newPoint, rocks, sands, floor)
	}

	//Posed in rock or sand
	return currentPos
}

func part1(input []string) int {

	var rocks, abysmStart = loadRocks(input)
	var sands = make(map[Point]bool, 0)

	for {
		var posePos = poseSand(Point{500, 0}, rocks, sands, abysmStart)
		if posePos == nil {
			break
		}

		sands[*posePos] = true
	}

	return len(sands)
}

func part2(input []string) int {

	var rocks, maxY = loadRocks(input)
	var sands = make(map[Point]bool, 0)

	for {
		var posePos = poseSand2(Point{500, 0}, rocks, sands, maxY+2)
		if posePos.x == 500 && posePos.y == 0 {
			sands[posePos] = true
			break
		}

		sands[posePos] = true
	}

	return len(sands)
}

func main() {
	var input = readInput("day14.input")

	fmt.Println("Part1: ", part1(input))
	fmt.Println("Part2: ", part2(input))
}
