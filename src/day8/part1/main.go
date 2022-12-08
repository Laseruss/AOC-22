package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	grid := bytes.Split(input, []byte("\n"))
	grid = grid[:len(grid)-1]
	visibleGrid := createGrid(len(grid), len(grid[0]))

	// Checking the rows from left and from right
	// If a tree is visible setting it to true in visibleGrid
	for i := 0; i < len(grid); i++ {
		startMax, endMax := grid[i][0], grid[i][len(grid[i])-1]
		visibleGrid[i][0], visibleGrid[i][len(grid[i])-1] = true, true

		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] > startMax {
				visibleGrid[i][j] = true
				startMax = grid[i][j]
			}
			if grid[i][len(grid[i])-1-j] > endMax {
				visibleGrid[i][len(grid[i])-1-j] = true
				endMax = grid[i][len(grid[i])-1-j]
			}
		}
	}

	for j := 0; j < len(grid[0]); j++ {
		startMax, endMax := grid[0][j], grid[len(grid)-1][j]
		visibleGrid[0][j], visibleGrid[len(grid)-1][j] = true, true

		for i := 0; i < len(grid); i++ {
			if grid[i][j] > startMax {
				visibleGrid[i][j] = true
				startMax = grid[i][j]
			}
			if grid[len(grid)-1-i][j] > endMax {
				visibleGrid[len(grid)-1-i][j] = true
				endMax = grid[len(grid)-1-i][j]
			}
		}
	}

	sum := 0
	for _, row := range visibleGrid {
		for _, cell := range row {
			if cell {
				sum++
			}
		}
	}

	fmt.Println(sum)
}

func createGrid(l, w int) [][]bool {
	grid := make([][]bool, l)
	for i := range grid {
		grid[i] = make([]bool, w)
	}

	return grid
}
