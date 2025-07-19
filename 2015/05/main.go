package main

import (
	"advent/utils"
	"fmt"
	//"github.com/glenn-brown/golang-pkg-pcre/src/pkg/pcre"
)

func solve(input []string) (int, int) {
	part1 := 0
	part2 := 0

	// Using a Perl-Compatible Regex to support backreferences `\1`
	//threeVowels := pcre.MustCompile(`([aeiou].*){3}`, 0)
	//repeatedChar := pcre.MustCompile(`(.)\1`, 0)
	//bannedPair := pcre.MustCompile(`ab|cd|pq|xy`, 0)
	//repeatedPair := pcre.MustCompile(`(..)\1`, 0)
	//repeatedCharSeparated := pcre.MustCompile(`(.).\1`, 0)

	//for _, s := range input {
	//	if threeVowels.MatcherString(s, 0).Matches() &&
	//		repeatedChar.MatcherString(s, 0).Matches() &&
	//		!bannedPair.MatcherString(s, 0).Matches() {
	//		part1++
	//	}
	//
	//	if repeatedPair.MatcherString(s, 0).Matches() &&
	//		repeatedCharSeparated.MatcherString(s, 0).Matches() {
	//		part2++
	//	}
	//}

	return part1, part2
}

func main() {
	input := utils.ReadInput(5)
	part1, part2 := solve(input)

	fmt.Println("-- Day 05 --")
	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
