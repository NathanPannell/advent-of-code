package main

import (
	"advent/utils"
	"fmt"
	"regexp"
	//"github.com/glenn-brown/golang-pkg-pcre/src/pkg/pcre"
)

func solve(input []string) (int, int) {
	part1 := 0
	part2 := 0

	threeVowels := regexp.MustCompile(`([aeiou].*){3}`)
	bannedPair := regexp.MustCompile(`ab|cd|pq|xy`)

	// `(.)\1`
	repeatedChar := func(s string) bool {
		for i := 0; i < len(s)-1; i++ {
			if s[i] == s[i+1] {
				return true
			}
		}
		return false
	}

	// `(.).\1`
	repeatedCharSpaced := func(s string) bool {
		for i := 0; i < len(s)-2; i++ {
			if s[i] == s[i+2] {
				return true
			}
		}
		return false
	}

	// `(..).*\1`
	pair := func(s string) bool {
		for i := 0; i < len(s)-3; i++ {
			for j := i + 2; j < len(s)-1; j++ {
				if s[i:i+2] == s[j:j+2] {
					return true
				}
			}
		}
		return false
	}

	for _, s := range input {
		if threeVowels.MatchString(s) &&
			repeatedChar(s) &&
			!bannedPair.MatchString(s) {
			part1++
		}

		if pair(s) &&
			repeatedCharSpaced(s) {
			part2++
		}
	}

	return part1, part2
}

func main() {
	input := utils.ReadInput(5)
	part1, part2 := solve(input)

	fmt.Println("-- Day 05 --")
	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
