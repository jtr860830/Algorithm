package main

import (
	"fmt"
	"math"
)

type graph struct {
	to int     // To what node
	wt float64 // To node length
}

func floydWarshall(g [][]graph) [][]float64 {
	dist := make([][]float64, len(g))
	for i := range dist {
		di := make([]float64, len(g))
		for j := range di {
			di[j] = math.Inf(1)
		}
		di[i] = 0
		dist[i] = di
	}
	for u, graphs := range g {
		for _, v := range graphs {
			dist[u][v.to] = v.wt
		}
	}
	for k, dk := range dist {
		for _, di := range dist {
			for j, dij := range di {
				if d := di[k] + dk[j]; dij > d {
					di[j] = d
				}
			}
		}
	}
	return dist
}

func main() {
	gra := [][]graph{
		1: {{4, 1}, {5, 5}, {6, 2}}, // Node 1 to every node's length and so on
		2: {{1, 9}, {3, 3}, {4, 2}},
		3: {{4, 4}},
		4: {{3, 2}, {5, 3}},
		5: {{1, 3}},
		6: {{2, 1}},
	}

	dist := floydWarshall(gra)
	for _, d := range dist {
		fmt.Printf("%4g\n", d)
	}
}
