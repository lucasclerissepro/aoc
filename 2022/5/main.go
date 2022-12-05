package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"golang.org/x/exp/constraints"
)

// Sum returns the sum of a slice of integers.
func Sum[T constraints.Integer](nums []T) int {
	var sum int = 0

	for i := 0; i < len(nums); i++ {
		sum += int(nums[i])
	}

	return sum
}

// Sum returns the sum of a slice of integers.
func Cumsum[T constraints.Integer](nums []T) []int {
	cumsum := []int{}
	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += int(nums[i])
		cumsum = append(cumsum, sum)
	}

	return cumsum
}

func SolutionA(lines []string) int {
	return 1
}

func SolutionB(lines []string) int {
	return 2
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
