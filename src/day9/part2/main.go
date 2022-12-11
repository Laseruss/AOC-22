// Not actually working, getting 2324 but want 2557
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type position struct {
	x, y int
}

func buildRope(n int) []*position {
	rope := make([]*position, n)
	for i := 0; i < n; i++ {
		rope[i] = &position{x: 0, y: 0}
	}

	return rope
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	rope := buildRope(10)
	tailPos := map[string]struct{}{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		direction := words[0]
		steps, _ := strconv.Atoi(words[1])

		tailPos = moveHead(tailPos, rope, direction, steps)
	}

	fmt.Println(len(tailPos))
}

func moveHead(places map[string]struct{}, rope []*position, direction string, steps int) map[string]struct{} {
	for i := 0; i < steps; i++ {
		switch direction {
		case "R":
			rope[0].x += 1
			break
		case "L":
			rope[0].x -= 1
			break
		case "U":
			rope[0].y += 1
			break
		case "D":
			rope[0].y -= 1
			break
		}

		for i := 1; i < 10; i++ {
			if !isTouching(*rope[i-1], *rope[i]) {
				moveTail(rope[i-1], rope[i])
			}
			places[strconv.Itoa(rope[9].x)+", "+strconv.Itoa(rope[9].y)] = struct{}{}
		}
	}

	return places
}

func isTouching(head, tail position) bool {
	return (tail.x-1 == head.x || tail.x+1 == head.x || tail.x == head.x) &&
		(tail.y-1 == head.y || tail.y+1 == head.y || tail.y == head.y)
}

func moveTail(head, tail *position) {
	if int(math.Abs(float64(head.x-tail.x)+math.Abs(float64(head.y-tail.y)))) == 3 {
		if head.x == tail.x-2 {
			tail.x -= 1
			tail.y += head.y - tail.y
		}
		if head.x == tail.x+2 {
			tail.x += 1
			tail.y += head.y - tail.y
		}
		if head.y == tail.y-2 {
			tail.y -= 1
			tail.x += head.x - tail.x
		}
		if head.y == tail.y+2 {
			tail.y += 1
			tail.x += head.x - tail.x
		}
	} else if int(math.Abs(float64(head.x-tail.x))) > 1 || int(math.Abs(float64(head.y-tail.y))) > 1 {
		if head.x == tail.x-2 {
			tail.x -= 1
		}
		if head.x == tail.x+2 {
			tail.x += 1
		}
		if head.y == tail.y-2 {
			tail.y -= 1
		}
		if head.y == tail.y+2 {
			tail.y += 1
		}

	}
}
