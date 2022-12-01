package main

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// open file
	input, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	tokens := strings.Split(string(input), "\n")
	maxCalories := -1
	currentCalories := 0

	for _, token := range tokens {
		if token == "" {
			if currentCalories > maxCalories {
				maxCalories = currentCalories
			}
			currentCalories = 0

			continue
		}

		cal, err := strconv.Atoi(token)
		if err != nil {
			log.Fatalf("failed to parse calories: %s", err)
		}

		currentCalories += cal
	}

	log.Printf("max calories: %d", maxCalories)
}
