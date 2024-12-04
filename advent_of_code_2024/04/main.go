package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
	Its not pretty but works
	A vector map would have been nice but golangs slice out of bounds errors make it tricky
*/

func main() {
	fmt.Println("-- Part 1 --")
	part1()
	fmt.Println("-- Part 2 --")
	part2()
}

func part1() {
	// Read the input file
	file, err := os.Open("inputs/04.txt")

	// If we don't have an input file return an error
	if err != nil {
		fmt.Println("Cannot find input file, please add your input file to inputs/04.txt")
		return
	}

	// Ensure we close the file
	defer file.Close()

	// File reader helper
	reader := bufio.NewReader(file)

	// Get each line of the file
	var lines = make([]string, 0)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		lines = append(lines, string(line))
	}

	total_xmas_count := 0

	// Loop through each line
	for n := 0; n < len(lines); n++ {
		// Loop through each letter on each line
		for j := 0; j < len(lines[n]); j++ {
			// Get the current line as a list of letters
			current_line := strings.Split(lines[n], "")
			// If the current letter is X check all possible directions for XMAS
			if current_line[j] == "X" {
				// Check reverse
				if (j > 2) {
					if strings.Join(current_line[j-3:j], "") == "SAM" {
						fmt.Println(current_line[j-3:j])
						total_xmas_count++
					}
				}
				// Check forward
				if (j < len(current_line)-3) {
					if strings.Join(current_line[j+1:j+4], "") == "MAS" {
						fmt.Println(current_line[j+1:j+4])
						total_xmas_count++
					}
				}
				// Check upwards and upwards diagonals
				if (n > 2) {
					if (
						strings.Split(lines[n-1], "")[j] == "M" &&
						strings.Split(lines[n-2], "")[j] == "A" &&
						strings.Split(lines[n-3], "")[j] == "S"){

						total_xmas_count++
					}
					// Diagonally up and left
					if (
						j > 2 &&
						strings.Split(lines[n-1], "")[j-1] == "M" &&
						strings.Split(lines[n-2], "")[j-2] == "A" &&
						strings.Split(lines[n-3], "")[j-3] == "S"){

						total_xmas_count++
					}
					// Diagonally up and right
					if (
						j < len(current_line)-3 &&
						strings.Split(lines[n-1], "")[j+1] == "M" &&
						strings.Split(lines[n-2], "")[j+2] == "A" &&
						strings.Split(lines[n-3], "")[j+3] == "S"){

						total_xmas_count++
					}
				}

				if (n < len(lines)-3) {
					if (
						strings.Split(lines[n+1], "")[j] == "M" &&
						strings.Split(lines[n+2], "")[j] == "A" &&
						strings.Split(lines[n+3], "")[j] == "S"){

						total_xmas_count++
					}
					// Diagonally up and left
					if (
						j > 2 &&
						strings.Split(lines[n+1], "")[j-1] == "M" &&
						strings.Split(lines[n+2], "")[j-2] == "A" &&
						strings.Split(lines[n+3], "")[j-3] == "S"){

						total_xmas_count++
					}
					// Diagonally up and right
					if (
						j < len(current_line)-3 &&
						strings.Split(lines[n+1], "")[j+1] == "M" &&
						strings.Split(lines[n+2], "")[j+2] == "A" &&
						strings.Split(lines[n+3], "")[j+3] == "S"){

						total_xmas_count++
					}
				}
			}
		}
	}

	// Print out the total distance
	fmt.Println("Total occurences of XMAS: ", total_xmas_count)
}

func part2() {
	// Read the input file
	file, err := os.Open("inputs/04.txt")

	// If we don't have an input file return an error
	if err != nil {
		fmt.Println("Cannot find input file, please add your input file to inputs/04.txt")
		return
	}

	// Ensure we close the file
	defer file.Close()

	// File reader helper
	reader := bufio.NewReader(file)

	// Get each line of the file
	var lines = make([]string, 0)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		lines = append(lines, string(line))
	}

	total_xmas_count := 0

	// Loop through each line
	for n := 0; n < len(lines); n++ {
		// Loop through each letter on each line
		for j := 0; j < len(lines[n]); j++ {
			// Get the current line as a list of letters
			current_line := strings.Split(lines[n], "")
			// If the current letter is A check all diagonals for MAS
			if current_line[j] == "A" {
				if (n > 0 && n < len(lines)-1 && j > 0 && j < len(current_line)-1) {
					if (
						strings.Split(lines[n-1], "")[j-1] == "M" &&
						strings.Split(lines[n+1], "")[j+1] == "S"){
							if (
								strings.Split(lines[n-1], "")[j+1] == "M" &&
								strings.Split(lines[n+1], "")[j-1] == "S"){
									total_xmas_count++
							} else if (
								strings.Split(lines[n-1], "")[j+1] == "S" &&
								strings.Split(lines[n+1], "")[j-1] == "M") {
									total_xmas_count++
								}
					} else if (
						strings.Split(lines[n-1], "")[j-1] == "S" &&
						strings.Split(lines[n+1], "")[j+1] == "M"){
							if (
								strings.Split(lines[n-1], "")[j+1] == "M" &&
								strings.Split(lines[n+1], "")[j-1] == "S"){
									total_xmas_count++
							} else if (
								strings.Split(lines[n-1], "")[j+1] == "S" &&
								strings.Split(lines[n+1], "")[j-1] == "M") {
									total_xmas_count++
								}
					}
				}

			}
		}
	}

	// Print out the total distance
	fmt.Println("Total occurences of X-MAS: ", total_xmas_count)
}