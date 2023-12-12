package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type partNumber struct {
	digits    []*digit
	start     location
	schematic bool
}

type digit struct {
	value    int
	location location
}

type symbol struct {
	location location
}

type location struct {
	x int
	y int
}

func main() {
	b, _ := os.ReadFile("input.txt")
	content := string(b)
	lines := strings.Split(content, "\r\n")

	partNumbers := []*partNumber{}
	symbols := []*symbol{}

	for y := 0; y < len(lines); y++ {
		line := lines[y]

		for x := 0; x < len(line); {
			var err error
			var value int
			var pn partNumber
			pn.digits = []*digit{}

			start := x

			for err == nil && x < len(line) {
				char := string(line[x])
				value, err = strconv.Atoi(char)

				if err == nil {
					pn.digits = append(pn.digits, &digit{value: value, location: location{x, y}})
				}

				if char != "." && err != nil {
					symbols = append(symbols, &symbol{location: location{x, y}})
				}

				x++
			}

			if len(pn.digits) > 0 {
				pn.start = location{start, y}
				partNumbers = append(partNumbers, &pn)
			}
		}
	}

	sum := 0
	for _, symbol := range symbols {
		nearby := []int{}
		total := 0

		for _, pn := range partNumbers {
			pn.schematic = partNearSymbol(pn, symbol.location)
			if pn.schematic {
				nearby = append(nearby, partNumberValue(pn))
				total++
			}
		}

		if total == 2 {
			sum += nearby[0] * nearby[1]
		}
	}
	fmt.Println(sum)

	//sum := 0
	//for _, pn := range partNumbers {
	//	if pn.schematic {
	//		sum += partNumberValue(pn)
	//	}
	//}
	//fmt.Println(sum)
}

func partNumberValue(pn *partNumber) int {
	multiplier := 1
	sum := 0
	for i := len(pn.digits) - 1; i >= 0; i-- {
		digit := pn.digits[i]
		sum += digit.value * multiplier
		multiplier = multiplier * 10
	}
	return sum
}

func partNearSymbol(pn *partNumber, l location) bool {
	y := false
	x := false

	if pn.start.y >= l.y-1 && pn.start.y <= l.y+1 {
		y = true
	}

	for _, digit := range pn.digits {
		if digit.location.x >= l.x-1 && digit.location.x <= l.x+1 {
			x = true
		}
	}

	return x && y
}
