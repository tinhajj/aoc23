package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type adventHand struct {
	suit        suit
	hand        string
	handCompare string
	bid         int
}

type suit int

const (
	highcard suit = iota
	onepair
	twopair
	threeofkind
	fullhouse
	fourofkind
	fiveofkind
)

var fiveofkindcheck = []int{5}
var fourofkindcheck = []int{1, 4}
var threeofkindcheck = []int{1, 1, 3}
var fullhousecheck = []int{2, 3}
var twopaircheck = []int{1, 2, 2}
var onepaircheck = []int{1, 1, 1, 2}
var highcheck = []int{1, 1, 1, 1, 1}

func main() {
	bytes, _ := os.ReadFile("input.txt")
	content := string(bytes)

	hands := []*adventHand{}

	for _, line := range strings.Split(content, "\r\n") {
		parts := strings.Split(line, " ")

		hand := parts[0]
		bid, err := strconv.Atoi(parts[1])
		if err != nil {
			panic("failed to parse")
		}

		ah := &adventHand{hand: hand, bid: bid}
		getSuit(ah)
		handReplace(ah)
		hands = append(hands, ah)
	}

	sort.Slice(hands, func(i, j int) bool {
		leftHand := hands[i]
		rightHand := hands[j]
		if leftHand.suit < rightHand.suit {
			return true
		}
		if leftHand.suit > rightHand.suit {
			return false
		}
		return leftHand.handCompare < rightHand.handCompare
	})

	product := 0
	for i, hand := range hands {
		product += (i + 1) * hand.bid
		fmt.Println(hand)
	}
	fmt.Println(product)
}

func getSuit(ah *adventHand) {
	m := make(map[string]int)
	for _, c := range ah.hand {
		m[string(c)]++
	}

	format := mapToCounts(m)
	if sameSlice(fiveofkindcheck, format) {
		ah.suit = fiveofkind
	}
	if sameSlice(fourofkindcheck, format) {
		ah.suit = fourofkind
	}
	if sameSlice(threeofkindcheck, format) {
		ah.suit = threeofkind
	}
	if sameSlice(fullhousecheck, format) {
		ah.suit = fullhouse
	}
	if sameSlice(twopaircheck, format) {
		ah.suit = twopair
	}
	if sameSlice(onepaircheck, format) {
		ah.suit = onepair
	}
	if sameSlice(highcheck, format) {
		ah.suit = highcard
	}
}

func handReplace(ah *adventHand) {
	ah.handCompare = ah.hand
	ah.handCompare = strings.Replace(ah.handCompare, "A", "Z", -1)
	ah.handCompare = strings.Replace(ah.handCompare, "K", "Y", -1)
	ah.handCompare = strings.Replace(ah.handCompare, "Q", "X", -1)
	ah.handCompare = strings.Replace(ah.handCompare, "J", "W", -1)
	ah.handCompare = strings.Replace(ah.handCompare, "T", "V", -1)
}

func mapToCounts(m map[string]int) []int {
	result := []int{}
	for _, value := range m {
		result = append(result, value)
	}
	sort.Ints(result)
	return result
}

func sameSlice(first, second []int) bool {
	if len(first) != len(second) {
		return false
	}

	for i := 0; i < len(first); i++ {
		if first[i] != second[i] {
			return false
		}
	}

	return true
}
