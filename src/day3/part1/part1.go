package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var sum rune
	for scanner.Scan() {
		chars := map[rune]struct{}{}
		firstHalf := scanner.Text()[:len(scanner.Text())/2]
		secondHalf := scanner.Text()[len(scanner.Text())/2:]

		for _, r := range firstHalf {
			chars[r] = struct{}{}
		}

		for _, r := range secondHalf {
			if _, ok := chars[r]; !ok {
				continue
			}

			if r >= 'A' && r <= 'Z' {
				sum += 27 + r - 'A'
				break
			} else {
				sum += r - 'a' + 1
				break
			}
		}
	}

	fmt.Println(sum)

}
