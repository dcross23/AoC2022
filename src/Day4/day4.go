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

func part1(sections []string) int {
	var overlaps = 0

	for _, pair := range sections {
		var sections = strings.Split(pair, ",")

		var x1, _ = strconv.Atoi(strings.Split(sections[0], "-")[0])
		var y1, _ = strconv.Atoi(strings.Split(sections[0], "-")[1])

		var x2, _ = strconv.Atoi(strings.Split(sections[1], "-")[0])
		var y2, _ = strconv.Atoi(strings.Split(sections[1], "-")[1])

		if x1 <= x2 && y1 >= y2 || x1 >= x2 && y1 <= y2 {
			overlaps++
		}
	}

	return overlaps
}

func part2(sections []string) int {
	var overlaps = 0

	for _, pair := range sections {
		var sections = strings.Split(pair, ",")

		var x1, _ = strconv.Atoi(strings.Split(sections[0], "-")[0])
		var y1, _ = strconv.Atoi(strings.Split(sections[0], "-")[1])

		var x2, _ = strconv.Atoi(strings.Split(sections[1], "-")[0])
		var y2, _ = strconv.Atoi(strings.Split(sections[1], "-")[1])

		if y1 < x2 || x1 > y2 {
			continue
		}

		overlaps++
	}

	return overlaps
}

func main() {
	var sections = readInput("day4.input")

	fmt.Println("Part1: ", part1(sections))
	fmt.Println("Part2: ", part2(sections))
}
