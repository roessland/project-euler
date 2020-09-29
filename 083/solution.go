// Runs instantly, using Dijkstra and priority queue
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
	ReadMatrix("p083_matrix.txt")
	//ReadMatrix("smallmatrix.txt")

	graph := digraph.Graph{make([]digraph.Node, N*N)}

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

			// Add edge rightwards if possible
			if j != 0 {
				graph.Nodes[nodeIdx].Neighbors =
					append(graph.Nodes[nodeIdx].Neighbors, digraph.Edge{Idx(i, j-1), mat[i][j-1]})
			}

		}
	}

	dist, _ := digraph.Dijkstra(graph, Idx(0, 0))
	fmt.Printf("%v\n", mat[0][0]+dist[Idx(N-1, N-1)])
}
