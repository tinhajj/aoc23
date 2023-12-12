package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var shorts = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
var longs = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

var sum = 0

func main() {
	input, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(input), "\n")
	lines = lines[0 : len(lines)-1]

	for _, line := range lines {
		numbers := []int{}
		// parse all the numbers on a line
		for i := 0; i < len(line); i++ {
			number, ok := matchShort(i, line)
			if ok {
				numbers = append(numbers, number)
				continue
			}

			number, ok = matchLong(i, line)
			if ok {
				numbers = append(numbers, number)
				continue
			}
		}

		var first int
		var last int
		// get and add the right numbers
		if len(numbers) == 1 {
			first = numbers[0]
			last = numbers[0]
		} else {
			first = numbers[0]
			last = numbers[len(numbers)-1]
		}
		lineResult := (first * 10) + last
		sum += lineResult
	}
	fmt.Println(sum)
}

func matchShort(i int, s string) (int, bool) {
	for _, short := range shorts {
		if short == string(s[i]) {
			number, _ := strconv.Atoi(short)
			return number, true
		}
	}
	return 0, false
}

func matchLong(start int, s string) (int, bool) {
	for i, long := range longs {
		if start+len(long) > len(s) {
			continue
		}
		if s[start:start+len(long)] == long {
			return i + 1, true
		}
	}
	return 0, false
}
