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

type obj struct {
	start, end int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	coordinates := parseInput(scanner)

	/*
		objects := map[int][]obj{}

		for _, row := range coordinates {
			for i := 0; i < len(row)-1; i++ {
				if row[i].x != row[i+1].x {
					if row[i].x < row[i+1].x {
						objects[row[i].y] = append(objects[row[i].y], obj{row[i].x, row[i+1].x})
					} else {
						objects[row[i].y] = append(objects[row[i].y], obj{row[i+1].x, row[i].x})
					}
				}
			}
		}
	*/

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

	cave := make([][]string, maxY+2)
	for i := range cave {
		cave[i] = make([]string, 1000)
	}

	for _, row := range coordinates {
		for i := range row[:len(row)-1] {
			if row[i].x == row[i+1].x {
				if row[i].y > row[i+1].y {
					for j := row[i+1].y; j <= row[i].y; j++ {
						cave[j][row[i].x] = "#"
					}
				} else {
					for j := row[i].y; j <= row[i+1].y; j++ {
						cave[j][row[i].x] = "#"
					}
				}
			} else {
				if row[i].x > row[i+1].x {
					for j := row[i+1].x; j <= row[i].x; j++ {
						cave[row[i].y][j] = "#"
					}
				} else {
					for j := row[i].x; j <= row[i+1].x; j++ {
						cave[row[i].y][j] = "#"
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

	dropIdx := 500
	dropping := true
	grains := 0

	for {
		y := 0
		x := dropIdx
		falling := true
		for falling {
			if cave[0][500] == "o" {
				dropping = false
				break
			}
			if y == len(cave)-1 {
				cave[y][x] = "o"
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
		grains++
	}

	fmt.Println(grains)

	/*
		sumObs := 0

		for _, row := range cave {
			fmt.Println(strings.Join(row, " "))
			for _, v := range row {
				if v == "#" {
					sumObs++
				}
			}
		}

		sumBlocked := 0
		for depth, blockers := range objects {
			d := depth
			for _, block := range blockers {
				w := block.end - block.start + 1
				for d < maxY+1 && w > 2 {
					d++
					c := true
					for _, v := range objects[d] {
						if (v.start >= block.start && v.start <= block.end) || (v.end >= block.start && v.end <= block.end) {
							c = false
						}
					}
					if !c {
						break
					}
					w = w - 2
					sumBlocked += w
				}
			}
		}
	*/

}

func totalGrains(y int) int {
	res := 0
	for i := 1; i <= y+2; i++ {
		res += i*2 - 1
	}

	return res
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
