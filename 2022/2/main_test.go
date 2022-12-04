package main

import (
	"log"
	"os"
	"strings"
	"testing"
)

func BenchmarkSolution(b *testing.B) {
	input, err := os.ReadFile("data/input.txt")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	lines := strings.Split(string(input), "\n")

	b.Run("Solution", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Solution(lines)
		}
	})
}
