// This code is not performing well. It's focused on readability and simplicity.
// A lot of allocations could be avoided but the performance are still reasonable.
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type ReduceFunc func(int, int, int) int

func ReduceFilesystem(cmds []string, init int, f ReduceFunc) int {
	path := []string{}
	size := map[string]int{}

	for i := 0; i < len(cmds); i++ {
		if cmds[i] == "" {
			continue
		}

		cmd := strings.Split(cmds[i], " ")

		if cmd[1] == "cd" {
			if cmd[2] == ".." {
				path = path[:len(path)-1]
				continue
			}
			path = append(path, cmd[2])
		} else if cmd[1] != "ls" && cmd[0] != "dir" {
			for j := 0; j < len(path)+1; j++ {
				s, err := strconv.Atoi(cmd[0])
        if err != nil {
          panic("failed to parse file size")
        }
				size["/"+strings.Join(path[:j], "+")] += s
			}
		}
	}

	ans := init
	for _, v := range size {
		ans = f(size["/"], ans, v)
	}

	return ans
}

func SolutionA(cmds []string) int {
	return ReduceFilesystem(cmds, 0, func(_, acc, val int) int {
		if val <= 100000 {
			return acc + val
		}
		return acc
	})
}

func SolutionB(cmds []string) int {
	return ReduceFilesystem(cmds, math.MaxInt, func(size, acc, val int) int {
		tofree := 30000000 - (70000000 - size)
		if val >= tofree && val < acc {
			return val
		}
		return acc
	})
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
