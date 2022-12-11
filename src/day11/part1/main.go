package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Monkey struct {
	items     []int
	operation func(int) int
	test      func(int) bool
	ifPass    int
	ifNotPass int
}

func (m Monkey) multiply(x int) func(int) int {
	return func(y int) int {
		return x * y
	}
}

func (m Monkey) multiplySelf() func(int) int {
	return func(x int) int {
		return x * x
	}
}

func (m Monkey) add(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func (m Monkey) mod(x int) func(int) bool {
	return func(y int) bool {
		return y%x == 0
	}
}

func main() {
	monkeys := parseInput("input.txt")
	inspected := make([]int, len(monkeys))
	for lap := 0; lap < 20; lap++ {
		for i := range monkeys {
			for _, item := range monkeys[i].items {
				inspected[i]++
				worry := monkeys[i].operation(item) / 3
				fmt.Println(worry)
				didPass := monkeys[i].test(worry)
				if didPass {
					monkeys[monkeys[i].ifPass].items = append(monkeys[monkeys[i].ifPass].items, worry)
				} else {
					monkeys[monkeys[i].ifNotPass].items = append(monkeys[monkeys[i].ifNotPass].items, worry)
				}

			}
			monkeys[i].items = []int{}
		}
	}
	fmt.Println(inspected)
}

func parseInput(filename string) []Monkey {
	res := []Monkey{}
	text, _ := os.ReadFile(filename)
	monkeys := strings.Split(string(text), "\n\n")
	for _, m := range monkeys {
		monkey := Monkey{}
		monkeyAttr := strings.Split(m, "\n")

		// Parsing the items
		dStr := strings.Fields(monkeyAttr[1])[2:]
		for _, v := range dStr {
			v := strings.TrimSuffix(v, ",")
			x, _ := strconv.Atoi(v)
			monkey.items = append(monkey.items, x)
		}

		// parsing the function
		fnStr := strings.Fields(monkeyAttr[2])
		operator := fnStr[len(fnStr)-2]
		num := fnStr[len(fnStr)-1]
		dig, _ := strconv.Atoi(num)

		if operator == "*" {
			monkey.operation = monkey.multiply(dig)
		} else if operator == "+" {
			monkey.operation = monkey.add(dig)
		}
		if num == "old" {
			monkey.operation = monkey.multiplySelf()
		}

		// parsing the test
		testStr := strings.Fields(monkeyAttr[3])
		num = testStr[len(testStr)-1]
		dig, _ = strconv.Atoi(num)

		monkey.test = monkey.mod(dig)

		ifTrueStr := strings.Fields(monkeyAttr[4])
		num = ifTrueStr[len(ifTrueStr)-1]
		dig, _ = strconv.Atoi(num)
		monkey.ifPass = dig

		ifFalseStr := strings.Fields(monkeyAttr[5])
		num = ifFalseStr[len(ifFalseStr)-1]
		dig, _ = strconv.Atoi(num)
		monkey.ifNotPass = dig

		res = append(res, monkey)
	}

	return res
}
