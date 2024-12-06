package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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
	file, err := os.Open("inputs/05.txt")

	// If we don't have an input file return an error
	if err != nil {
		fmt.Println("Cannot find input file, please add your input file to inputs/05.txt")
		return
	}

	// Ensure we close the file
	defer file.Close()

	// File reader helper
	reader := bufio.NewReader(file)

	var rules, sequences = make(map[int][]int), make([][]int, 0)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		// If string contains "|" assume its a rule
		if strings.Contains(string(line), "|") {
			var l, r int
			// Scan the line and read the two values
			_, err = fmt.Sscanf(string(line), "%d|%d", &l, &r)
			if err != nil {
				fmt.Println("Error parsing line:", err)
				break
			}
			// Add rules to that number
			rules[l] = append(rules[l], r)
		} else {
			// Skip empty lines
			if string(line) == "" {
				continue
			}

			// Get line as list of strings
			sequences_list := strings.Split(string(line), ",")
			sequence_list_ints := make([]int, 0)
			// Convert strings to ints
			for n := 0; n < len(sequences_list); n++ {
				sequence, err := strconv.Atoi(sequences_list[n])
				if err != nil {
					fmt.Println(err)
					break
				}
				sequence_list_ints = append(sequence_list_ints, sequence)
			}
			// Add list of ints to sequences
			sequences = append(sequences, sequence_list_ints)
		}
	}

	var total_of_middle_numbers = 0
	// Loop through each sequence
	for n := 0; n < len(sequences); n++ {
		var valid = true
		// Loop through each int in the sequence
		for j := 0; j < len(sequences[n]); j++ {
			// Get the rules for the int
			var rules = rules[sequences[n][j]]
			// If there are no rules its safe to continue
			if rules == nil {
				continue
			}
			// If there are rules then check all previous elements conform
			for i := 0; i < j; i++ {
				if slices.Contains(rules, sequences[n][i]) {
					valid = false
					break
				}
			}
		}
		// If sequence is valid then add the middle number to the total
		if valid {
			middle_num := sequences[n][len(sequences[n])/2]
			total_of_middle_numbers = total_of_middle_numbers + middle_num
		}
	}

	fmt.Println(total_of_middle_numbers)
}

func part2() {
	// Read the input file
	file, err := os.Open("inputs/05.txt")

	// If we don't have an input file return an error
	if err != nil {
		fmt.Println("Cannot find input file, please add your input file to inputs/05.txt")
		return
	}

	// Ensure we close the file
	defer file.Close()

	// File reader helper
	reader := bufio.NewReader(file)

	var rules, sequences = make(map[int][]int), make([][]int, 0)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		// If string contains "|" assume its a rule
		if strings.Contains(string(line), "|") {
			var l, r int
			// Scan the line and read the two values
			_, err = fmt.Sscanf(string(line), "%d|%d", &l, &r)
			if err != nil {
				fmt.Println("Error parsing line:", err)
				break
			}
			// Add rules to that number
			rules[l] = append(rules[l], r)
		} else {
			// Skip empty lines
			if string(line) == "" {
				continue
			}

			// Get line as list of strings
			sequences_list := strings.Split(string(line), ",")
			sequence_list_ints := make([]int, 0)
			// Convert strings to ints
			for n := 0; n < len(sequences_list); n++ {
				sequence, err := strconv.Atoi(sequences_list[n])
				if err != nil {
					fmt.Println(err)
					break
				}
				sequence_list_ints = append(sequence_list_ints, sequence)
			}
			// Add list of ints to sequences
			sequences = append(sequences, sequence_list_ints)
		}
	}

	var total_of_middle_numbers = 0
	// Loop through each sequence
	for n := 0; n < len(sequences); n++ {
		var valid = true
		// Loop through each int in the sequence
		for j := 0; j < len(sequences[n]); j++ {
			// Get the rules for the int
			var rules = rules[sequences[n][j]]
			// If there are no rules its safe to continue
			if rules == nil {
				continue
			}
			// If there are rules then check all previous elements conform
			for i := 0; i < j; i++ {
				if slices.Contains(rules, sequences[n][i]) {
					valid = false
					break
				}
			}
		}
		// While invalid fix the lists
		for !valid {
			var is_valid = true
			// Loop through each int in the sequence
			for j := 0; j < len(sequences[n]); j++ {
				// Get the rules for the int
				var rules = rules[sequences[n][j]]
				// If there are no rules its safe to continue
				if rules == nil {
					continue
				}
				// If there are rules then check all previous elements conform
				for i := 0; i < j; i++ {
					// If a previous element does not conform then swap them
					if slices.Contains(rules, sequences[n][i]) {
						sequences[n][j], sequences[n][i] = sequences[n][i], sequences[n][j]
						is_valid = false
					}
				}
			}
			// If sequence is now valid then add the middle number to the total
			if (is_valid) {
				middle_num := sequences[n][len(sequences[n])/2]
				total_of_middle_numbers = total_of_middle_numbers + middle_num
				valid = true
			}
		}
	}

	fmt.Println(total_of_middle_numbers)
}
