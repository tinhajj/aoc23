package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Card struct {
	id      int
	winning []int
	values  []int

	matches []int
}

func main() {
	bytes, _ := os.ReadFile("input.txt")
	content := string(bytes)
	lines := strings.Split(content, "\r\n")
	var part string

	cards := []*Card{}

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
		parts = strings.Split(part, " ")
		for _, part := range parts {
			number, err := strconv.Atoi(part)
			if err != nil {
				continue
			}
			myNumbers = append(myNumbers, number)
		}

		cards = append(cards, &Card{id: cardNumber, winning: winningNumbers, values: myNumbers})
	}

	sum := 0
	for _, card := range cards {
		for _, value := range card.values {
			for _, win := range card.winning {
				if value == win {
					card.matches = append(card.matches, value)
					break
				}
			}
		}

		fmt.Println(int(math.Pow(2, float64(len(card.matches)))))
		if len(card.matches) == 0 {
			continue
		}
		sum += int(math.Pow(2, float64(len(card.matches)-1)))
	}
	fmt.Println(sum)

	scratches := 0
	stack := cards
	for i := 0; i < len(stack); i++ {
		scratches++
		card := stack[i]
		matches := len(card.matches)
		end := card.id + matches
		if card.id == len(cards) {
			continue
		}
		if end > len(cards)-1 {
			end = len(cards) - 1
		}
		stack = append(stack, cards[card.id:end]...)
	}
	fmt.Println(len(stack))
}

func capture(exp string, haystack string) string {
	re := regexp.MustCompile(exp)
	matches := re.FindAllStringSubmatch(haystack, -1)

	//for i, result := range matches {
	//	fmt.Printf("idx: %d, value: %s\n", i, result)
	//}

	for _, match := range matches {
		//fmt.Println("looking at", i, match, len(match))
		if len(match) >= 2 && match[1] != "" {
			//fmt.Println("the chosen submatch", match[1])
			return match[1]
		}
	}
	panic(0)
}
