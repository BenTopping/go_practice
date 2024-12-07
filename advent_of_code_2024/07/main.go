package main

import (
	"bufio"
	"fmt"
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
	file, err := os.Open("inputs/07.txt")

	// If we don't have an input file return an error
	if err != nil {
		fmt.Println("Cannot find input file, please add your input file to inputs/07.txt")
		return
	}

	// Ensure we close the file
	defer file.Close()

	// File reader helper
	reader := bufio.NewReader(file)

	// Create a map of total to its component values
	var lines = make(map[int][]int, 0)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		// Split string to get total and string of values
		split_string := strings.Split(string(line), ": ")
		total, err := strconv.Atoi(string(split_string[0]))
		if err != nil {
			fmt.Println("Error: Unable to identify total")
			break
		}

		// split string of values
		split_string = strings.Split(split_string[1], " ")
		for n := 0; n < len(split_string); n++ {
			// Convert string values to ints
			value, err := strconv.Atoi(string(split_string[n]))
			if err != nil {
				fmt.Println("Error: Unable to read value")
				break
			}
			// Add value to total key
			lines[total] = append(lines[total], value)
		}
	}

	var sum_valid_totals = 0
	// For each key pair check if its solvable
	for total, values := range lines {
		if (canSolve(total, values, []string{"*", "+"})) {
			sum_valid_totals += total
		}
	}

	fmt.Println("Sum valid total: ", sum_valid_totals)
}

func canSolve(total int, values []int, operators []string)  bool {
	// If there is only 1 value then check if it equals the total
	if (len(values) == 1) {
		return total == values[0]
	}
	// Get the current value and the next value
	left_value := values[0]
	right_value := values[1]
	// Get the remaining values
	remaining_values := []int{}
	if len(values) > 2 {
		remaining_values = values[2:]
	}
	// For each operator calculate the current values
	for _, operator := range operators {
		var result int
		switch operator {
		case "*":
			result = left_value * right_value
		case "+":
			result = left_value + right_value
		case "||":
			result, _ = strconv.Atoi((strconv.Itoa(left_value) + strconv.Itoa(right_value)))
		}

		// set the next set of values to be the currenlt result plus remain values
		next_values := []int{result}
		next_values = append(next_values, remaining_values...)

		// If there are no remaining values and result equals total we have solved
		// If the result is less than the total check if we can still solve with next values
		if len(remaining_values) == 0 && result == total {
			return true
		} else if result <= total && canSolve(total, next_values, operators) {
			return true
		}
	}
	return false
}

func part2() {
	// Read the input file
	file, err := os.Open("inputs/07.txt")

	// If we don't have an input file return an error
	if err != nil {
		fmt.Println("Cannot find input file, please add your input file to inputs/07.txt")
		return
	}

	// Ensure we close the file
	defer file.Close()

	// File reader helper
	reader := bufio.NewReader(file)

	// Create a map of total to its component values
	var lines = make(map[int][]int, 0)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		// Split string to get total and string of values
		split_string := strings.Split(string(line), ": ")
		total, err := strconv.Atoi(string(split_string[0]))
		if err != nil {
			fmt.Println("Error: Unable to identify total")
			break
		}

		// split string of values
		split_string = strings.Split(split_string[1], " ")
		for n := 0; n < len(split_string); n++ {
			// Convert string values to ints
			value, err := strconv.Atoi(string(split_string[n]))
			if err != nil {
				fmt.Println("Error: Unable to read value")
				break
			}
			// Add value to total key
			lines[total] = append(lines[total], value)
		}
	}

	var sum_valid_totals = 0
	// For each key pair check if its solvable
	for total, values := range lines {
		if (canSolve(total, values, []string{"*", "+", "||"})) {
			sum_valid_totals += total
		}
	}

	fmt.Println("Sum valid total: ", sum_valid_totals)
}
