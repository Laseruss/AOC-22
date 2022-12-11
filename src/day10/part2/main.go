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
	screen := [6][40]string{}

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		words := strings.Fields(scanner.Text())

		if words[0] == "noop" {
			row := cycle / 40
			col := cycle % 40
			if col == register-1 || col == register || col == register+1 {
				screen[row][col] = "#"
			} else {
				screen[row][col] = "."
			}
			cycle++
		} else {
			row := cycle / 40
			col := cycle % 40
			if col == register-1 || col == register || col == register+1 {
				screen[row][col] = "#"
			} else {
				screen[row][col] = "."
			}
			cycle++
			row = cycle / 40
			col = cycle % 40
			if col == register-1 || col == register || col == register+1 {
				screen[row][col] = "#"
			} else {
				screen[row][col] = "."
			}
			reg, _ := strconv.Atoi(words[1])
			register += reg
			cycle++
		}

	}

	for _, row := range screen {
		fmt.Println(row)
	}

}
