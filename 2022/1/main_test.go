package main

import (
	"log"
	"os"
	"strings"
	"testing"
)

func benchmarkSolution(k int, b *testing.B) {
	input, err := os.ReadFile("data/part_two.txt")
	if err != nil {
		b.Fatalf("failed to read input: %s", err)
	}
	tokens := strings.Split(string(input), "\n")

	b.Run("solution", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			tops, err := Topk(tokens, k)
			if len(tops) != k || err != nil {
				log.Fatalf("failed to find top 3: %s", err)
			}
		}
	})
}

func BenchmarkSolutionTopk3(b *testing.B)  { benchmarkSolution(3, b) }
func BenchmarkSolutionTopk5(b *testing.B)  { benchmarkSolution(5, b) }
func BenchmarkSolutionTopk10(b *testing.B) { benchmarkSolution(10, b) }
func BenchmarkSolutionTopk25(b *testing.B) { benchmarkSolution(25, b) }
func BenchmarkSolutionTopk50(b *testing.B) { benchmarkSolution(50, b) }
