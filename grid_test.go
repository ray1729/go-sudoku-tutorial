package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestElementAt(t *testing.T) {
	G := Grid{} // an empty grid
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			v := G.ElementAt(row, col)
			if v != 0 {
				t.Errorf("ElementAt(%d, %d) = %d, expected 0", row, col, v)
			}
		}
	}
}

func TestElementAtPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("ElementAt did not panic when out of bounds")
		}
	}()
	G := Grid{}
	G.ElementAt(10, 10)
}

func TestWithElementAt(t *testing.T) {
	G := Grid{}
	for i := 0; i < 10; i++ {
		row := rand.Intn(9)
		col := rand.Intn(9)
		val := rand.Intn(10)
		G = G.WithElementAt(row, col, val)
		if G.ElementAt(row, col) != val {
			t.Errorf("WithElementAt(%d, %d, %d) did not return expected value", row, col, val)
		}
	}
}

func TestParseGrid(t *testing.T) {
	data := []byte(`010020706
700913040
380004001
000007010
500109003
090500000
200300094
040762005
105090070`)
	G, err := ParseGrid(data)
	if err != nil {
		t.Errorf("Unexpected error parsing grid: %v", err)
	}
	// Check that row 2 is as expected
	row := 2
	expected := []int{3, 8, 0, 0, 0, 4, 0, 0, 1}
	for col, want := range expected {
		got := G.ElementAt(row, col)
		if got != want {
			t.Errorf("ElementAt(%d,%d) = %d, expected %d", row, col, got, want)
		}
	}
	fmt.Println(G)
}

func TestParseGridError(t *testing.T) {
	data := []byte(`01002070
700913040
380004001
000007010
500109003
090500000
200300094
040762005
105090070`)
	_, err := ParseGrid(data)
	if err == nil {
		t.Errorf("ParseGrid() did not return an expected error")
	}
}

var tests = []struct {
	data           []byte
	grid           Grid
	firstEmptyCell int
	neighbours     []int
	candidates     []int
}{
	{
		[]byte("010020706700913040380004001000007010500109003090500000200300094040762005105090070"),
		Grid{},
		0,
		[]int{1, 2, 3, 5, 6, 7, 8},
		[]int{4, 9},
	},
	{
		[]byte("000000800785090006040700002068300051400000007570009320900002060800060719006000000"),
		Grid{},
		0,
		[]int{4, 5, 7, 8, 9},
		[]int{1, 2, 3, 6},
	},
	{
		[]byte("320000809400600027005000040000401000002000900000502000080000100640003008107000032"),
		Grid{},
		2,
		[]int{2, 3, 4, 5, 7, 8, 9},
		[]int{1, 6},
	},
	{
		[]byte("000008000040100390000900561006094000003000200000320800372001000095003020000200000"),
		Grid{},
		0,
		[]int{3, 4, 8},
		[]int{1, 2, 5, 6, 7, 9},
	},
}

func init() {
	for i := 0; i < len(tests); i++ {
		grid, err := ParseGrid(tests[i].data)
		if err != nil {
			panic(err)
		}
		tests[i].grid = grid
	}
}

func TestFirstEmptyCell(t *testing.T) {
	for i, x := range tests {
		got := tests[i].grid.FirstEmptyCell()
		if got != x.firstEmptyCell {
			t.Errorf("FirstEmptyCell(grid[%d]) = %d, expected %d", i, got, x.firstEmptyCell)
		}
	}
}

func TestNeighbours(t *testing.T) {
	for i, x := range tests {
		r := x.firstEmptyCell / 9
		c := x.firstEmptyCell % 9
		got := x.grid.Neighbours(r, c)
		if len(got) != len(x.neighbours) {
			t.Errorf("len(Neighbours(grid[%d]) = %d, expected %d", i, len(got), len(x.neighbours))
		}
		for _, v := range x.neighbours {
			if !got.Contains(v) {
				t.Errorf("Neighbours(grid[%d]) did not contain %d", i, v)
			}
		}
	}
}

func TestCandidates(t *testing.T) {
	for i, x := range tests {
		r := x.firstEmptyCell / 9
		c := x.firstEmptyCell % 9
		got := x.grid.Candidates(r, c)
		if len(got) != len(x.candidates) {
			t.Errorf("len(Candidates(grid[%d]) = %d, expected %d", i, len(got), len(x.candidates))
		}
		for _, v := range x.candidates {
			found := false
			for _, w := range got {
				if v == w {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("Candidates(grid[%d]) did not contain %d", i, v)
			}
		}
	}
}

func TestSolve(t *testing.T) {
	for i, x := range tests {
		res, err := x.grid.Solve()
		if err != nil {
			t.Errorf("Could not solve grid %d", i)
		} else {
			fmt.Println(res)
		}
	}
}
