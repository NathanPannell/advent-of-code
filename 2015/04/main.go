package main

import (
	"advent/utils"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"strconv"
)

func solve(input string, upperBound int) (int, int) {
	fiveLeadingZeroes := -1
	for i := 0; i < upperBound; i++ {
		candidate := input + strconv.Itoa(i)

		hashFunc := md5.New()
		io.WriteString(hashFunc, candidate)
		checksum := hex.EncodeToString(hashFunc.Sum(nil))

		if fiveLeadingZeroes == -1 && checksum[:5] == "00000" {
			fmt.Printf("%s: %s\n", candidate, checksum)
			fiveLeadingZeroes = i
		}

		if checksum[:6] == "000000" {
			fmt.Printf("%s: %s\n", candidate, checksum)
			return fiveLeadingZeroes, i
		}
	}

	fmt.Printf("Unable to find both solutions with i < %d\n", upperBound)
	return fiveLeadingZeroes, -1
}

func main() {
	input := utils.ReadInput(4)[0]
	part1, part2 := solve(input, 10000000)

	fmt.Println("-- Day 04 --")
	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
