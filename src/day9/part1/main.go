package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type position struct {
	x int
	y int
}

func main() {
	head := position{
		x: 0,
		y: 0,
	}

	tail := position{
		x: 0,
		y: 0,
	}

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	res := map[string]struct{}{}
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		direction := words[0]
		steps := words[1]

		res = moveHead(res, &head, &tail, direction, steps)
	}

	fmt.Println(len(res))
}

func moveHead(places map[string]struct{}, head *position, tail *position, direction, steps string) map[string]struct{} {
	s, _ := strconv.Atoi(steps)
	for i := 0; i < s; i++ {
		switch direction {
		case "R":
			head.x += 1
			break
		case "L":
			head.x -= 1
			break
		case "U":
			head.y += 1
			break
		case "D":
			head.y -= 1
			break
		}

		if !isTouching(*head, *tail) {
			moveTail(head, tail)
		}
		places[strconv.Itoa(tail.x)+", "+strconv.Itoa(tail.y)] = struct{}{}

	}

	return places

}

func isTouching(head, tail position) bool {
	return (tail.x-1 == head.x || tail.x+1 == head.x || tail.x == head.x) &&
		(tail.y-1 == head.y || tail.y+1 == head.y || tail.y == head.y)
}

func moveTail(head, tail *position) {
	if head.x != tail.x && head.y != tail.y {
		if head.x-2 == tail.x {
			tail.x = head.x - 1
			tail.y = head.y
		} else if head.x+2 == tail.x {
			tail.x = head.x + 1
			tail.y = head.y
		}
		if head.y-2 == tail.y {
			tail.y = head.y - 1
			tail.x = head.x
		} else if head.y+2 == tail.y {
			tail.y = head.y + 1
			tail.x = head.x
		}
		return
	}
	if head.x != tail.x {
		if head.x > tail.x {
			tail.x = head.x - 1
		} else {
			tail.x = head.x + 1
		}
		return
	}

	if head.y > tail.y {
		tail.y = head.y - 1
	} else {
		tail.y = head.y + 1
	}
}
