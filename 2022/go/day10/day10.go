package main

import (
	"fmt"
	"os"
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

func printScreen(screen [240]rune) {
	for x := 0; x < 240; x++ {
		if x%40 == 0 {
			fmt.Printf("\n")
		}
		fmt.Printf("%c", screen[x])
	}
	fmt.Printf("\n")
}

func parse(line string) (string, int) {
	var in string
	var val int
	if len(line) == 0 {
		return "", 0
	}
	parts := strings.Split(line, " ")
	in = parts[0]
	if in == "noop" {
		return "noop", 0
	} else {
		val, _ = strconv.Atoi(parts[1])
		return "addx", val
	}
}

func drawPixel(screen [240]rune, cycle int, x int) [240]rune {
	if x == (cycle%40) || x-1 == (cycle%40) || x+1 == (cycle%40) {
		screen[cycle] = '▉'
	} else {
		screen[cycle] = ' '
	}

	return screen
}

func solve(data []string) (int, [240]rune) {
	var (
		screen     [240]rune
		x          = 1
		cyclesLeft int
		i          int
		val        int
		strength   int
		in         string
	)

	for cycle := 0; cycle < 240; cycle++ {
		if (cycle == 20) || ((cycle-20)%40) == 0 {
			strength += (cycle * x)
		}
		if cyclesLeft == 0 {
			x += val
			val = 0
			in, val = parse(data[i])
			i++
			switch in {
			case "noop":
				screen = drawPixel(screen, cycle, x)
			case "addx":
				screen = drawPixel(screen, cycle, x)
				cyclesLeft = 1
			}
			continue
		}
		screen = drawPixel(screen, cycle, x)
		cyclesLeft -= 1
	}

	return strength, screen
}

func main() {
	data := parseInput()
	part1, screen := solve(data)
	fmt.Printf("Part 1:\t%v\n", part1)
	fmt.Printf("Part 2:\n")
	printScreen(screen)
}
