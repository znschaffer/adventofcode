package main

import (
	"fmt"
	"math/bits"
	"os"
	"strings"
	"time"
)

type bitSet struct {
	bits uint64
}

func (c *bitSet) setRune(i rune) {
	c.bits |= 1 << (i - 'A')
}

func strToBits(s string) bitSet {
	var cSet bitSet
	for _, rune := range s {
		cSet.setRune(rune)
	}
	return cSet
}

func formatLetter(i int) int {
	if i > 96 {
		i -= 96
	} else {
		i -= 38
	}
	return i
}

func compareLine(s string) int {
	halfLen := len(s) / 2

	l, r := strToBits(s[:halfLen]), strToBits(s[halfLen:])

	letter := bits.TrailingZeros64(l.bits&r.bits) + 'A'

	return formatLetter(letter)
}

func compareGroup(group []string) int {
	s1, s2, s3 := strToBits(group[0]), strToBits(group[1]), strToBits(group[2])

	letter := bits.TrailingZeros64(s1.bits&s2.bits&s3.bits) + 'A'

	return formatLetter(letter)
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
	var numLines int
	var group []string

	for _, line := range input {
		part1 += compareLine(line)

		group = append(group, line)
		numLines = (numLines + 1) % 3

		if numLines == 0 && group != nil {
			part2 += compareGroup(group)
			group = nil
		}
	}
	return part1, part2
}

func main() {
	start := time.Now()
	data := parseInput()
	solveStart := time.Now()
	part1, part2 := solve(data)
	solveTime := time.Since(solveStart)

	fmt.Printf("Part 1:\t%v\n", part1)
	fmt.Printf("Part 2:\t%v\n", part2)
	fmt.Printf("Solve Time:\t%s\n", solveTime)
	fmt.Printf("Run Time:\t%s\n", time.Since(start))
}
