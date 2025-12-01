package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// readLines reads all lines from a file path.
func readLines(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func part1(lines []string) string {
	// TODO: implement part 1
	return ""
}

func part2(lines []string) string {
	// TODO: implement part 2
	return ""
}

func main() {
	inputPath := "input.txt" // put AoC input in this file
	if len(os.Args) > 1 {
		inputPath = os.Args[1]
	}
	lines, err := readLines(inputPath)
	if err != nil {
		log.Fatalf("reading input: %v", err)
	}
	fmt.Println("Part 1:", part1(lines))
	fmt.Println("Part 2:", part2(lines))
}
