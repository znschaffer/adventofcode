package main

import (
	"fmt"
	"os"
	"strings"
)

var scores = map[string]map[string][]int{
	"A": {
		"X": []int{4, 3},
		"Y": []int{8, 4},
		"Z": []int{3, 8},
	},
	"B": {
		"X": []int{1, 1},
		"Y": []int{5, 5},
		"Z": []int{9, 9},
	},
	"C": {
		"X": []int{7, 2},
		"Y": []int{2, 6},
		"Z": []int{6, 7},
	},
}

func parseInput() []string {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return strings.Split(string(data), "\n")
}

func solve(input []string) (int, int) {
	part1 := 0
	part2 := 0

	for _, line := range input {
		round := strings.Split(line, " ")
		results := scores[round[0]][round[1]]
		part1 += results[0]
		part2 += results[1]
	}
	return part1, part2
}

func main() {
	data := parseInput()
	part1, part2 := solve(data)
	fmt.Printf("Part 1:\t%v\n", part1)
	fmt.Printf("Part 2:\t%v\n", part2)
}
