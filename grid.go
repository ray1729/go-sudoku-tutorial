package main

import (
	"fmt"
	"strings"
)

type Grid [81]int

func (g Grid) ElementAt(row, col int) int {
	return g[col+row*9]
}

func (g Grid) WithElementAt(row, col int, val int) Grid {
	g[col+row*9] = val
	return g
}

func (g Grid) String() string {
	var sb strings.Builder
	for r := 0; r < 9; r++ {
		fmt.Fprintf(&sb, "%d %d %d | %d %d %d | %d %d %d\n",
			g.ElementAt(r, 0), g.ElementAt(r, 1), g.ElementAt(r, 2),
			g.ElementAt(r, 3), g.ElementAt(r, 4), g.ElementAt(r, 5),
			g.ElementAt(r, 6), g.ElementAt(r, 7), g.ElementAt(r, 8))
		if r == 2 || r == 5 {
			sb.WriteString("------+-------+------\n")
		}
	}
	return sb.String()
}

func (g Grid) Neighbours(row, col int) IntSet {
	neighbours := IntSet{}

	// Neighbouring row
	for c := 0; c < 9; c++ {
		if v := g.ElementAt(row, c); v != 0 {
			neighbours.Insert(v)
		}
	}

	// Neighbouring col
	for r := 0; r < 9; r++ {
		if v := g.ElementAt(r, col); v != 0 {
			neighbours.Insert(v)
		}
	}

	// Top-left corner of the 3x3 subgrid
	tlRow := (row / 3) * 3
	tlCol := (col / 3) * 3

	// Neighbouring subgrid
	for r := tlRow; r < tlRow+3; r++ {
		for c := tlCol; c < tlCol+3; c++ {
			if v := g.ElementAt(r, c); v != 0 {
				neighbours.Insert(v)
			}
		}
	}

	return neighbours
}

func (g Grid) FirstEmptyCell() int {
	for i, v := range g {
		if v == 0 {
			return i
		}
	}
	return -1
}

func (g Grid) Candidates(row, col int) []int {
	neighbours := g.Neighbours(row, col)
	var candidates []int
	for _, v := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9} {
		if !neighbours.Contains(v) {
			candidates = append(candidates, v)
		}
	}
	return candidates
}

func (g Grid) Solve() (Grid, error) {
	i := g.FirstEmptyCell()
	if i == -1 {
		// No unfilled cells: we have found a solution!
		return g, nil
	}
	row := i / 9
	col := i % 9
	candidates := g.Candidates(row, col)
	// Try each of the candidates in turn
	for _, v := range candidates {
		result, err := g.WithElementAt(row, col, v).Solve()
		if err == nil {
			return result, nil
		}
	}
	// There were no valid candidates
	return g, fmt.Errorf("No solutions found")
}

func ParseGrid(data []byte) (Grid, error) {
	G := Grid{}
	i := 0
	for _, x := range data {
		x := int(x) - int('0')
		if x >= 0 && x <= 9 {
			G[i] = x
			i++
		}
	}
	if i < len(G) {
		return G, fmt.Errorf("Invalid input: found %d elements, expected %d", i, len(G))
	}
	return G, nil
}
