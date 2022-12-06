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

func solve(input []string) ([]byte, []byte) {

    // hardcoded input like a wimp
	stacks := [][]byte{
		{'H', 'R', 'B', 'D', 'Z', 'F', 'L', 'S'},
		{'T', 'B', 'M', 'Z', 'R'},
		{'Z', 'L', 'C', 'H', 'N', 'S'},
		{'S', 'C', 'F', 'J'},
		{'P', 'G', 'H', 'W', 'R', 'Z', 'B'},
		{'V', 'J', 'Z', 'G', 'D', 'N', 'M', 'T'},
		{'G', 'L', 'N', 'W', 'F', 'S', 'P', 'Q'},
		{'M', 'Z', 'R'},
		{'M', 'C', 'L', 'G', 'V', 'R', 'T'},
	}

	p1stacks := append([][]byte(nil), stacks...)
	p2stacks := append([][]byte(nil), stacks...)

	var part1 []byte
	var part2 []byte

	for _, line := range input {
		var p1x byte
		var p2x []byte
		inst := strings.Split(line, " ")
		if inst[0] != "move" {
			continue
		}

		amt, _ := strconv.Atoi(inst[1])
		src, _ := strconv.Atoi(inst[3])
		dest, _ := strconv.Atoi(inst[5])
		src--
		dest--

        // part 1 instructions
		for i := 0; i < amt; i++ {
			p1x, p1stacks[src] = p1stacks[src][len(p1stacks[src])-1], p1stacks[src][:(len(p1stacks[src])-1)]
			p1stacks[dest] = append(p1stacks[dest], p1x)
		}

        // part 2 instructions
		p2x, p2stacks[src] = p2stacks[src][(len(p2stacks[src])-amt):], p2stacks[src][:(len(p2stacks[src])-amt)]
		p2stacks[dest] = append(p2stacks[dest], p2x...)
	}

	for _, stack := range p1stacks {
		part1 = append(part1, stack[len(stack)-1])
	}

	for _, stack := range p2stacks {
		part2 = append(part2, stack[len(stack)-1])
	}

	return part1, part2
}

func main() {
	data := parseInput()
	part1, part2 := solve(data)
	fmt.Printf("Part 1:\t%c\n", part1)
	fmt.Printf("Part 2:\t%c\n", part2)
}
