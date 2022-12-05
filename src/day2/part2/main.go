package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scores := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}

	options := map[string]map[string]string{
		"wins": {
			"A": "B",
			"B": "C",
			"C": "A",
		},
		"lose": {
			"B": "A",
			"C": "B",
			"A": "C",
		},
	}

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	score := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			break
		}

		choices := strings.Fields(scanner.Text())
		if choices[1] == "X" {
			score += scores[options["lose"][choices[0]]]
		} else if choices[1] == "Y" {
			score += 3 + scores[choices[0]]
		} else {
			score += 6 + scores[options["wins"][choices[0]]]
		}
	}

	fmt.Println(score)
}
