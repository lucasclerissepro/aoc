// This code is not performing well. It's focused on readability and simplicity.
// A lot of allocations could be avoided but the performance are still reasonable.
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Unique returns true if the string contains only unique characters
func Unique(slice string) bool {
	m := make(map[rune]bool)
	for _, v := range slice {
		if m[v] {
			return false
		}
		m[v] = true
	}
	return true
}

func FindSequenceN(buffer string, n int) int {
  for i := 0; i < len(buffer)-n; i++ {
    b := buffer[i : i+n]
    if Unique(b) {
      return i + n
    }
  }

  return -1
}

func SolutionA(buffer string) int {
	return FindSequenceN(buffer, 4)
}

func SolutionB(buffer string) int {
  return FindSequenceN(buffer, 14)
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

	if os.Args[2] == "one" {
		fmt.Println("Solution is: ", SolutionA(string(input)))
		os.Exit(0)
	}

	fmt.Println("Solution is: ", SolutionB(string(input)))
}
