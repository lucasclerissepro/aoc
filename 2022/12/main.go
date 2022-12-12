// This code is not performing well. It's focused on readability and simplicity.
// A lot of allocations could be avoided but the performance are still reasonable.
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strings"

	"github.com/yourbasic/graph"
)

func Abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func SolutionA(lines []string) any {
	m := [][]int{}
	x, y := 0, 0
	xe, ye := 0, 0

	// build map
	m = make([][]int, len(lines)-1)
	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			continue
		}
		m[i] = make([]int, len(lines[i]))
		for j := 0; j < len(lines[i]); j++ {
			curr := lines[i][j]
			toappend := int(curr - 'a')

			if curr == 'S' {
				toappend = 0
				x, y = j, i
			}

			if curr == 'E' {
				toappend = int('z' - 'a')
				xe, ye = j, i
			}

			m[i][j] = toappend
		}
	}

	// build graph
	g := graph.New(len(lines) * len(lines[0]))

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {

			// check left side
			if j > 0 && (m[i][j] >= m[i][j-1] || m[i][j]+1 == m[i][j-1]) {
				g.AddCost(i*len(m[i])+j, i*len(m[i])+j-1, 1)
			}

			// check right side
			if j < len(m[i])-1 && (m[i][j] >= m[i][j+1] || m[i][j]+1 == m[i][j+1]) {
				g.AddCost(i*len(m[i])+j, i*len(m[i])+j+1, 1)
			}

			// check if top side is smaller or +1
			if i > 0 && (m[i][j] >= m[i-1][j] || m[i][j]+1 == m[i-1][j]) {
				g.AddCost(i*len(m[i])+j, (i-1)*len(m[i])+j, 1)
			}

			// check bottom side
			if i < len(m)-1 && (m[i][j] >= m[i+1][j] || m[i][j]+1 == m[i+1][j]) {
				g.AddCost(i*len(m[i])+j, (i+1)*len(m[i])+j, 1)
			}
		}
	}

	_, dist := graph.ShortestPath(g, y*len(m[0])+x, ye*len(m[0])+xe)

	return dist
}

func SolutionB(lines []string) any {
	m := [][]int{}
	xe, ye := 0, 0
	sty, stx := []int{}, []int{}

	// build map
	m = make([][]int, len(lines)-1)
	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			continue
		}
		m[i] = make([]int, len(lines[i]))
		for j := 0; j < len(lines[i]); j++ {
			curr := lines[i][j]
			toappend := int(curr - 'a')

			if curr == 'S' || curr == 'a' {
				toappend = 0
				stx = append(stx, j)
				sty = append(sty, i)
			}

			if curr == 'E' {
				toappend = int('z' - 'a')
				xe, ye = j, i
			}

			m[i][j] = toappend
		}
	}

	// build graph
	g := graph.New(len(lines) * len(lines[0]))

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {

			// check left side
			if j > 0 && (m[i][j] >= m[i][j-1] || m[i][j]+1 == m[i][j-1]) {
				g.AddCost(i*len(m[i])+j, i*len(m[i])+j-1, 1)
			}

			// check right side
			if j < len(m[i])-1 && (m[i][j] >= m[i][j+1] || m[i][j]+1 == m[i][j+1]) {
				g.AddCost(i*len(m[i])+j, i*len(m[i])+j+1, 1)
			}

			// check if top side is smaller or +1
			if i > 0 && (m[i][j] >= m[i-1][j] || m[i][j]+1 == m[i-1][j]) {
				g.AddCost(i*len(m[i])+j, (i-1)*len(m[i])+j, 1)
			}

			// check bottom side
			if i < len(m)-1 && (m[i][j] >= m[i+1][j] || m[i][j]+1 == m[i+1][j]) {
				g.AddCost(i*len(m[i])+j, (i+1)*len(m[i])+j, 1)
			}
		}
	}

	ans := int64(math.MaxInt64)

	for i := 0; i < len(stx); i++ {
		x := stx[i]
		y := sty[i]
		_, dist := graph.ShortestPath(g, y*len(m[0])+x, ye*len(m[0])+xe)

		if dist > 0 && dist < ans {
			ans = dist
		}
	}

	return ans
}

func main() {
	// open file
	input, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	if len(os.Args) != 3 {
		fmt.Println("args should be '[input_path] [part]'. With [part] equal 'one' or 'two'")
		os.Exit(1)
	}

	lines := strings.Split(string(input), "\n")
	if os.Args[2] == "one" {
		fmt.Println("Solution is: ", SolutionA(lines))
		os.Exit(0)
	}

	fmt.Println("Solution is: ", SolutionB(lines))
}
