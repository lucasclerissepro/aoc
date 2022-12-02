package main

import (
	"io/ioutil"
	"log"
	"os"

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

func main() {
	// open file
	_, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
}
