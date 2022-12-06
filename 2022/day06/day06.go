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
	return strings.Split(string(data), "")
}

func checkUnique(s string) bool {
	var a [256]bool
	for _, ascii := range s {
		if a[ascii] {
			return false
		}
		a[ascii] = true
	}
	return true
}

func solve(input []string) (int, int) {
	part1 := 1
	part2 := 1

	for idx := range input {
		if idx >= 13 {
			var line string
			for i := idx; i > (idx - 14); i-- {
				line += input[i]
			}
			if checkUnique(line) {
				break
			}
			part2++
			continue
		}
		part2++
		continue
	}

	for idx := range input {
		if idx >= 3 {
			var line string
			for i := idx; i > (idx - 4); i-- {
				line += input[i]
			}
			if checkUnique(line) {
				break
			}
			part1++
			continue
		}
		part1++
		continue
	}

	return part1, part2
}

func main() {
	data := parseInput()
	part1, part2 := solve(data)
	fmt.Printf("Part 1:\t%d\n", part1)
	fmt.Printf("Part 2:\t%d\n", part2)
}
