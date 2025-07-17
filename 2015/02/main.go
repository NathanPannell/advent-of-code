package main

import (
	"2015/utils"
	"fmt"
	"strconv"
	"strings"
)

func part1(input []string) int {
	total := 0
	for _, dimensions := range input {
		split := strings.Split(dimensions, "x")

		l, _ := strconv.Atoi(split[0])
		w, _ := strconv.Atoi(split[1])
		h, _ := strconv.Atoi(split[2])

		leastSide := l * w * h / max(l, w, h)
		total += leastSide + 2*(l*w+w*h+l*h)
	}
	return total
}

func part2(input []string) int {
	return 1
}

func main() {
	input := utils.ReadInput(2)

	fmt.Println("-- Day 02 --")
	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}
