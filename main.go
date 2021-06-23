package main

import "fmt"

type Matrix [][]uint8

// count adjacent alive cells, skip over out of range and skip over
// the center(x,y).
func count_neighbors(x int, y int, input Matrix) int {
	count := 0
	for i := x - 1; i <= x+1; i++ {
		if i >= 0 && i < len(input) {
			for j := y - 1; j <= y+1; j++ {
				if (i != x || j != y) && j >= 0 && j < len(input[i]) {
					count += int(input[i][j])
				}
			}
		}
	}
	return count
}

// run a single generation
func generation(input Matrix) Matrix {

	// Need an output matrix that is sized to the input matrix.
	output := make([][]uint8, len(input))
	for i := range input {
		output[i] = make([]uint8, len(input[i]))
	}

	// run the rules to create the output generation
	for i, _ := range input {
		for j, _ := range input[i] {
			cell := input[i][j]
			alive := count_neighbors(i, j, input)
			switch cell {
			case 0:
				if alive == 3 {
					cell = 1
				}
			default:
				if alive < 2 || alive > 3 {
					cell = 0
				}
			}
			output[i][j] = cell
		}
	}
	return output
}

func main() {
	a := Matrix{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	for i := 0; i < 4; i++ {
		fmt.Println("gen:", i, " ", a)
		a = generation(a)
	}
}
