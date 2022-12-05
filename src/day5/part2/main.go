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
	scanner.Scan()
	stacks := make([][]byte, (len(scanner.Text())+1)/4)

	for {
		idx := 0
		if scanner.Bytes()[1] >= '0' && scanner.Bytes()[1] <= '9' {
			break
		}
		for i := 1; i < len(scanner.Bytes()); i += 4 {
			if scanner.Bytes()[i] == 32 {
				idx++
				continue
			}

			stacks[idx] = append(stacks[idx], scanner.Bytes()[i])
			idx++
		}
		scanner.Scan()
	}
	scanner.Scan()

	// Parse all the moves
	for scanner.Scan() {
		move := strings.Fields(scanner.Text())
		amount, _ := strconv.Atoi(move[1])
		from, _ := strconv.Atoi(move[3])
		from--
		to, _ := strconv.Atoi(move[5])
		to--
		l := len(stacks[from]) - amount

		moved := make([]byte, amount)
		copy(moved, stacks[from][:amount])

		rest := make([]byte, l)
		copy(rest, stacks[from][amount:])

		stacks[from] = rest
		stacks[to] = append(moved, stacks[to]...)
	}

	res := make([]string, len(stacks))
	for i, v := range stacks {
		res[i] = string(v[0])
	}

	fmt.Println(res)

}
