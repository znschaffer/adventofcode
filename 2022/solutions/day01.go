package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var totals []int
	var totalLen int
	var currNum int

	f, err := os.Open("../input/day01.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		if len(s.Bytes()) == 0 {
			totals = append(totals, currNum)
			totalLen++
			currNum = 0
			continue
		}
		if num, err := strconv.Atoi(s.Text()); err == nil {
			currNum += num
		}
	}

	sort.Ints(totals)

	fmt.Printf("Largest Total: %d\n", totals[totalLen-1])
	fmt.Printf("Combined Three Largest: %d\n", (totals[totalLen-1] + totals[totalLen-2] + totals[totalLen-3]))
}
