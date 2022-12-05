package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	rows := bytes.Split(input, []byte("\n"))
	sum := 0
	for i := 0; i < len(rows); i += 3 {
		if i+2 >= len(rows) {
			break
		}
		chars1 := map[byte]struct{}{}
		for _, b := range rows[i] {
			chars1[b] = struct{}{}
		}

		chars2 := map[byte]struct{}{}
		for _, b := range rows[i+1] {
			if _, ok := chars1[b]; ok {
				chars2[b] = struct{}{}
			}
		}

		for _, b := range rows[i+2] {
			if _, ok := chars2[b]; ok {
				if b >= 'A' && b <= 'Z' {
					sum += 27 + int(b-'A')
				} else {
					sum += int(b-'a') + 1
				}
				break
			}
		}
	}

	fmt.Println(sum)
}
