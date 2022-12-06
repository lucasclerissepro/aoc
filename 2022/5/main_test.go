package main

import (
	"io/ioutil"
	"log"
	"strings"
	"testing"
)

func BenchmarkSolutionA(b *testing.B) {
	// open file
	input, err := ioutil.ReadFile("./data/input.txt")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	lines := strings.Split(string(input), "\n")

	b.Run("SolutionA", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
		  SolutionA(lines)
    }
	})
}

func BenchmarkSolutionB(b *testing.B) {
	// open file
	input, err := ioutil.ReadFile("./data/input.txt")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	lines := strings.Split(string(input), "\n")

	b.Run("SolutionB", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
		  SolutionB(lines)
    }
	})
}


func BenchmarkSolutionAF(b *testing.B) {
	// open file
	input, err := ioutil.ReadFile("./data/input.txt")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	lines := strings.Split(string(input), "\n")

	b.Run("SolutionAFabien", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
		  ex1(lines)
    }
	})
}

func BenchmarkSolutionBF(b *testing.B) {
	// open file
	input, err := ioutil.ReadFile("./data/input.txt")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	lines := strings.Split(string(input), "\n")

	b.Run("SolutionBFabien", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
		  ex2(lines)
    }
	})
}
