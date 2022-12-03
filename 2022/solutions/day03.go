package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/bits"
	"os"
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

func main() {
	part := flag.String("part", "a", "part of aoc problem")
	flag.Parse()
	f, err := os.Open("../input/day03.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	var sum int
	var numLines int
	var lines []string

	for s.Scan() {
		switch *part {
		case "a":
			sum += compareOneString(s.Text())
		case "b":
			lines = append(lines, s.Text())
			numLines = (numLines + 1) % 3
			if numLines == 0 && lines != nil {
				badgeNumber := compareThreeStrings(lines[0], lines[1], lines[2])
				sum += badgeNumber
				lines = nil
			}
		}
	}

	fmt.Printf("%d\n", sum)
}
