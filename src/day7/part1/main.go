package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Directory struct {
	name    string
	parent  *Directory
	size    int
	subDirs map[string]*Directory
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	root := Directory{
		name:    "/",
		parent:  nil,
		size:    0,
		subDirs: map[string]*Directory{},
	}

	currDir := &root

	for scanner.Scan() {
		words := strings.Fields(scanner.Text())

		if words[0] != "$" {
			if words[0] == "dir" {
				currDir.subDirs[words[1]] = &Directory{
					name:    words[1],
					parent:  currDir,
					size:    0,
					subDirs: map[string]*Directory{},
				}
			} else {
				size, _ := strconv.Atoi(words[0])
				currDir.size += size
			}
		} else if words[1] == "cd" {
			if words[2] == ".." {
				currDir = currDir.parent
			} else {
				currDir = currDir.subDirs[words[2]]
			}
		}
	}

	walk(&root)

	res := sum(&root, 0)

	fmt.Println(res)
}

func sum(root *Directory, n int) int {
	if len(root.subDirs) == 0 {
		if root.size <= 100000 {
			return n + root.size
		}
	}

	for _, v := range root.subDirs {
		if root.size <= 100000 {
			n += root.size
		}
		n = sum(v, n)
	}

	return n
}

func walk(root *Directory) {
	if len(root.subDirs) == 0 {
		return
	}

	for _, v := range root.subDirs {
		walk(v)
		root.size += v.size
	}
}
