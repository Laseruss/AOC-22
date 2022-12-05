package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	ranges := [][]int{}
	for scanner.Scan() {
		pair := strings.Split(scanner.Text(), ",")
		for i := range pair {
			nums := make([]int, 2)
			s := strings.Split(pair[i], "-")
			for i, d := range s {
				n, _ := strconv.Atoi(d)
				nums[i] = n
			}
			ranges = append(ranges, nums)
		}
	}

	res := 0
	for i := 0; i < len(ranges); i += 2 {
		if ranges[i][0] <= ranges[i+1][1] && ranges[i+1][0] <= ranges[i][1] {
			res++
		}
	}

	fmt.Println(res)
}
