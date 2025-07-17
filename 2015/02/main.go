package main

import (
	"2015/utils"
	"fmt"
	"strconv"
	"strings"
)

func solve(input []string) (int, int) {
	wrappingPaperArea := 0
	ribbonLength := 0

	for _, dimensions := range input {
		split := strings.Split(dimensions, "x")
		l, _ := strconv.Atoi(split[0])
		w, _ := strconv.Atoi(split[1])
		h, _ := strconv.Atoi(split[2])

		totalSurfaceArea := 2 * (l*w + w*h + l*h)
		smallestSideArea := l * w * h / max(l, w, h)
		smallestPerimeter := 2 * (l + w + h - max(l, w, h))
		totalVolume := l * w * h

		wrappingPaperArea += totalSurfaceArea + smallestSideArea
		ribbonLength += smallestPerimeter + totalVolume
	}

	return wrappingPaperArea, ribbonLength
}

func main() {
	input := utils.ReadInput(2)
	part1, part2 := solve(input)

	fmt.Println("-- Day 02 --")
	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
