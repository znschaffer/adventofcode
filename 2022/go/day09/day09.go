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

type rope []pos

type pos struct {
	X, Y int
}

func printGrid(r rope) {
	rlen := len(r)
	for y := 0; y < 5; y++ {
		for x := 0; x < 6; x++ {
			marker := "."
			for rl := 0; rl < rlen; rl++ {
				if r[rl].X == x && r[rl].Y == y {
					if rl == 0 {
						marker = "H"
					} else {
						marker = fmt.Sprintf("%d", rl)
					}
					break
				}
			}
			fmt.Printf("%s", marker)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func moveRope(r rope, data []string) int {
	rlen := len(r) - 1
	m := make(map[pos]bool)
	m[r[rlen]] = true

	for _, line := range data {
		if len(line) == 0 {
			continue
		}
		dir, strN := line[0], line[2:]
		n, err := strconv.Atoi(strN)
		if err != nil {
			fmt.Println(err.Error())
		}

		// fmt.Printf("== %c %d ==\n", dir, n)
		for i := 0; i < n; i++ {

			switch dir {
			case 'R':
				r[0].X++
			case 'L':
				r[0].X--
			case 'D':
				r[0].Y++
			case 'U':
				r[0].Y--
			}

			for i := 1; i < len(r); i++ {
				if (r[i-1].X == r[i].X+1) && (r[i-1].Y == r[i].Y+1) {
					continue
				} else if (r[i-1].X == r[i].X+1) && (r[i-1].Y == r[i].Y-1) {
					continue
				} else if (r[i-1].X == r[i].X-1) && (r[i-1].Y == r[i].Y-1) {
					continue
				} else if (r[i-1].X == r[i].X-1) && (r[i-1].Y == r[i].Y+1) {
					continue
				} else if (r[i-1].X == r[i].X+2) && (r[i-1].Y == r[i].Y) {
					r[i].X++
				} else if (r[i-1].X == r[i].X-2) && (r[i-1].Y == r[i].Y) {
					r[i].X--
				} else if (r[i-1].Y == r[i].Y+2) && (r[i-1].X == r[i].X) {
					r[i].Y++
				} else if (r[i-1].Y == r[i].Y-2) && (r[i-1].X == r[i].X) {
					r[i].Y--
				} else if (r[i-1].X > r[i].X) && (r[i-1].Y > r[i].Y) {
					r[i].X++
					r[i].Y++
				} else if (r[i-1].X < r[i].X) && (r[i-1].Y < r[i].Y) {
					r[i].X--
					r[i].Y--
				} else if (r[i-1].X > r[i].X) && (r[i-1].Y < r[i].Y) {
					r[i].X++
					r[i].Y--
				} else if (r[i-1].X < r[i].X) && (r[i-1].Y > r[i].Y) {
					r[i].X--
					r[i].Y++
				}
				// printGrid(r)
			}

			_, ok := m[r[rlen]]
			if !ok {
				m[r[rlen]] = true
			}

		}

	}

	return len(m)
}

func solve(data []string) (int, int) {
	p1rope := rope{
		pos{0, 4},
		pos{0, 4},
	}

	p2rope := rope{
		pos{0, 4},
		pos{0, 4},
		pos{0, 4},
		pos{0, 4},
		pos{0, 4},
		pos{0, 4},
		pos{0, 4},
		pos{0, 4},
		pos{0, 4},
		pos{0, 4},
	}

	part1 := moveRope(p1rope, data)
	part2 := moveRope(p2rope, data)

	return part1, part2
}

func main() {
	data := parseInput()
	part1, part2 := solve(data)
	fmt.Printf("Part 1:\t%v\n", part1)
	fmt.Printf("Part 2:\t%v\n", part2)
}
