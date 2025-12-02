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

func part1(lines []string) string {
	// input is a single line of comma-separated ranges.
	line := strings.TrimSpace(lines[0])
	parts := strings.Split(line, ",")
	var total int64
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		bounds := strings.Split(p, "-")
		if len(bounds) != 2 {
			continue
		}
		a, _ := strconv.ParseInt(bounds[0], 10, 64)
		b, _ := strconv.ParseInt(bounds[1], 10, 64)
		if a > b {
			a, b = b, a
		}

		// Determine maximum half-length to consider (based on digits of b)
		maxDigits := len(strconv.FormatInt(b, 10))
		for L := 1; L*2 <= maxDigits; L++ {
			// pow10L = 10^L
			pow10L := int64(1)
			for i := 0; i < L; i++ {
				pow10L *= 10
			}
			divisor := pow10L + 1 // number = x*divisor
			// minX = ceil(a/divisor), maxX = floor(b/divisor)
			minX := (a + divisor - 1) / divisor
			maxX := b / divisor
			// enforce L-digit x (no leading zeros)
			lower := pow10L / 10
			upper := pow10L - 1
			if minX < lower {
				minX = lower
			}
			if maxX > upper {
				maxX = upper
			}
			if minX > maxX {
				continue
			}
			count := (maxX - minX + 1)
			sumX := (minX + maxX) * count / 2
			total += sumX * divisor
		}
	}
	return fmt.Sprintf("%d", total)
}

func part2(lines []string) string {
	line := strings.TrimSpace(lines[0])
	parts := strings.Split(line, ",")
	seen := make(map[int64]struct{})
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		bounds := strings.Split(p, "-")
		if len(bounds) != 2 {
			continue
		}
		a, _ := strconv.ParseInt(strings.TrimSpace(bounds[0]), 10, 64)
		b, _ := strconv.ParseInt(strings.TrimSpace(bounds[1]), 10, 64)
		if a > b {
			a, b = b, a
		}

		maxDigits := len(strconv.FormatInt(b, 10))
		// L = length of base block
		for L := 1; L <= maxDigits-1; L++ {
			for k := 2; L*k <= maxDigits; k++ {
				// multiplier = (10^{L*k} - 1) / (10^L - 1)
				powL := pow10(L)
				powLK := pow10(L * k)
				if powL == 1 {
					continue
				}
				multiplier := (powLK - 1) / (powL - 1)
				if multiplier <= 0 {
					continue
				}
				minX := (a + multiplier - 1) / multiplier
				maxX := b / multiplier
				lower := powL / 10
				upper := powL - 1
				if minX < lower {
					minX = lower
				}
				if maxX > upper {
					maxX = upper
				}
				if minX > maxX {
					continue
				}
				// add each produced number to set to avoid duplicates
				for x := minX; x <= maxX; x++ {
					n := x * multiplier
					if n < a || n > b {
						continue
					}
					seen[n] = struct{}{}
				}
			}
		}
	}

	var total int64
	for n := range seen {
		total += n
	}
	return fmt.Sprintf("%d", total)
}

func pow10(n int) int64 {
	v := int64(1)
	for i := 0; i < n; i++ {
		v *= 10
	}
	return v
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
