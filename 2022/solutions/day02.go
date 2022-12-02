package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func checkA(in string, part string) int {
	switch in {
	case "X":
		switch part {
		case "a":
			return (1 + 3)
		case "b":
			return (3 + 0)
		}
	case "Y":
		switch part {
		case "a":
			return (2 + 6)
		case "b":
			return (1 + 3)
		}
	case "Z":
		switch part {
		case "a":
			return (3 + 0)
		case "b":
			return (2 + 6)
		}
	}
	return 0
}

func checkB(in string, part string) int {
	switch in {
	case "X":
		switch part {
		case "a":
			return (1 + 0)
		case "b":
			return (1 + 0)
		}
	case "Y":
		switch part {
		case "a":
			return (2 + 3)
		case "b":
			return (2 + 3)
		}
	case "Z":
		switch part {
		case "a":
			return (3 + 6)
		case "b":
			return (3 + 6)
		}
	}
	return 0
}

func checkC(in string, part string) int {
	switch in {
	case "X":
		switch part {
		case "a":
			return (1 + 6)
		case "b":
			return (2 + 0)
		}
	case "Y":
		switch part {
		case "a":
			return (2 + 0)
		case "b":
			return (3 + 3)
		}
	case "Z":
		switch part {
		case "a":
			return (3 + 3)
		case "b":
			return (1 + 6)
		}
	}
	return 0
}

func main() {
	part := flag.String("part", "a", "part of aoc day")
	flag.Parse()
	f, err := os.Open("../input/day02.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var totalScore int
	s := bufio.NewScanner(f)
	for s.Scan() {
		round := strings.Split(s.Text(), " ")
		switch round[0] {
		case "A":
			totalScore += checkA(round[1], *part)
		case "B":
			totalScore += checkB(round[1], *part)
		case "C":
			totalScore += checkC(round[1], *part)
		}
	}

	fmt.Println(totalScore)
}
