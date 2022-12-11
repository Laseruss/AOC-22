package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	cycle := 1
	register := 1
	sum := 0

	toSave := []int{20, 60, 100, 140, 180, 220}
	i := 0

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if i > 5 {
			break
		}
		words := strings.Fields(scanner.Text())

		if words[0] == "noop" {
			cycle++
			continue
		}

		reg, _ := strconv.Atoi(words[1])
		register += reg
		cycle += 2

		if cycle == toSave[i] {
			sum += register * toSave[i]
			i++
		} else if cycle > toSave[i] {
			sum += (register - reg) * toSave[i]
			i++
		}
	}

	fmt.Println(sum)
}
