package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("-- Part 1 --")
	part1()
	fmt.Println("-- Part 2 --")
	part2()
}

type Node struct {
	X, Y int
}
func part1() {
	// Read the input file
	file, err := os.Open("inputs/08.txt")

	// If we don't have an input file return an error
	if err != nil {
		fmt.Println("Cannot find input file, please add your input file to inputs/08.txt")
		return
	}

	// Ensure we close the file
	defer file.Close()

	// File reader helper
	reader := bufio.NewReader(file)

	// Create a map of total to its component values
	var frequencies = make(map[string][]Node)
	var y_pos = 1
	var max_x = 0
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		// bit of a hacky way to figure out and store max x
		max_x = len(line) + 1
		for n := 0; n < len(line); n++ {
			if (string(line[n]) == ".") { continue }
			
			frequencies[string(line[n])] = append(frequencies[string(line[n])], Node{X: n+1, Y: y_pos})
		}
		y_pos++
	}

	var anti_nodes = make(map[Node]int)

	for _, frequency := range frequencies {
		for _, node := range frequency {
			for n := 0; n < len(frequency); n++ {
				// If on the current node skip
				if (frequency[n] == node) {
					continue
				}

				var anti_node = Node{}
				// Anti node position is 2 times current node minus paired node
				anti_node.X = 2 * node.X - frequency[n].X
				anti_node.Y = 2 * node.Y - frequency[n].Y

				// If anti node is in a valid position add it to the set
				if anti_node.X > 0 && anti_node.X < max_x && anti_node.Y > 0 && anti_node.Y < y_pos {
					anti_nodes[anti_node] = 1
				}

			}
		}
	}

	fmt.Println("Number of unique antinodes: ", len(anti_nodes))
}

func part2() {}