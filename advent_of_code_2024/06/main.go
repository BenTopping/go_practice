package main

import (
	"bufio"
	"fmt"
	"os"
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
	file, err := os.Open("inputs/06.txt")

	// If we don't have an input file return an error
	if err != nil {
		fmt.Println("Cannot find input file, please add your input file to inputs/06.txt")
		return
	}

	// Ensure we close the file
	defer file.Close()

	// File reader helper
	reader := bufio.NewReader(file)

	var lines = make([][]string, 0)
	guard_pos_x, guard_pos_y := 0, 0
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		for n := 0; n < len(line); n++ {
			// Find the guard character and gets its position
			if string(line)[n] == '^' {
				guard_pos_x, guard_pos_y = n, len(lines)
			}
		}
		lines = append(lines, strings.Split(string(line), ""))
	}

	// Possible directions
	guard_directions := [][]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	// Initial direction
	guard_direction := 0
	// Positions visited
	type Pos struct {
		X, Y int
	}
	// Keep track of positions visited
	positions_visited := map[Pos]int{}
	// While the guard is still in the map loop
	for guard_pos_x < len(lines[0]) && guard_pos_y < len(lines) {
		// Get the next positions
		next_pos_x := guard_pos_x + guard_directions[guard_direction][0]
		next_pos_y := guard_pos_y + guard_directions[guard_direction][1]

		// Check the next positions aren't out of bound and not an obstacle
		if next_pos_x < len(lines[0]) && next_pos_y < len(lines) && lines[next_pos_y][next_pos_x] == "#" {
			// If the guard is at an obstacle turn and loop back to start of direction list if fully turned (3)
			if guard_direction == 3 {
				guard_direction = 0
			} else {
				guard_direction++
			}
		} else {
			// Add the positions to visited list and move the guard
			positions_visited[Pos{guard_pos_x, guard_pos_y}] = 1
			guard_pos_x = next_pos_x
			guard_pos_y = next_pos_y
		}
	}

	fmt.Println(len(positions_visited))
}

func part2() {
	// Read the input file
	file, err := os.Open("inputs/06.txt")

	// If we don't have an input file return an error
	if err != nil {
		fmt.Println("Cannot find input file, please add your input file to inputs/06.txt")
		return
	}

	// Ensure we close the file
	defer file.Close()

	// File reader helper
	reader := bufio.NewReader(file)

	var lines = make([][]string, 0)
	guard_pos_x, guard_pos_y := 0, 0
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		for n := 0; n < len(line); n++ {
			// Find the guard character and gets its position
			if string(line)[n] == '^' {
				guard_pos_x, guard_pos_y = n, len(lines)
			}
		}
		lines = append(lines, strings.Split(string(line), ""))
	}

	// Possible directions
	guard_directions := [][]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	// Positions visited - note includes direction so we can track loops (revisted in same direction)
	type Pos struct {
		X, Y, direction int
	}
	// Keep track of how many obstacles cause loops
	recurrsive_obstacle_count := 0
	// For each line
	for n := 0; n < len(lines); n++ {
		// For each position on the line replace it with an obstacle
		for j := 0; j < len(lines[n]); j++ {
			initial_line_value := lines[n][j]
			// Keep track of the initial guard positions so we can restore them after each loop
			initial_guard_x, initial_guard_y := guard_pos_x, guard_pos_y
			// Set the current position to an obstacle if its not the guard position
			if (lines[n][j] != "^") {lines[n][j] = "#"}
			// Initial direction
			guard_direction := 0
			// Keep track of positions visited
			positions_visited := map[Pos]int{}
			// While the guard is still in the map loop
			for 0 <= guard_pos_x && guard_pos_x < len(lines[0]) && 0 <= guard_pos_y && guard_pos_y < len(lines) {
				// Get the next positions
				next_pos_x := guard_pos_x + guard_directions[guard_direction][0]
				next_pos_y := guard_pos_y + guard_directions[guard_direction][1]

				// Check the next positions aren't out of bound and not an obstacle
				if 0 <= next_pos_x && 0 <= next_pos_y && next_pos_x < len(lines[0]) && next_pos_y < len(lines) && lines[next_pos_y][next_pos_x] == "#" {
					// If the guard is at an obstacle turn and loop back to start of direction list if fully turned (3)
					if guard_direction == 3 {
						guard_direction = 0
					} else {
						guard_direction++
					}
				} else {
					// Add the positions to visited list and move the guard
					positions_visited[Pos{guard_pos_x, guard_pos_y, guard_direction}]++
					// If the positiion has been seen before add loop to count and go to the next position
					if positions_visited[Pos{guard_pos_x, guard_pos_y, guard_direction}] > 1 {
						recurrsive_obstacle_count++
						break
					}
					guard_pos_x = next_pos_x
					guard_pos_y = next_pos_y
				}
			}
			// Restore the guard position and line value
			guard_pos_x, guard_pos_y = initial_guard_x, initial_guard_y
			lines[n][j] = initial_line_value
		}
	}

	fmt.Println(recurrsive_obstacle_count)
}
