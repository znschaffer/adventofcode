package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
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

type monkey struct {
	num            int
	items          []int
	worryFunc      func(old int) int
	testFunc       func(item int) bool
	testNum        int
	testT          int
	testF          int
	itemsInspected int
}

var (
	monkeyNumberRegex    = regexp.MustCompile(`\d+`)
	monkeyItemsRegex     = regexp.MustCompile(`\d+`)
	monkeyWorryFuncRegex = regexp.MustCompile(`[^=]*.$`)
	monkeyTestFuncRegex  = regexp.MustCompile(`\d+`)
	monkeyTestTrueRegex  = regexp.MustCompile(`\d+`)
	monkeyTestFalseRegex = regexp.MustCompile(`\d+`)
)

func parseWorry(d string) func(old int) int {
	opString := monkeyWorryFuncRegex.FindString(d)
	ops := strings.Split(strings.TrimSpace(opString), " ")
	if ops[2] == "old" {
		switch ops[1] {
		case "*":
			return func(old int) int {
				return old * old
			}
		case "+":
			return func(old int) int {
				return old + old
			}
		}
	} else {
		val, _ := strconv.Atoi(ops[2])
		switch ops[1] {
		case "*":
			return func(old int) int {
				return old * val
			}
		case "+":
			return func(old int) int {
				return old + val
			}
		}
	}

	return nil
}

func parseTest(data string) (func(item int) bool, int) {
	num, _ := strconv.Atoi(monkeyTestFuncRegex.FindString(data))
	return func(item int) bool {
		return item%num == 0
	}, num
}

func parseTF(lines []string) (int, int) {
	tNum, _ := strconv.Atoi(monkeyTestTrueRegex.FindString(lines[0]))
	fNum, _ := strconv.Atoi(monkeyTestFalseRegex.FindString(lines[1]))

	return tNum, fNum
}

func newMonkey(data []string) (*monkey, int) {
	m := &monkey{}
	m.num, _ = strconv.Atoi(monkeyNumberRegex.FindString(data[0]))
	itemsString := monkeyItemsRegex.FindAllString(data[1], -1)
	for i := 0; i < len(itemsString); i++ {
		it, _ := strconv.Atoi(itemsString[i])
		m.items = append(m.items, it)
	}

	m.worryFunc = parseWorry(data[2])
	m.testFunc, m.testNum = parseTest(data[3])
	m.testT, m.testF = parseTF(data[4:])
	return m, m.testNum
}

func solve(data []string) (int, int) {
	var (
		part1       int
		part1totals []int
		part2       int
		part2totals []int
		monkeys     []*monkey
		xItem       int
		totalM      = 1
	)

	dlen := len(data)

	for i := 0; i < dlen; i += 7 {
		nm, sm := newMonkey(data[i : i+7])
		totalM *= sm
		monkeys = append(monkeys, nm)
	}

	fmt.Printf("superM: %v\n", totalM)

	for r := 0; r < 20; r++ {
		for i := 0; i < len(monkeys); i++ {
			m := monkeys[i]

			for i := 0; len(m.items) > 0; {
				m.itemsInspected++
				m.items[i] = m.worryFunc(m.items[i])
				m.items[i] /= 3

				if m.items[i]%m.testNum == 0 {
					mNum := m.testT
					xItem, m.items = m.items[0], m.items[1:]
					monkeys[mNum].items = append(monkeys[mNum].items, xItem)
				} else {
					mNum := m.testF
					xItem, m.items = m.items[0], m.items[1:]
					monkeys[mNum].items = append(monkeys[mNum].items, xItem)
				}
			}
		}
	}

	for i := 0; i < len(monkeys); i++ {
		fmt.Printf("Monkey %d inspected items %d times\n", monkeys[i].num, monkeys[i].itemsInspected)
		part1totals = append(part1totals, monkeys[i].itemsInspected)
	}

	sort.Slice(part1totals, func(i, j int) bool {
		return part1totals[i] > part1totals[j]
	})

	part1 = part1totals[0] * part1totals[1]

	// reset
	monkeys = nil
	totalM = 1

	for i := 0; i < dlen; i += 7 {
		nm, sm := newMonkey(data[i : i+7])
		totalM *= sm
		monkeys = append(monkeys, nm)
	}

	for r := 0; r < 10000; r++ {
		for i := 0; i < len(monkeys); i++ {
			m := monkeys[i]

			for i := 0; len(m.items) > 0; {
				m.itemsInspected++
				m.items[i] = m.worryFunc(m.items[i])
				m.items[i] %= totalM

				if m.items[i]%m.testNum == 0 {
					mNum := m.testT
					xItem, m.items = m.items[0], m.items[1:]
					monkeys[mNum].items = append(monkeys[mNum].items, xItem)
				} else {
					mNum := m.testF
					xItem, m.items = m.items[0], m.items[1:]
					monkeys[mNum].items = append(monkeys[mNum].items, xItem)
				}
			}
		}
	}

	for i := 0; i < len(monkeys); i++ {
		fmt.Printf("Monkey %d inspected items %d times\n", monkeys[i].num, monkeys[i].itemsInspected)
		part2totals = append(part2totals, monkeys[i].itemsInspected)
	}

	sort.Slice(part2totals, func(i, j int) bool {
		return part2totals[i] > part2totals[j]
	})

	part2 = part2totals[0] * part2totals[1]

	return part1, part2
}

func main() {
	data := parseInput()
	part1, part2 := solve(data)
	fmt.Printf("Part 1:\t%v\n", part1)
	fmt.Printf("Part 2:\t%v\n", part2)
}
