package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	cycle := 0
	register := 1
	screen := [6][]string{}
	for i := range screen {
		screen[i] = make([]string, 40)
	}

	file, _ := os.Open("test.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if cycle >= 240 {
			break
		}
		words := strings.Fields(scanner.Text())
		col := cycle % 40
		row := cycle / 40

		if words[0] == "noop" {
			if cycle == register-1 || cycle == register || cycle == register+1 {
				screen[row][col] = "#"
			} else {
				screen[row][col] = "."
			}
			cycle++
			continue
		}

		reg, _ := strconv.Atoi(words[1])
		if col == register-1 || col == register || col == register+1 {
			screen[row][col] = "#"
		} else {
			screen[row][col] = "."
		}
		cycle++
		col = cycle % 40
		row = cycle / 40
		if col == register-1 || col == register || col == register+1 {
			screen[row][col] = "#"
		} else {
			screen[row][col] = "."
		}
		cycle++
		register += reg

	}

	for _, row := range screen {
		fmt.Println(row)
	}

}
