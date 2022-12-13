package main

import (
	"bytes"
	"fmt"
	"math"
	"os"
)

type vertex struct {
	value     byte
	cost      int
	neighbors []*vertex
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	inputRows := bytes.Split(input, []byte("\n"))
	inputRows = inputRows[:len(inputRows)-1]

	graph, start, end := buildGraph(&inputRows)

	for i, row := range graph {
		for j, vertex := range row {
			if i > 0 {
				// check up neighbor
				if checkNeighbor(vertex.value, graph[i-1][j].value) {
					graph[i][j].neighbors = append(graph[i][j].neighbors, &graph[i-1][j])
				}
			}
			if i < len(graph)-1 {
				// check down neighbor
				if checkNeighbor(vertex.value, graph[i+1][j].value) {
					graph[i][j].neighbors = append(graph[i][j].neighbors, &graph[i+1][j])

				}
			}
			if j > 0 {
				// check left neighbor
				if checkNeighbor(vertex.value, graph[i][j-1].value) {
					graph[i][j].neighbors = append(graph[i][j].neighbors, &graph[i][j-1])

				}
			}
			if j < len(row)-1 {
				// check right neighbor
				if checkNeighbor(vertex.value, graph[i][j+1].value) {
					graph[i][j].neighbors = append(graph[i][j].neighbors, &graph[i][j+1])

				}
			}
		}
	}

	currentNode := start
	unvisitedNodes := map[*vertex]struct{}{
		currentNode: {},
	}
	visited := map[*vertex]struct{}{}

	for len(unvisitedNodes) > 0 {
		for _, vertice := range currentNode.neighbors {
			if _, ok := visited[vertice]; !ok {
				unvisitedNodes[vertice] = struct{}{}
			}

			cost := currentNode.cost + 1
			if cost < vertice.cost {
				vertice.cost = cost
			}
		}

		visited[currentNode] = struct{}{}
		delete(unvisitedNodes, currentNode)

		currentNode, err = findLowest(unvisitedNodes)
		if err != nil {
			break
		}
	}

	fmt.Println(start, end)
}

func removeCurrent(current *vertex, unvisited []*vertex) []*vertex {
	idx := 0
	for i, v := range unvisited {
		if v == current {
			fmt.Println("hello")
			idx = i
			break
		}
	}

	res := make([]*vertex, 0, len(unvisited)-1)
	res = append(res, unvisited[:idx]...)
	res = append(res, unvisited[idx+1:]...)

	return res
}

func findLowest(vertices map[*vertex]struct{}) (*vertex, error) {
	if len(vertices) == 0 {
		return nil, fmt.Errorf("oops")
	}

	min := math.MaxInt
	var minV *vertex
	for k := range vertices {
		if k.cost < min {
			min = k.cost
			minV = k
		}
	}

	return minV, nil
}

func buildGraph(input *[][]byte) ([][]vertex, *vertex, *vertex) {
	graph := [][]vertex{}
	var start *vertex
	var end *vertex

	for i, row := range *input {
		r := make([]vertex, len(row))
		graph = append(graph, r)
		for j, b := range row {
			v := vertex{
				value:     b,
				cost:      math.MaxInt,
				neighbors: []*vertex{},
			}
			graph[i][j] = v
			if b == 'S' {
				start = &graph[i][j]
			} else if b == 'E' {
				end = &graph[i][j]
			}
		}
	}

	start.cost = 0
	start.value = 'a'
	end.value = 'z'

	return graph, start, end
}

func checkNeighbor(value, possible byte) bool {
	return value >= possible || value+1 == possible
}
