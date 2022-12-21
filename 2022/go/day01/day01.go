package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
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
	var totals []int
	var totalLen int
	var currNum int

	for _, line := range input {
		if len(line) == 0 {
			totals = append(totals, currNum)
			totalLen++
			currNum = 0
			continue
		}
		if num, err := strconv.Atoi(line); err == nil {
			currNum += num
		}
	}
	sort.Ints(totals)

	part1 := totals[totalLen-1]
	part2 := totals[totalLen-1] + totals[totalLen-2] + totals[totalLen-3]
	return part1, part2
}

func main() {
	data := parseInput()
	part1, part2 := solve(data)
	fmt.Printf("Part 1:\t%v\n", part1)
	fmt.Printf("Part 2:\t%v\n", part2)
}
