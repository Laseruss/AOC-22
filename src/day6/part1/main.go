package main

import (
	"fmt"
	"os"
)

func main() {
	text, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	chars := map[byte]int{}
	res := 0

	for i := 0; i < len(text); i++ {
		if idx, ok := chars[text[i]]; ok {
			i = idx
			chars = map[byte]int{}
			continue
		}

		chars[text[i]] = i

		if len(chars) == 4 {
			res = i + 1
			break
		}
	}

	fmt.Println(res)
}
