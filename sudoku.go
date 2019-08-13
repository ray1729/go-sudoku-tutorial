package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: sudoku FILENAME")
		os.Exit(2)
	}
	filename := os.Args[1]
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading puzzle file: %v\n", err)
		os.Exit(3)
	}
	grid, err := ParseGrid(data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing grid: %v\n", err)
		os.Exit(3)
	}
	solution, err := grid.Solve()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(3)
	}
	fmt.Println(solution)
}
