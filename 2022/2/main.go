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

func Solution(lines []string) int {
	sum := 0

	for _, line := range lines {
		choices := strings.Split(line, " ")

		if len(choices) != 2 {
			continue
		}

		switch choices[1] {
		// loss
		case "X":
			sum += 0
			switch choices[0] {
			case "A":
				sum += 3
			case "B":
				sum += 1
			case "C":
				sum += 2
			}

			// draw
		case "Y":
			sum += 3
			switch choices[0] {
			case "A":
				sum += 1
			case "B":
				sum += 2
			case "C":
				sum += 3
			}

			// win
		case "Z":
			sum += 6
			switch choices[0] {
			case "A":
				sum += 2
			case "B":
				sum += 3
			case "C":
				sum += 1
			}
		}
	}
	return sum
}

func main() {
	// open file
	input, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	lines := strings.Split(string(input), "\n")
	sum := Solution(lines)

	fmt.Println(sum)
}
