package main

import (
	"advent/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const gridSize = 1000

var commandRegex = regexp.MustCompile(`(\d*),(\d*) through (\d*),(\d*)$`)

type coords struct {
	X int
	Y int
}

type CommandType int

const (
	TurnOn CommandType = iota
	TurnOff
	Toggle
	Default
)

func getCommandType(command string) CommandType {
	if strings.HasPrefix(command, "turn on") {
		return TurnOn
	} else if strings.HasPrefix(command, "turn off") {
		return TurnOff
	} else if strings.HasPrefix(command, "toggle") {
		return Toggle
	}
	return Default
}

func apply(lights *[][gridSize]bool, dynamicLights *[][gridSize]int, start coords, end coords, commandType CommandType) {
	switch commandType {
	case TurnOn:
		for x := start.X; x <= end.X; x++ {
			for y := start.Y; y <= end.Y; y++ {
				(*lights)[x][y] = true
				(*dynamicLights)[x][y]++
			}
		}
	case TurnOff:
		for x := start.X; x <= end.X; x++ {
			for y := start.Y; y <= end.Y; y++ {
				(*lights)[x][y] = false
				if (*dynamicLights)[x][y] > 0 {
					(*dynamicLights)[x][y]--
				}
			}
		}
	case Toggle:
		for x := start.X; x <= end.X; x++ {
			for y := start.Y; y <= end.Y; y++ {
				(*lights)[x][y] = !(*lights)[x][y]
				(*dynamicLights)[x][y] += 2
			}
		}
	}
}

func countLightsOn(lights [][gridSize]bool, dynamicLights [][gridSize]int) (int, int) {
	count, brightness := 0, 0
	for x := 0; x < len(lights); x++ {
		for y := 0; y < len(lights[x]); y++ {
			if lights[x][y] {
				count++
			}
			brightness += dynamicLights[x][y]
		}
	}
	return count, brightness
}

func solve(input []string) (int, int) {
	part1 := 0
	part2 := 0

	lights := make([][gridSize]bool, gridSize)
	dynamicLights := make([][gridSize]int, gridSize)

	for _, command := range input {
		start, end := coords{}, coords{}

		matches := commandRegex.FindStringSubmatch(command)
		start.X, _ = strconv.Atoi(matches[1])
		start.Y, _ = strconv.Atoi(matches[2])
		end.X, _ = strconv.Atoi(matches[3])
		end.Y, _ = strconv.Atoi(matches[4])

		commandType := getCommandType(command)
		apply(&lights, &dynamicLights, start, end, commandType)
	}

	part1, part2 = countLightsOn(lights, dynamicLights)
	return part1, part2
}

func main() {
	input := utils.ReadInput(6)
	part1, part2 := solve(input)

	fmt.Println("-- Day 06 --")
	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
