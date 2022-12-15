package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type cord struct {
	x, y int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	coordinates := parseInput(scanner)

	minX := coordinates[0][0].x
	maxX := coordinates[0][0].x
	maxY := coordinates[0][0].y

	for _, row := range coordinates {
		for _, c := range row {
			if c.x > maxX {
				maxX = c.x
			}
			if c.x < minX {
				minX = c.x
			}
			if c.y > maxY {
				maxY = c.y
			}
		}
	}

	cave := make([][]string, maxY+1)
	for i := range cave {
		cave[i] = make([]string, maxX-minX+1)
	}

	for _, row := range coordinates {
		for i := range row[:len(row)-1] {
			if row[i].x == row[i+1].x {
				if row[i].y > row[i+1].y {
					for j := row[i+1].y; j <= row[i].y; j++ {
						cave[j][row[i].x-minX] = "#"
					}
				} else {
					for j := row[i].y; j <= row[i+1].y; j++ {
						cave[j][row[i].x-minX] = "#"
					}
				}
			} else {
				if row[i].x > row[i+1].x {
					for j := row[i+1].x; j <= row[i].x; j++ {
						cave[row[i].y][j-minX] = "#"
					}
				} else {
					for j := row[i].x; j <= row[i+1].x; j++ {
						cave[row[i].y][j-minX] = "#"
					}

				}
			}
		}
	}

	for i := range cave {
		for j := range cave[i] {
			if cave[i][j] == "" {
				cave[i][j] = "."
			}
		}
	}

	// Drop sand
	dropIdx := 500 - minX
	dropping := true
	grains := 0

	for {
		grains++
		y := 0
		x := dropIdx
		falling := true
		for falling {
			if x <= 0 || x == len(cave[0])-1 || y == len(cave)-1 {
				dropping = false
				break
			}
			// check if it can fall left and then right
			if cave[y+1][x] == "." {
				y++
			} else if cave[y+1][x-1] == "." {
				y++
				x--
			} else if cave[y+1][x+1] == "." {
				y++
				x++
			} else {
				cave[y][x] = "o"
				falling = false
			}
		}
		if !dropping {
			break
		}
	}

	fmt.Println(grains - 1)
}

func parseInput(scanner *bufio.Scanner) [][]cord {
	res := [][]cord{}
	for scanner.Scan() {
		line := []cord{}
		pairs := strings.Split(scanner.Text(), "->")
		for _, pair := range pairs {
			p := strings.Split(pair, ",")
			x, _ := strconv.Atoi(strings.TrimSpace(p[0]))
			y, _ := strconv.Atoi(strings.TrimSpace(p[1]))
			line = append(line, cord{
				x: x,
				y: y,
			})
		}
		res = append(res, line)
	}

	return res
}
