package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func ToAlphabetOrdering(a rune) int {
	if a > 'a' {
		return int(a - 96)
	}

	return int(a - 38)
}

func Intersect(candidates ...string) rune {
	if len(candidates) < 2 {
		panic("cannot intersect less than two candidates")
	}

	found := map[rune]int{}
	for idx, c := range candidates {
		for _, v := range c {

			if found[v] == idx {
				found[v]++
			}

			if found[v] == len(candidates) {
				return v
			}
		}
	}

	return '0'
}

func SolutionA(lines []string) int {
	sum := 0

	for _, line := range lines {

		if line == "" {
			continue
		}

		n := len(line)
		k := n / 2

		inter := Intersect(line[:k], line[k:])
		priority := ToAlphabetOrdering(inter)

		sum += int(priority)
	}

	return sum
}

func SolutionB(lines []string) int {
	sum := 0

	for i := 0; i < len(lines)-3; i += 3 {
		badge := Intersect(lines[i : i+3]...)
		priority := ToAlphabetOrdering(badge)
		sum += priority
	}

	return sum
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
