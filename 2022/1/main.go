package main

import (
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

// Topk returns the top k elves with the highest total calories.
func Topk(tokens []string, k int) ([]int, error) {
	tops := make([]int, k)
	totalCalories := 0

	for _, token := range tokens {
		if token == "" {
			index := sort.Search(k, func(index int) bool { return tops[index] <= totalCalories })

			if index >= k {
				// sort.Search returns the index of last element if it could not find the
				// element using predicate. This condition just make sure we only replace
				// last element only if the new topCalories is bigger.
				if totalCalories > tops[index-1] {
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

		cal, err := strconv.Atoi(token)
		if err != nil {
			log.Fatalf("failed to parse calories: %s", err)
		}

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
