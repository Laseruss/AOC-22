package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	points := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
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
		choices := strings.Split(scanner.Text(), " ")
		score += points[choices[1]]
		score += checkWinner(choices[0], choices[1])
	}

	fmt.Println(score)
}

func checkWinner(opp, you string) int {
	if (you == "X" && opp == "C") || (you == "Y" && opp == "A") || (you == "Z" && opp == "B") {
		return 6
	} else if (you == "X" && opp == "A") || (you == "Y" && opp == "B") || (you == "Z" && opp == "C") {
		return 3
	} else {
		return 0
	}

}
