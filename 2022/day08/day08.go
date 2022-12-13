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

func byteToInt(b byte) int {
	n, _ := strconv.Atoi(string(b))
	return n
}

func solve(data []string) (int, int) {
	var (
		part1     int
		part2     int
		treeCount int
	)

	width := len(data[0])
	height := len(data) - 1

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {

			currTree, err := strconv.Atoi(string(data[y][x]))
			if err != nil {
				fmt.Println(err.Error())
			}

			var (
				ln int
				rn int
				tn int
				bn int
				lh bool
				rh bool
				th bool
				bh bool
			)

			// check if edge
			if x == 0 || x == width-1 || y == 0 || y == height-1 {
				treeCount++
				continue
			}

			// check left
			for i := x - 1; i >= 0; i-- {
				ln++
				if currTree > byteToInt(data[y][i]) {
					lh = true
				} else {
					lh = false
					break
				}
			}

			// check right
			for i := x + 1; i < width; i++ {
				rn++
				if currTree > byteToInt(data[y][i]) {
					rh = true
				} else {
					rh = false
					break
				}
			}

			for i := y - 1; i >= 0; i-- {
				tn++
				if currTree > byteToInt(data[i][x]) {
					th = true
				} else {
					th = false
					break
				}
			}

			// check bottom
			for i := y + 1; i < height; i++ {
				bn++
				if currTree > byteToInt(data[i][x]) {
					bh = true
				} else {
					bh = false
					break
				}
			}

			if lh || rh || bh || th {
				treeCount += 1
			}

			score := (ln * rn * bn * tn)

			if score > part2 {
				part2 = score
			}
		}
	}

	part1 += treeCount
	return part1, part2
}

func main() {
	data := parseInput()
	part1, part2 := solve(data)
	fmt.Printf("Part 1:\t%v\n", part1)
	fmt.Printf("Part 2:\t%v\n", part2)
}
