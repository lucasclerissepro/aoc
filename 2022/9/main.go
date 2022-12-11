// This code is not performing well. It's focused on readability and simplicity.
// A lot of allocations could be avoided but the performance are still reasonable.
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"golang.design/x/clipboard"
)

func ShouldMove(x, y, dx, dy int) bool {
	return (                  // check side movement
	(x == dx && y+2 == dy) || // up
		(x == dx && y-2 == dy) || // down
		(x+2 == dx && y == dy) || // right
		(x-2 == dx && y == dy) || // left

		(x+2 == dx && y+1 == dy) || // up right
		(x+2 == dx && y-1 == dy) || // down right
		(x-2 == dx && y+1 == dy) || // up left
		(x-2 == dx && y-1 == dy) || // down left

		(x+1 == dx && y+2 == dy) || // up right
		(x+1 == dx && y-2 == dy) || // down right
		(x-1 == dx && y+2 == dy) || // up left
		(x-1 == dx && y-2 == dy))
}

func SolutionA(lines []string) string {
	visited := map[string]bool{}

	tx := 0
	ty := 0
	hx := 0
	hy := 0

	visited["0,0"] = true

	for i := 0; i < len(lines); i++ {

		if lines[i] == "" {
			continue
		}

		a := strings.Fields(lines[i])
		q, _ := strconv.Atoi(a[1])

		for j := 0; j < q; j++ {
			lastx := hx
			lasty := hy

			switch a[0] {
			case "U":
				hy++
			case "D":
				hy--
			case "R":
				hx++
			case "L":
				hx--
			}

			if ShouldMove(tx, ty, hx, hy) {
				visited[fmt.Sprintf("%d,%d", lastx, lasty)] = true
				tx = lastx
				ty = lasty
			}
		}

	}

	return strconv.Itoa(len(visited))
}

func SolutionB(lines []string) string {
	visited := map[string]bool{}

	n := 10

	xs := make([]int, n)
	ys := make([]int, n)

	visited["0,0"] = true

	for i := 0; i < len(lines); i++ {

		if lines[i] == "" {
			continue
		}

		a := strings.Fields(lines[i])
		q, _ := strconv.Atoi(a[1])

		for j := 0; j < q; j++ {

			switch a[0] {
			case "U":
				ys[n-1]++
			case "D":
				ys[n-1]--
			case "R":
				xs[n-1]++
			case "L":
				xs[n-1]--
			}

      for k := 1; k < n-1; k++ {
        fmt.Printf("Checking if %d,%d should move to %d,%d\n", xs[k-1], ys[k-1], xs[k], ys[k])
				if ShouldMove(xs[k-1], ys[k-1], xs[k], ys[k]) {
					xs[k-1] = lastxs[k]
					ys[k-1] = lastys[k]
				}
			}
		}

	}

	return strconv.Itoa(len(visited))
}

func main() {
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}

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
	var ans []byte

	if os.Args[2] == "one" {
		ans = []byte(SolutionA(lines))
		fmt.Println(string(ans))
	} else {
		ans = []byte(SolutionB(lines))
		fmt.Println(string(ans))
	}

	clipboard.Write(clipboard.FmtText, ans)
}
