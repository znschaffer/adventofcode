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

func solve[T any](data []string) (T, T) {
	var part1 T
	var part2 T

	return part1, part2
}

func main() {
	data := parseInput()
	part1, part2 := solve[int](data)
	fmt.Printf("Part 1:\t%v", part1)
	fmt.Printf("Part 2:\t%v", part2)
}
