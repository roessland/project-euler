package main

import "fmt"
import "strings"
import "strconv"
import "io/ioutil"
import "github.com/roessland/gopkg/digraph"

var N int
var mat [][]float64

func ReadMatrix(filename string) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err.Error())
	}
	s := string(buf)
	lines := strings.Split(s, "\n")
	N = len(lines) - 1
	lines = lines[0:N]

	mat = make([][]float64, N)
	for i, line := range lines {
		mat[i] = make([]float64, N)
		for j, numStr := range strings.Split(line, ",") {
			num, _ := strconv.ParseFloat(numStr, 64)
			mat[i][j] = num
		}
	}
}

func Idx(i, j int) int {
	return N*i + j
}

func main() {
	ReadMatrix("p082_matrix.txt")

	graph := digraph.Graph{make([]digraph.Node, N*N+2)}
	A := N * N   // supersource
	B := N*N + 1 // supersink

	// Set up supersource
	graph.Nodes[A].Neighbors = make([]digraph.Edge, N)
	for i := 0; i < N; i++ {
		graph.Nodes[A].Neighbors[i].To = Idx(i, 0)
		graph.Nodes[A].Neighbors[i].Weight = mat[i][0]
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			nodeIdx := Idx(i, j)

			// Add edge upwards if possible
			if i != 0 {
				graph.Nodes[nodeIdx].Neighbors =
					append(graph.Nodes[nodeIdx].Neighbors, digraph.Edge{Idx(i-1, j), mat[i-1][j]})
			}

			// Add edge downwards if possible
			if i != N-1 {
				graph.Nodes[nodeIdx].Neighbors =
					append(graph.Nodes[nodeIdx].Neighbors, digraph.Edge{Idx(i+1, j), mat[i+1][j]})
			}

			// Add edge rightwards if possible
			if j != N-1 {
				graph.Nodes[nodeIdx].Neighbors =
					append(graph.Nodes[nodeIdx].Neighbors, digraph.Edge{Idx(i, j+1), mat[i][j+1]})
			}

		}
	}

	// Set up supersink
	for i := 0; i < N; i++ {
		graph.Nodes[Idx(i, N-1)].Neighbors =
			append(graph.Nodes[Idx(i, N-1)].Neighbors, digraph.Edge{B, 0.0})
	}

	dist, _ := digraph.Dijkstra(graph, A)
	fmt.Printf("%v\n", dist[B])
}
