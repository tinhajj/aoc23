package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Card struct {
	winning []int
}

func main() {
	bytes, _ := os.ReadFile("input.txt")
	content := string(bytes)
	lines := strings.Split(content, "\r\n")
	var part string

	for _, line := range lines {
		part = capture(`Card\s*(\d*):`, line)
		cardNumber, _ := strconv.Atoi(part)

		part = capture(`Card.*:([\d\s]*) |`, line)
		parts := strings.Split(part, " ")
		winningNumbers := []int{}

		for _, part := range parts {
			number, err := strconv.Atoi(part)
			if err != nil {
				continue
			}
			winningNumbers = append(winningNumbers, number)
		}

		myNumbers := []int{}
		part = capture(`Card.*:[\d\s]* | ([\d\s]*)`, line)
		fmt.Println("the part is", part)
		parts = strings.Split(part, " ")
		for _, part := range parts {
			number, err := strconv.Atoi(part)
			if err != nil {
				continue
			}
			myNumbers = append(myNumbers, number)
		}

		fmt.Println("card number", cardNumber)
		fmt.Println("winning numbers", winningNumbers)
		fmt.Println("my numbers", myNumbers)
	}
}

func capture(exp string, haystack string) string {
	re := regexp.MustCompile(exp)
	matches := re.FindAllStringSubmatch(haystack, -1)
	fmt.Println("all submatches:", matches)
	//fmt.Println("from the capture")
	//for i, result := range results {
	//	fmt.Printf("idx: %d, value: %s ", i, result)
	//}

	for i, match := range matches {
		fmt.Println("looking at", i, match, len(match))
		if len(match) >= 2 {
			fmt.Println("the chosen submatch", match[1])
			return match[1]
		}
	}
	panic(0)
}
