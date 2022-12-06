package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseInput() []string {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	return strings.Split(string(data), "\n")
}

func solve(input []string) (int, int) {
	part1 := 0
	part2 := 0
	for _, pair := range input {
		sections := strings.Split(pair, ",")
		left := strings.Split(sections[0], "-")
		right := strings.Split(sections[1], "-")

		var leftNums, rightNums []int

		for _, nums := range left {
			n, _ := strconv.Atoi(nums)
			leftNums = append(leftNums, n)
		}

		for _, nums := range right {
			n, _ := strconv.Atoi(nums)
			rightNums = append(rightNums, n)
		}

		if (leftNums[0] <= rightNums[0] && leftNums[1] >= rightNums[1]) || (rightNums[0] <= leftNums[0] && rightNums[1] >= leftNums[1]) {
			part1++
		}

		if (leftNums[0] <= rightNums[0] && leftNums[1] >= rightNums[0]) || (rightNums[0] <= leftNums[0] && rightNums[1] >= leftNums[0]) {
			part2++
		}
	}
	return part1, part2
}

func main() {
	data := parseInput()
	part1, part2 := solve(data)
	fmt.Printf("Part 1:\t%v\n", part1)
	fmt.Printf("Part 2:\t%v\n", part2)
}
