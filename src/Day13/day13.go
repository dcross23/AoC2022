package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Tree struct {
	info     int
	children []*Tree
	parent   *Tree
}

func parseTree(treeStr string) Tree {
	var tree = Tree{-1, []*Tree{}, nil}
	var current = &tree
	var number string

	for _, c := range treeStr {
		if c == '[' {
			var newTree = Tree{-1, []*Tree{}, current}
			current.children = append(current.children, &newTree)
			current = &newTree

		} else if c == ',' || c == ']' {
			if len(number) > 0 {
				var numberInt, _ = strconv.Atoi(number)
				current.info = numberInt
				number = ""
			}
			current = current.parent

			if c == ',' {
				var newTree = Tree{-1, []*Tree{}, current}
				current.children = append(current.children, &newTree)
				current = &newTree
			}

		} else {
			number += string(c)
		}
	}

	return tree
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

func printTree(tree *Tree, level int) {
	fmt.Println("Level "+strconv.Itoa(level)+":", tree.info)

	for _, node := range tree.children {
		printTree(node, level+1)
	}
}

func compareTrees(tree1, tree2 Tree) int {

	//No childs, just a leaf
	if len(tree1.children) == 0 && len(tree2.children) == 0 {
		if tree1.info > tree2.info {
			return -1
		} else if tree1.info < tree2.info {
			return 1
		} else {
			return 0
		}
	}

	//Reached leafs
	if tree1.info >= 0 {
		return compareTrees(Tree{-1, []*Tree{&tree1}, nil}, tree2)
	}

	if tree2.info >= 0 {
		return compareTrees(tree1, Tree{-1, []*Tree{&tree2}, nil})
	}

	//Compare trees if they are notleafs
	var i int
	for i = 0; i < len(tree1.children) && i < len(tree2.children); i++ {
		var comparedValue = compareTrees(*tree1.children[i], *tree2.children[i])
		if comparedValue != 0 {
			return comparedValue
		}
	}

	if i < len(tree1.children) {
		return -1
	} else if i < len(tree2.children) {
		return 1
	} else {
		return 0
	}
}

func part1(input []string) int {
	var sum = 0
	var pairNum = 1

	for i := 0; i < len(input); i += 3 {
		var pair1 = input[i]
		var pair2 = input[i+1]

		var tree1 = parseTree(pair1)
		var tree2 = parseTree(pair2)

		var comparedValue = compareTrees(tree1, tree2)
		if comparedValue == 1 {
			sum += pairNum
		}

		pairNum++
	}

	return sum
}

func part2(input []string) int {
	var trees []Tree

	for i := 0; i < len(input); i++ {
		if input[i] == "" {
			continue
		}

		var treeStr = input[i]
		var tree = parseTree(treeStr)
		trees = append(trees, tree)
	}

	var treeNew1 = parseTree("[[2]]")
	var treeNew2 = parseTree("[[6]]")

	trees = append(trees, treeNew1)
	trees = append(trees, treeNew2)

	sort.Slice(trees, func(i, j int) bool {
		return compareTrees(trees[i], trees[j]) == 1
	})

	var key = 1
	for i := 0; i < len(trees); i++ {
		if compareTrees(treeNew1, trees[i]) == 0 || compareTrees(treeNew2, trees[i]) == 0 {
			key *= (i + 1)
		}

	}

	return key
}
func main() {
	var input = readInput("day13.input")

	fmt.Println("Part1: ", part1(input))
	fmt.Println("Part2: ", part2(input))
}
