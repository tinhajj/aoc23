package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type record struct {
	location *location
	step     int
}

type location struct {
	name string

	// jank?
	leftName  string
	rightName string

	left  *location
	right *location
}

func main() {
	p := message.NewPrinter(language.English)
	bytes, _ := os.ReadFile("input.txt")
	content := string(bytes)
	lines := strings.Split(content, "\r\n")

	directions := lines[0]

	lines = lines[2:]
	locations := []*location{}

	for _, line := range lines {
		locName := capture(`^(\w\w\w) =`, line)
		loc := &location{name: locName}

		left := capture(`^\w\w\w = \((\w\w\w)`, line)
		loc.leftName = left

		right := capture(`^\w\w\w = \(\w\w\w, (\w\w\w)`, line)
		loc.rightName = right

		locations = append(locations, loc)
	}

	for _, loc := range locations {
		loc.left = findLoc(loc.leftName, locations)
		loc.right = findLoc(loc.rightName, locations)
	}

	// solve
	clones := findLocEndsWithA(locations)

	directionStep := 0
	totalSteps := 0
	//endsWithZ := 0

	hits := []*record{}
	index := 0
	current := clones[index]
	fmt.Print("starting ", current.name, " ")

	for {

		if current.name[2] == 'Z' {
			r := &record{current, totalSteps}
			hits = append(hits, r)
			if len(hits) >= 2 {
				difference := hits[len(hits)-1].step - hits[len(hits)-2].step
				fmt.Printf("%s @ %d, ", r.location.name, difference)
			} else {
				fmt.Printf("%s @ %d, ", r.location.name, totalSteps)
			}
		}

		if len(hits) > 5 {
			index++
			hits = []*record{}
			totalSteps = 0

			if index >= len(clones) {
				os.Exit(1)
			}

			fmt.Println()
			fmt.Print("starting ", current.name, " ")
			current = clones[index]
		}

		totalSteps++
		direction := directions[directionStep]
		if direction == 'L' {
			current = current.left
		}
		if direction == 'R' {
			current = current.right
		}

		directionStep++
		if directionStep == len(directions) {
			directionStep = 0
		}
	}

	//for {
	//	if endsWithZ == len(clones) {
	//		break
	//	}

	//	endsWithZ = 0
	//	for i := 0; i < len(clones); i++ {
	//		direction := directions[directionStep]
	//		if direction == 'L' {
	//			clones[i] = clones[i].left
	//		}
	//		if direction == 'R' {
	//			clones[i] = clones[i].right
	//		}
	//		if clones[i].name[2] == 'Z' {
	//			endsWithZ++
	//		}
	//	}

	//	directionStep++
	//	totalSteps++

	//	if totalSteps%100_000_000 == 0 {
	//		p.Printf("%d\n", totalSteps)
	//	}

	//	if directionStep == len(directions) {
	//		directionStep = 0
	//	}
	//}
	p.Printf("%d\n", totalSteps)
}

func findLocEndsWithA(locations []*location) []*location {
	matches := []*location{}
	for _, loc := range locations {
		if strings.HasSuffix(loc.name, "A") {
			matches = append(matches, loc)
		}
	}
	return matches
}

func findLoc(name string, locations []*location) *location {
	for _, loc := range locations {
		if loc.name == name {
			return loc
		}
	}
	panic("unable to find matching location")
}

func capture(exp string, haystack string) string {
	re := regexp.MustCompile(exp)

	submatches := re.FindAllStringSubmatch(haystack, -1)
	for _, submatch := range submatches {
		if submatch[1] != "" {
			return submatch[1]
		}
	}
	panic("unable to find capture")
}
