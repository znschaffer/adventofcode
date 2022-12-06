package main

import (
	"fmt"
	"math/bits"
	"os"
	"strings"
)

type bitSet struct {
	bits uint64
}

func (c *bitSet) setRune(i rune) {
	c.bits |= 1 << (i - 'A')
}

func setString(s string) bitSet {
	var cSet bitSet
	for _, rune := range s {
		cSet.setRune(rune)
	}
	return cSet
}

func formatItem(i int) int {
	if i > 96 {
		i -= 96
	} else {
		i -= 38
	}
	return i
}

func compareOneString(s string) int {
	splitStringLen := len(s) / 2
	l, r := setString(s[:splitStringLen]), setString(s[splitStringLen:])
	item := bits.TrailingZeros64(l.bits&r.bits) + 'A'
	return formatItem(item)
}

func compareThreeStrings(s1 string, s2 string, s3 string) int {
	s1Bits, s2Bits, s3Bits := setString(s1), setString(s2), setString(s3)
	item := bits.TrailingZeros64(s1Bits.bits&s2Bits.bits&s3Bits.bits) + 'A'
	return formatItem(item)
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
	var lines []string
	for _, line := range input {
		part1 += compareOneString(line)
		lines = append(lines, line)
		numLines = (numLines + 1) % 3
		if numLines == 0 && lines != nil {
			badgeNumber := compareThreeStrings(lines[0], lines[1], lines[2])
			part2 += badgeNumber
			lines = nil
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
