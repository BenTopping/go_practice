package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
)

func main() {
	fmt.Println("-- Part 1 --")
	part1()
	fmt.Println("-- Part 2 --")
	part2()
}

func part1() {
	// Read the input file
	file, err := os.Open("inputs/day_1.txt")

	// If we don't have an input file return an error
	if err != nil {
		fmt.Println("Cannot find input file, please add your input file to inputs/day_1.txt")
		return
	}

	// Ensure we close the file
	defer file.Close()

	// File reader helper
	reader := bufio.NewReader(file)

	// Create the two lists from the
	var left_list, right_list = make([]int, 0), make([]int, 0)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		var l, r int
		// Scan the line and read the two values
		_, err = fmt.Sscanf(string(line), "%d %d", &l, &r)
		if err != nil {
			fmt.Println("Error parsing line:", err)
			break
		}
		// Append the values to the relevant lists
		left_list = append(left_list, l)
		right_list = append(right_list, r)
	}

	// Sort the lists to make calculating distance easier
	slices.Sort(left_list)
	slices.Sort(right_list)

	var total_distance int
	for n := 0; n < len(left_list); n++ {
		// Add the distance of both values to total value. Using math.Abs to ensure a positive value
		total_distance = total_distance + int(math.Abs(float64(left_list[n])-float64(right_list[n])))
	}

	// Print out the total distance
	fmt.Println("Total distance: ", total_distance)
}

func part2() {
	// Read the input file
	file, err := os.Open("inputs/day_1.txt")

	// If we don't have an input file return an error
	if err != nil {
		fmt.Println("Cannot find input file, please add your input file to inputs/day_1.txt")
		return
	}

	// Ensure we close the file
	defer file.Close()

	// File reader helper
	reader := bufio.NewReader(file)

	// Create the a left list and a right map
	var left_list, right_dict = make([]int, 0), make(map[int]int)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		var l, r int
		// Scan the line and read the two values
		_, err = fmt.Sscanf(string(line), "%d %d", &l, &r)
		if err != nil {
			fmt.Println("Error parsing line:", err)
			break
		}
		// Append the left value to the left list
		left_list = append(left_list, l)
		// Add 1 to the number of occurences for that value in a dict
		right_dict[r] = right_dict[r] + 1
	}

	var similarity_score int
	for n := 0; n < len(left_list); n++ {
		// Add the number of the left list value times the right list occurences to the similiarity score
		similarity_score = similarity_score + (left_list[n] * right_dict[left_list[n]])
	}

	fmt.Println("Similarity score: ", similarity_score)
}
