package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

func maxSubseqDigits(s string, k int) string {
	n := len(s)
	if k >= n {
		return s
	}
	var b strings.Builder
	start := 0
	for rem := k; rem > 0; rem-- {
		end := n - rem
		maxIdx := start
		maxCh := s[start]
		// find largest digit in s[start : end+1]
		for i := start; i <= end; i++ {
			if s[i] > maxCh {
				maxCh = s[i]
				maxIdx = i
				if maxCh == '9' {
					break
				}
			}
		}
		b.WriteByte(maxCh)
		start = maxIdx + 1
	}
	return b.String()
}

func solve(lines []string, k int) string {
	// find the largest combination of two ordered digits from each line and sum them, e.g. for line "123435" it is 45
	// iterate lines
	var total int64
	for _, ln := range lines {
		if len(ln) < k {
			k = len(ln)
		}
		numStr := maxSubseqDigits(ln, k)
		v, _ := strconv.ParseInt(numStr, 10, 64)
		total += v
	}
	return fmt.Sprintf("%d", total)
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
	fmt.Println("Part 1:", solve(lines, 2))
	fmt.Println("Part 2:", solve(lines, 12))
}
