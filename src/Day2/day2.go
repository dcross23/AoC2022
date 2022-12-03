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

	var plays = make([]string, 0)
	for fileScanner.Scan() {
		var text = fileScanner.Text()
		plays = append(plays, text)
	}

	readFile.Close()
	return plays
}

func part1(plays []string) int {
	var points = 0
	var possiblePlays = make(map[string][]int)
	//Points acording to possible plays for each enemy play
	possiblePlays["A"] = []int{3, 6, 0}
	possiblePlays["B"] = []int{0, 3, 6}
	possiblePlays["C"] = []int{6, 0, 3}

	for _, s := range plays {
		var play = strings.Split(s, " ")
		var enemy = play[0]
		var response = play[1][0]

		var playPoints = possiblePlays[enemy][response-'X'] + int(response-'X'+1)
		points += playPoints

	}

	return points
}

// X lose    Y draw   Z win
func part2(plays []string) int {
	var points = 0
	var possiblePlays = make(map[string]int)
	//Points acording to how the round must end
	possiblePlays["X"] = 0
	possiblePlays["Y"] = 3
	possiblePlays["Z"] = 6

	//Points acording to my play
	var pointsForMyPlay = make(map[string][]int)
	pointsForMyPlay["A"] = []int{3, 1, 2}
	pointsForMyPlay["B"] = []int{1, 2, 3}
	pointsForMyPlay["C"] = []int{2, 3, 1}

	for _, s := range plays {
		var play = strings.Split(s, " ")
		var enemy = play[0]
		var response = play[1]

		var playPoints = int(possiblePlays[response]) + pointsForMyPlay[enemy][response[0]-'X']
		points += playPoints
	}

	return points
}

// enemy     A rock       B paper     C scissors
// response  X rock	     Y paper     Z scissors
// round points -> 1 rock, 2 paper, 3 scissors   +  6 win, 3 draw, 0 lose
func main() {
	var plays = readInput("day2.input")

	fmt.Println("Part1: " + strconv.Itoa(part1(plays)))
	fmt.Println("Part2: " + strconv.Itoa(part2(plays)))
}
