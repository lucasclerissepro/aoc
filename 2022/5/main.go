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
)

type Stack []string

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push new value at the begining of the stack
func (s *Stack) PushFront(str string) {
	*s = append([]string{str}, *s...)
}

func (s *Stack) PushTail(str string) {
	*s = append(*s, str)
}

// Remove and return first element of stack. Return false if stack is empty.
func (s *Stack) Pop() string {
	if s.IsEmpty() {
		return ""
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element
	}
}

// Pop N element from the stack
func (s *Stack) PopN(n int) []string {
	if s.IsEmpty() {
		return []string{}
	} else {
		index := len(*s) - n    // Get the index of the top most element.
		element := (*s)[index:] // Index into the slice and obtain the element.
		*s = (*s)[:index]       // Remove it from the stack by slicing it off.
		return element
	}
}

// Push all given elements to the stack
func (s *Stack) PushAll(elements []string) {
	*s = append(*s, elements...)
}

// Peek returns the first element of the stack without removing it.
func (s *Stack) Peek() string {
	return (*s)[len(*s)-1]
}


// SplitSpecs split the input into header and moves
func SplitSpecs(lines []string) ([]string, []string) {
	moveSectionIdx := 0
	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			moveSectionIdx = i
			break
		}
	}

	return lines[:moveSectionIdx], lines[moveSectionIdx+1:]
}

// ParseStacks parse the different stacks from the header
func ParseStacks(lines []string) []Stack {
	stacks := []Stack{}

	for i := 0; i < len(lines)-1; i++ {
		for j := 1; j < len(lines[i]); j += 4 {

			if i == 0 {
				stacks = append(stacks, []string{})
			}

			if lines[i][j] == ' ' {
				continue
			}

			stacks[j/4].PushFront(string(lines[i][j]))
		}
	}

	return stacks
}

// A fonction that runs some actions given the move to do
type MoveFunc func(int, int, int)

// ForeachMoves run the given function for each move
func ForeachMoves(moves []string, f MoveFunc) {
	for i := 0; i < len(moves); i++ {
		if moves[i] == "" {
			continue
		}

		// get all digit from the string without using regex
		m := make([]int, 3)
		cursor := 0
		digit := ""
		for j := 0; j < len(moves[i]); j++ {
			if moves[i][j] >= '0' && moves[i][j] <= '9' {
				for k := j; k < len(moves[i]); k++ {
					if moves[i][k] < '0' || moves[i][k] > '9' {
						break
					}
					digit += string(moves[i][k])
					j = k
				}
				m[cursor], _ = strconv.Atoi(digit)
				digit = ""
				cursor++
			}
		}

		f(m[0], m[1], m[2])
	}
}

func SolutionA(lines []string) string {
	header, moves := SplitSpecs(lines)
	stacks := ParseStacks(header)

	ForeachMoves(moves, func(quantity, from, to int) {
		for i := 0; i < quantity; i++ {
			stacks[to-1].PushTail(stacks[from-1].Pop())
		}
	})

	ans := ""
	for i := 0; i < len(stacks); i++ {
		ans += stacks[i].Peek()
	}

	return ans
}

func SolutionB(lines []string) string {
	header, moves := SplitSpecs(lines)
	stacks := ParseStacks(header)

	ForeachMoves(moves, func(quantity, from, to int) {
		stacks[to-1].PushAll(stacks[from-1].PopN(quantity))
	})

	ans := ""
	for i := 0; i < len(stacks); i++ {
		ans += stacks[i].Peek()
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
