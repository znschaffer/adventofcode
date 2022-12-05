package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func pairsContain(line string, part string) bool {
	lines := strings.Split(line, ",")
	left := strings.Split(lines[0], "-")
	right := strings.Split(lines[1], "-")

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
		return true
	}

	if part == "b" {
		if (leftNums[0] <= rightNums[0] && leftNums[1] >= rightNums[0]) || (rightNums[0] <= leftNums[0] && rightNums[1] >= leftNums[0]) {
			return true
		}
	}
	return false
}

func readInput(part string) int {
	f, err := os.Open("../input/day04.txt")
	if err != nil {
		panic("can't open file")
	}

	defer f.Close()

	s := bufio.NewScanner(f)

	var total int

	for s.Scan() {
		if pairsContain(s.Text(), part) {
			total++
		}
	}
	return total
}

func main() {
	part := flag.String("part", "a", "part of aoc problem")
	flag.Parse()

	fmt.Println(readInput(*part))
}
