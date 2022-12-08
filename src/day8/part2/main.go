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

	maxView := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			view := calculateView(&grid, i, j)
			if view > maxView {
				maxView = view
			}
		}
	}

	fmt.Println(maxView)
}

func calculateView(grid *[][]byte, row, col int) int {
	cell := (*grid)[row][col]
	return checkLeft(grid, cell, row, col) *
		checkRight(grid, cell, row, col) *
		checkUp(grid, cell, row, col) *
		checkDown(grid, cell, row, col)
}

func checkLeft(grid *[][]byte, cell byte, row, col int) int {
	res := 0
	for i := col - 1; i >= 0; i-- {
		if (*grid)[row][i] >= cell {
			res++
			break
		}
		res++
	}

	return res
}

func checkRight(grid *[][]byte, cell byte, row, col int) int {
	res := 0
	for i := col + 1; i < len((*grid)[0]); i++ {
		if (*grid)[row][i] >= cell {
			res++
			break
		}
		res++
	}

	return res
}

func checkUp(grid *[][]byte, cell byte, row, col int) int {
	res := 0
	for i := row - 1; i >= 0; i-- {
		if (*grid)[i][col] >= cell {
			res++
			break
		}
		res++
	}

	return res
}

func checkDown(grid *[][]byte, cell byte, row, col int) int {
	res := 0
	for i := row + 1; i < len(*grid); i++ {
		if (*grid)[i][col] >= cell {
			res++
			break
		}
		res++
	}

	return res

}
