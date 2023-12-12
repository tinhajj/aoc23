package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var max = struct {
	red   int
	green int
	blue  int
}{
	red:   12,
	green: 13,
	blue:  14,
}

var exp = struct {
	red   *regexp.Regexp
	green *regexp.Regexp
	blue  *regexp.Regexp
	game  *regexp.Regexp
}{
	red:   regexp.MustCompile(`(?P<number>\d{1,2}) red`),
	green: regexp.MustCompile(`(?P<number>\d{1,2}) green`),
	blue:  regexp.MustCompile(`(?P<number>\d{1,2}) blue`),
	game:  regexp.MustCompile(`Game (?P<number>\d{1,3})`),
}

func main() {
	bytes, _ := os.ReadFile("input.txt")
	input := string(bytes)
	lines := strings.Split(input, "\n")
	sum := 0
	//lineloop:
	for _, line := range lines {
		//game := get(exp.game, line)[0]

		mins := struct {
			red   int
			green int
			blue  int
		}{}

		reds := get(exp.red, line)
		for _, red := range reds {
			if red > mins.red {
				mins.red = red
			}
			//if red > max.red {
			//	continue lineloop
			//}
		}
		greens := get(exp.green, line)
		for _, green := range greens {
			if green > mins.green {
				mins.green = green
			}
			//if green > max.green {
			//	continue lineloop
			//}
		}
		blues := get(exp.blue, line)
		for _, blue := range blues {
			if blue > mins.blue {
				mins.blue = blue
			}
			//if blue > max.blue {
			//	continue lineloop
			//}
		}
		sum += mins.blue * mins.green * mins.red
	}
	fmt.Println(sum)
}

func get(re *regexp.Regexp, line string) []int {
	result := []int{}
	matches := re.FindAllStringSubmatch(line, -1)
	for _, match := range matches {
		num, _ := strconv.Atoi(match[1])
		result = append(result, num)
	}
	return result
}
