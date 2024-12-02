package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("-- Part 1 --")
	part1()
	fmt.Println("-- Part 2 --")
	part2()
}

func part1() {
	// Read the input file
	file, err := os.Open("inputs/02.txt")

	// If we don't have an input file return an error
	if err != nil {
		fmt.Println("Cannot find input file, please add your input file to inputs/02.txt")
		return
	}

	// Ensure we close the file
	defer file.Close()

	// File reader helper
	reader := bufio.NewReader(file)

	var safe_reports = 0
	// Loop over each file line
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		// Get each level in the line
		string_levels := strings.Split(string(line), " ")

		var safe = true
		// direction false is increasing, true is decreasing
		var direction = false
		for n := 0; n < len(string_levels)-1; n++ {
			level1, err := strconv.Atoi(string_levels[n])
			if err != nil {
				fmt.Println(err)
				break
			}
			level2, err := strconv.Atoi(string_levels[n+1])

			// Check difference is valid and not 0
			diff := math.Abs(float64(level1 - level2))
			if diff == 0 || diff > 3 {
				safe = false
				break
			}

			// If it is the first level, determine the direction
			if n == 0 {
				direction = level1 > level2
			}
			// If the direction is not the same as existing then break
			if level1 > level2 != direction {
				safe = false
				break
			}

		}

		if safe == true {
			safe_reports++
		}
	}

	// Print out the total safe reports
	fmt.Println("Total safe reports: ", safe_reports)
}

func part2() {
	// Read the input file
	file, err := os.Open("inputs/02.txt")

	// If we don't have an input file return an error
	if err != nil {
		fmt.Println("Cannot find input file, please add your input file to inputs/02.txt")
		return
	}

	// Ensure we close the file
	defer file.Close()

	// File reader helper
	reader := bufio.NewReader(file)

	var safe_reports = 0
	// Loop over each file line
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		// Get each level in the line
		string_levels := strings.Split(string(line), " ")

		// Check if the current report is safe
		safe := isSafe(string_levels)

		if safe == true {
			safe_reports++
		} else {
			// If the current report is not safe, check all permutations with one missing element
			for n := 0; n < len(string_levels); n++ {
				sliced_levels := make([]string, 0)
				sliced_levels = append(sliced_levels, string_levels[:n]...)
				sliced_levels = append(sliced_levels, string_levels[n+1:]...)
				if isSafe(sliced_levels) { 
					safe_reports++
					break 
				}
			}
		}
	}

	// Print out the total safe reports
	fmt.Println("Total safe reports: ", safe_reports)
}

func isSafe(levels []string) bool {
	var safe = true
	// direction false is increasing, true is decreasing
	var direction = false
	for n := 0; n < len(levels)-1; n++ {
		level1, err := strconv.Atoi(levels[n])
		if err != nil {
			fmt.Println(err)
			break
		}
		level2, err := strconv.Atoi(levels[n+1])

		// Check difference is valid and not 0
		diff := math.Abs(float64(level1 - level2))
		if diff == 0 || diff > 3 {
			safe = false
			break
		}

		// If it is the first level, determine the direction
		if n == 0  {
			direction = level1 > level2
		}
		// If the direction is not the same as existing then break
		if level1 > level2 != direction {
			safe = false
		}
	}
	return safe
}
