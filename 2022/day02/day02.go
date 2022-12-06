package main

import (
	"fmt"
	"os"
	"strings"
)

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
		switch round[0] {
		case "A":
			part1result, part2result := checkA(round[1])
			part1 += part1result
			part2 += part2result
		case "B":
			part1result, part2result := checkB(round[1])
			part1 += part1result
			part2 += part2result
		case "C":
			part1result, part2result := checkC(round[1])
			part1 += part1result
			part2 += part2result
		}
	}

	return part1, part2
}

func checkA(in string) (int, int) {
	switch in {
	case "X":
		return (1 + 3), (3 + 0)
	case "Y":
		return (2 + 6), (1 + 3)
	case "Z":
		return (3 + 0), (2 + 6)
	}
	return 0, 0
}

func checkB(in string) (int, int) {
	switch in {
	case "X":
		return (1 + 0), (1 + 0)
	case "Y":
		return (2 + 3), (2 + 3)
	case "Z":
		return (3 + 6), (3 + 6)
	}
	return 0, 0
}

func checkC(in string) (int, int) {
	switch in {
	case "X":
		return (1 + 6), (2 + 0)
	case "Y":
		return (2 + 0), (3 + 3)
	case "Z":
		return (3 + 3), (1 + 6)
	}
	return 0, 0
}

func main() {
	data := parseInput()
	part1, part2 := solve(data)
	fmt.Printf("Part 1:\t%v\n", part1)
	fmt.Printf("Part 2:\t%v\n", part2)
}
