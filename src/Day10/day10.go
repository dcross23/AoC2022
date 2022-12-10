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
	var checkCycle = 20
	var cycle = 0
	var X = 1
	var sum = 0

	for _, inst := range input {
		var opc = strings.Split(inst, " ")[0]

		switch opc {
		case "noop":
			cycle++
			if cycle == checkCycle {
				sum += (X * cycle)
				checkCycle += 40
			}
			break

		case "addx":
			var value, _ = strconv.Atoi(strings.Split(inst, " ")[1])

			for i := 0; i < 2; i++ {
				cycle++
				if cycle == checkCycle {
					sum += (X * cycle)
					checkCycle += 40
				}
			}

			X += value
			break

		}
	}

	return sum
}

func part2(input []string) [][]string {
	var cycle = 0
	var X = 1

	var drawing = make([][]string, 6)
	for i := range drawing {
		drawing[i] = make([]string, 40)
	}

	for _, inst := range input {
		var opc = strings.Split(inst, " ")[0]

		switch opc {
		case "noop":
			var row = cycle / 40
			var column = cycle % 40

			if column >= (X-1) && column <= (X+1) {
				drawing[row][column] = "#"
			} else {
				drawing[row][column] = " "
			}

			cycle++
			break

		case "addx":
			var value, _ = strconv.Atoi(strings.Split(inst, " ")[1])

			for i := 0; i < 2; i++ {
				var row = cycle / 40
				var column = cycle % 40

				if column >= (X-1) && column <= (X+1) {
					drawing[row][column] = "#"
				} else {
					drawing[row][column] = " "
				}

				cycle++
			}

			X += value
			break

		}
	}

	return drawing
}

func main() {
	var input = readInput("day10.input")

	fmt.Println("Part1: ", part1(input))

	fmt.Println("Part2: ")
	var drawing = part2(input)
	for i := 0; i < 6; i++ {
		fmt.Println(drawing[i])
	}
}
