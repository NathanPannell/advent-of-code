package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadInput(day int) []string {
	inputFileName := fmt.Sprintf("../inputs/%d.in", day)
	file, err := os.Open(inputFileName)
	if err != nil {
		// Kill the program, there is no need to do error handling in each script
		log.Fatalf("unable to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	lines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error during scan: %s", err)
	}

	return lines
}
