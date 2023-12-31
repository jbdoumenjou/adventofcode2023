package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("part2-input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var powerSum int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		content := scanner.Text()
		powerSum += powerOfCubes(content)
	}

	fmt.Println(powerSum)
}

func powerOfCubes(s string) int {
	minBag := make(map[string]int)
	// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green

	game := strings.Split(s, ":")
	// 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	sets := strings.Split(game[1], ";")
	for _, set := range sets {
		// 3 blue, 4 red
		cubes := strings.Split(set, ",")
		for _, cube := range cubes {
			// 3 blue
			elt := strings.Split(strings.TrimSpace(cube), " ")
			val, err := strconv.Atoi(elt[0])
			if err != nil {
				log.Fatal(err)
			}

			if val > minBag[elt[1]] {
				minBag[elt[1]] = val
			}
		}
	}

	power := 1
	for _, v := range minBag {
		power *= v
	}

	return power
}
