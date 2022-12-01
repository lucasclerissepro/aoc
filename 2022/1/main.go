package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"golang.org/x/exp/constraints"
)

const END_OF_ELF = ""

func FastAtoi(s string) int {
	var val int

	for _, c := range s {
		val = val*10 + int(c-'0')
	}

	return val
}

// Topk returns the top k elves with the highest total calories.
func Topk(tokens []string, k int) ([]int, error) {
	tops := make([]int, k)
	totalCalories := 0

	for i := 0; i < len(tokens); i++ {
		if tokens[i] == END_OF_ELF {

			// Find index of the lowest calories.
			index := 0
			found := false
			for ; index < k; index++ {
				if tops[index] <= totalCalories {
					found = true
					break
				}
			}

			// If index of the lowest calories is equal to k, then we set the latest
			// element to the total calories if it is greater than the lowest calories.
			if index == k {
				if found {
					tops[index-1] = totalCalories
				}
				totalCalories = 0
				continue
			}

			copy(tops[index+1:], tops[index:])
			tops[index] = totalCalories

			totalCalories = 0

			continue
		}

		cal := FastAtoi(tokens[i])
		totalCalories += cal
	}

	return tops, nil
}

// Sum returns the sum of a slice of integers.
func Sum[T constraints.Integer](nums []T) int {
	var sum int = 0

	for _, num := range nums {
		sum += int(num)
	}

	return sum
}

func main() {
	// open file
	input, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	tokens := strings.Split(string(input), "\n")
	top, err := Topk(tokens, 3)
	if err != nil {
		log.Fatalf("failed to find top 3: %s", err)
	}

	log.Printf("top k = %d", top)
	log.Printf("max calories: %d", Sum(top))
}
