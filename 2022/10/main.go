// This code is not performing well. It's focused on readability and simplicity.
// A lot of allocations could be avoided but the performance are still reasonable.
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

// Sum returns the sum of a slice of integers.
func Sum[T constraints.Integer](nums []T) int {
	var sum int = 0

	for _, num := range nums {
		sum += int(num)
	}

	return sum
}

// Instruction represents a CPU instruction that can be executed when matched.
// The Regex is used to match the instruction.
// The Cost is the number of cycles the instruction takes.
// The Op is the operation to execute when the instruction is matched.
type Instruction struct {
	Regex *regexp.Regexp
	Op    func(int, []int) int
	Cost  int
}

// Currently available instructions
var instructions = []Instruction{
	{
		Regex: regexp.MustCompile(`^addx ([-+]?[0-9]*\.?[0-9]+)$`),
		Cost:  2,
		Op: func(x int, captures []int) int {
			return x + captures[0]
		},
	},
	{
		Regex: regexp.MustCompile(`^noop`),
		Cost:  1,
		Op:    nil,
	},
}

// ForEachInstruction runs the instructions and calls the callback for each instruction.
// The callback is called multiple time depending of the cost of each
// instruction.
func ForEachInstruction(lines []string, f func(*Instruction, int, int) error) error {
	var err error

	cycles := 1
	x := 1

	for i := 0; i < len(lines); i++ {
		for tx := 0; tx < len(instructions); tx++ {
			instruction := instructions[tx]
			captures := instruction.Regex.FindStringSubmatch(lines[i])

			if captures == nil {
				continue
			}

			inputs := make([]int, len(captures)-1)

			// Convert the captures to integers
			if len(captures) > 1 {
				for j := 1; j < len(captures); j++ {
					inputs[j-1], err = strconv.Atoi(captures[j])
					if err != nil {
						return err
					}
				}
			}

			for cost := 1; cost <= instruction.Cost; cost++ {

				f(&instruction, cycles, x)

				if cost == instruction.Cost && instruction.Op != nil {
					x = instruction.Op(x, inputs)
				}

				cycles++
			}
		}
	}

	return nil
}

func SolutionA(lines []string) int {
	ans := []int{}
	nextMilestone := 0
	milestones := []int{
		20,
		60,
		100,
		140,
		180,
		220,
	}

	err := ForEachInstruction(lines, func(instruction *Instruction, cycles int, x int) error {
		if nextMilestone < len(milestones) && cycles == milestones[nextMilestone] {
			ans = append(ans, cycles*x)
			nextMilestone++
		}
		return nil
	})

	if err != nil {
		log.Fatalf("failed to run instructions: %s", err)
	}

	return Sum(ans)
}

func SolutionB(lines []string) string {
	crt := make([]rune, 40*6)

	err := ForEachInstruction(lines, func(instruction *Instruction, cycles int, x int) error {
		crtCursor := cycles - 1
		if (crtCursor%40) >= x-1 && (crtCursor%40) <= x+1 {
			crt[crtCursor] = '#'
		} else {
			crt[crtCursor] = '.'
		}
		return nil
	})

	if err != nil {
		log.Fatalf("failed to run instructions: %s", err)
	}

	rendered := strings.Builder{}
	for y := 0; y < 6; y++ {
		for x := 0; x < 40; x++ {
			rendered.WriteRune(crt[(y*40)+x])
		}
		rendered.WriteRune('\n')
	}

	return rendered.String()
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

	fmt.Println("Solution is: \n", SolutionB(lines))
}
