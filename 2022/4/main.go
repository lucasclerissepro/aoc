package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
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

func ParseElves(s string) []string {
	return strings.Split(s, ",")
}

func ParsePair(s string) (int, int) {
	sectors := strings.Split(s, "-")
	begin, _ := strconv.Atoi(sectors[0])
	end, _ := strconv.Atoi(sectors[1])
	return begin, end
}

func SolutionA(lines []string) int {
	ans := 0

	for _, line := range lines {

		if line == "" {
			continue
		}

		elves := ParseElves(line)
		ls, le := ParsePair(elves[0])
		rs, re := ParsePair(elves[1])

		// left covers right
		if ls <= rs && le >= re {
			ans += 1
		} else if rs <= ls && re >= le {
			ans += 1
		}
	}

	return ans
}

func SolutionB(lines []string) int {
	ans := 0

	for _, line := range lines {

		if line == "" {
			continue
		}

		elves := ParseElves(line)
		ls, le := ParsePair(elves[0])
		rs, re := ParsePair(elves[1])

		if ls <= rs && le >= rs || le >= re && ls <= re {
			ans += 1
		} else if rs <= ls && re >= ls || re >= le && rs <= le {
			ans += 1
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
