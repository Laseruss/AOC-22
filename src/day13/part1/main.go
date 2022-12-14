package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type List []any

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	p := strings.Split(string(input), "\n\n")
	pairs := make([][]string, len(p))
	for i := range pairs {
		pair := strings.Split(p[i], "\n")
		pairs[i] = pair
	}
	res := 0
	for i, row := range pairs {
		left, right := parse(row[0]), parse(row[1])
		v := checkOrder(left, right)
		if v == 1 {
			res += i + 1
		} else if v == 0 {
			if len(left) < len(right) {
				res += i + 1
			}
		}
	}

	fmt.Println(res)
}

// Checking the order of each of the pairs

func checkOrder(left List, right List) int {
	if len(left) == 0 && len(right) != 0 {
		return 1
	} else if len(left) != 0 && len(right) == 0 {
		return -1
	} else if len(left) == 0 && len(right) == 0 {
		return 0
	}

	switch leftItem := left[0].(type) {
	case int:
		switch rightItem := right[0].(type) {
		case int:
			if leftItem < rightItem {
				return 1
			} else if rightItem < leftItem {
				return -1
			} else {
				return checkOrder(left[1:], right[1:])
			}
		case List:
			item := List{leftItem}
			v := checkOrder(item, rightItem)
			if v == 0 {
				return checkOrder(left[1:], right[1:])
			} else {
				return v
			}
		}
	case List:
		switch rightItem := right[0].(type) {
		case int:
			item := List{rightItem}
			v := checkOrder(leftItem, item)
			if v == 0 {
				return checkOrder(left[1:], right[1:])
			} else {
				return v
			}
		case List:
			v := checkOrder(leftItem, rightItem)
			if v == 0 {
				return checkOrder(left[1:], right[1:])
			} else {
				return v
			}
		}
	}

	return 0
}

// Parsing of the input

func parse(s string) List {
	pos := 1
	return parseList(s, &pos)
}

func parseList(s string, pos *int) List {
	res := List{}

	for s[*pos] != ']' {
		if s[*pos] == '[' {
			*pos++
			res = append(res, parseList(s, pos))
			*pos++
		}

		if isNumber(s[*pos]) {
			res = append(res, parseNumber(s, pos))
		}

		if s[*pos] == ',' {
			*pos++
		}

	}

	return res
}

func parseNumber(number string, pos *int) int {
	start := *pos
	for isNumber(number[*pos]) {
		*pos++
	}
	num, _ := strconv.Atoi(number[start:*pos])

	return num
}

// Helpers

func isNumber(c byte) bool {
	return c >= '0' && c <= '9'
}
