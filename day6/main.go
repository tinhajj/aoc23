package main

import "fmt"

type race struct {
	time   int
	record int
}

var races = []race{
	{59796575, 597123410321328},
}

func main() {
	raceWays := []int{}
	for _, race := range races {
		won := 0
		for charge := 0; charge < race.time; charge++ {
			distance := (race.time - charge) * charge
			if distance > race.record {
				won++
			}
		}
		raceWays = append(raceWays, won)
	}

	product := 1
	for _, raceWays := range raceWays {
		product = product * raceWays
	}
	fmt.Println(product)
}
