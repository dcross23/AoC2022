package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type File struct {
	Name string
	Size int
}

type Folder struct {
	Name    string
	Files   []File
	Folders []Folder
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

func part1(input []string) (int, map[string]int) {
	var sizes = make(map[string]int)
	var path = make([]string, 0)

	for i := 0; i < len(input); i++ {
		var line = input[i]

		if strings.HasPrefix(line, "$") {
			var command = strings.Split(line[2:], " ")[0]

			switch command {
			case "cd":
				var param = strings.Split(line[2:], " ")[1]

				if param == ".." {
					path = path[:len(path)-1]
				} else {
					path = append(path, param)
				}
				break
			case "ls":
				break

			}

		} else {
			var sizeOrDir = strings.Split(line, " ")[0]

			if strings.HasPrefix(sizeOrDir, "dir") {
				continue
			}

			var fileSize, _ = strconv.Atoi(sizeOrDir)
			for j := 0; j <= len(path); j++ {
				var fullPath = strings.Join(path[:j], "/")
				sizes[fullPath] += fileSize
			}
		}
	}

	var result = 0
	for _, size := range sizes {
		if size <= 100000 {
			result += size
		}
	}

	return result, sizes
}

func part2(sizes map[string]int) int {
	var spaceToFree = sizes["/"] - (70000000 - 30000000)

	var result = math.MaxInt
	for _, size := range sizes {
		if size >= spaceToFree {
			if size < result {
				result = size
			}
		}
	}
	return result
}

func main() {
	var input = readInput("day7.input")

	var result, sizes = part1(input)
	fmt.Println("Part1: ", result)

	fmt.Println("Part2: ", part2(sizes))
}
