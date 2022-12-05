package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	topThree := make([]int, 3)
	currCals := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			if currCals > topThree[0] {
				topThree[0] = currCals
				for i := 0; i < len(topThree)-1; i++ {
					if topThree[i] > topThree[i+1] {
						topThree[i], topThree[i+1] = topThree[i+1], topThree[i]
					}
				}
			}
			currCals = 0
			continue
		}

		cal, err := strconv.Atoi(text)
		if err != nil {
			panic(err)
		}
		currCals += cal
	}

	res := 0
	for _, v := range topThree {
		res += v
	}

	fmt.Println(res)
}
