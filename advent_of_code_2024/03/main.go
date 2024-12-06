package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	fmt.Println("-- Part 1 --")
	part1()
	fmt.Println("-- Part 2 --")
	part2()
}

func part1() {
	// Read the input file
	file, err := os.ReadFile("inputs/03.txt")

	// If we don't have an input file return an error
	if err != nil {
		fmt.Println("Cannot find input file, please add your input file to inputs/03.txt")
		return
	}

	var total_value = 0
	// Regexp for matching the multipliers
	re := regexp.MustCompile(`mul\([0-9]{0,3},[0-9]{0,3}\)`)

	// Loop over the file matches
	matches := re.FindAllStringSubmatch(string(file), -1)
	for n := 0; n < len(matches); n++ {
		var v1, v2 int
		// Scan the line and read the two values
		_, err = fmt.Sscanf(string(matches[n][0]), "mul(%d,%d)", &v1, &v2)
		if err != nil {
			break
		}
		total_value = total_value + (v1 * v2)
	}

	fmt.Println("Total value: ", total_value)
}

func part2() {
	// Read the input file
	file, err := os.ReadFile("inputs/03.txt")

	// If we don't have an input file return an error
	if err != nil {
		fmt.Println("Cannot find input file, please add your input file to inputs/03.txt")
		return
	}

	var total_value = 0
	// Regexp for matching the multipliers
	re := regexp.MustCompile(`mul\([0-9]{0,3},[0-9]{0,3}\)|do\(\)|don't\(\)`)

	// Set default do to true
	var do = true

	// Loop over the file matches
	matches := re.FindAllStringSubmatch(string(file), -1)
	for n := 0; n < len(matches); n++ {
		if matches[n][0] == "do()" {
			do = true
		} else if matches[n][0] == "don't()" {
			do = false
		} else {
			var v1, v2 int

			// Scan the line and read the two values
			_, err = fmt.Sscanf(string(matches[n][0]), "mul(%d,%d)", &v1, &v2)
			if err != nil {
				break
			}
			if do {
				total_value = total_value + (v1 * v2)
			}
		}
	}

	fmt.Println("Total value: ", total_value)
}
