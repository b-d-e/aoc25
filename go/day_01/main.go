package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	pos := 50
	countZero := 0
	for _, ln := range lines {
		if ln == "" {
			continue
		}
		dir := ln[0]

		var dist int
		_, err := fmt.Sscanf(ln[1:], "%d", &dist)
		if err != nil {
			continue
		}
		switch dir {
		case 'L':
			pos = (pos - dist) % 100
			if pos < 0 {
				pos += 100
			}
		case 'R':
			pos = (pos + dist) % 100
		default:
			continue
		}
		if pos == 0 {
			countZero++
		}
	}
	return fmt.Sprintf("%d", countZero)
}

func part2(lines []string) string {
	pos := 50
	count := 0
	for _, ln := range lines {
		if ln == "" {
			continue
		}
		dir := ln[0]
		var dist int
		_, err := fmt.Sscanf(ln[1:], "%d", &dist)
		if err != nil {
			continue
		}

		// For each rotation count how many times the dial points at 0 during the clicks.
		// Moving right: positions s+1 .. s+dist. It hits 0 when (s + t) % 100 == 0.
		// Moving left: positions s-1 .. s-dist. It hits 0 when (s - t) % 100 == 0.
		switch dir {
		case 'R':
			// steps to first 0 from pos when moving right
			offset := (100 - (pos % 100)) % 100
			if offset == 0 {
				offset = 100
			}
			if dist >= offset {
				count += 1 + (dist-offset)/100
			}
			pos = (pos + dist) % 100
		case 'L':
			// steps to first 0 from pos when moving left
			offset := pos % 100
			if offset == 0 {
				offset = 100
			}
			if dist >= offset {
				count += 1 + (dist-offset)/100
			}
			pos = (pos - dist) % 100
			if pos < 0 {
				pos += 100
			}
		default:
			continue
		}
	}
	return fmt.Sprintf("%d", count)
}

func main() {
	inputPath := "input.txt"
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
