package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type adventMap struct {
	ranges []*adventRange
}

type adventRange struct {
	destination int
	source      int
	run         int
}

var header = regexp.MustCompile(`\w*-\w*-\w* map:$`)

func main() {
	fmt.Println("start")
	bytes, _ := os.ReadFile("input.txt")
	content := string(bytes)
	lines := strings.Split(content, "\r\n")
	first := capture(`seeds: ([\d\s]*)$`, lines[0])

	seeds := digitize(splitTrim(first, " "))

	amaps := []adventMap{}

	lines = lines[1:]
	for i := 0; i < len(lines); i++ {
		line := lines[i]

		adventMap := adventMap{ranges: []*adventRange{}}

		if header.MatchString(line) {
			i++
			line = lines[i]
			for line != "" && i < len(lines) {
				all := digitize(strings.Split(line, " "))
				destination := all[0]
				source := all[1]
				run := all[2]
				adventMap.ranges = append(adventMap.ranges, &adventRange{destination, source, run})

				i++
				if i < len(lines) {
					line = lines[i]
				}
			}
			fmt.Println("done reading a range")
			amaps = append(amaps, adventMap)
		}
	}
	fmt.Println(seeds)
	fmt.Println(amaps)
}

func capture(exp string, haystack string) string {
	re := regexp.MustCompile(exp)
	matches := re.FindAllStringSubmatch(haystack, -1)
	for _, match := range matches {
		if len(match) > 1 && match[1] != "" {
			return match[1]
		}
	}

	panic("nothing to capture in haystack")
}

func splitTrim(line string, sep string) []string {
	parts := []string{}
	for _, part := range strings.Split(line, sep) {
		trim := strings.TrimSpace(part)
		if trim != "" {
			parts = append(parts, trim)
		}
	}
	return parts
}

func digitize(parts []string) []int {
	digits := []int{}
	for _, part := range parts {
		digit, err := strconv.Atoi(part)
		if err != nil {
			panic("failed to part string")
		}
		digits = append(digits, digit)
	}
	return digits
}
