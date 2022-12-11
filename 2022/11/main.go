// This code is not performing well. It's focused on readability and simplicity.
// A lot of allocations could be avoided but the performance are still reasonable.
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"sort"
	"strings"
)

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b float64) float64 {
	for b != 0 {
		t := b
		b = math.Mod(a, b)
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(monkeys ...*Monkey) float64 {
	result := 1.

	for i := 0; i < len(monkeys); i++ {
		a := result
		b := monkeys[i].DivisibleBy
		result = a / GCD(a, b) * b
	}

	return result
}

type Monkey struct {
	// Items is a list of items with their corresponding worry level.
	Items []float64
	// Operation is the function that will be applied to the items to calculate
	// the new worry level.
	Operation func(float64) float64
	// DivisibleBy is the amount by which worry level should be divided each
	// round.
	DivisibleBy float64
	// Peers are all the monkeys that can receive items from this monkey
	Peers []*Monkey

	// Inspected is the total number of items this monkey inspected.
	Inspected int
}

func (m *Monkey) AddPeers(peers ...*Monkey) {
	m.Peers = append(m.Peers, peers...)
}

func (m *Monkey) nextItem() bool {
	return len(m.Items) > 0
}

func (m *Monkey) popItem() float64 {
	item := m.Items[0]
	m.Items = m.Items[1:]
	return item
}

func (m *Monkey) AddItem(item float64) {
	m.Items = append(m.Items, item)
}

func (m *Monkey) Round(postProcess func(float64) float64) {
	// For each item, apply the operation and test if the monkey is happy
	for m.nextItem() {
		item := m.popItem()
		new := postProcess(m.Operation(item))
		if math.Mod(new, m.DivisibleBy) == 0 {
			m.Peers[0].AddItem(new)
		} else {
			m.Peers[1].AddItem(new)
		}
		m.Inspected++
	}

}

type ByInspection []*Monkey

func (a ByInspection) Len() int      { return len(a) }
func (a ByInspection) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByInspection) Less(i, j int) bool {
	return a[i].Inspected < a[j].Inspected
}

func Input() []*Monkey {
	zero := Monkey{
		Items: []float64{91, 66},
		Operation: func(old float64) float64 {
			return old * 13
		},
		DivisibleBy: 19,
	}

	one := Monkey{
		Items: []float64{78, 97, 59},
		Operation: func(old float64) float64 {
			return old + 7
		},
		DivisibleBy: 5,
	}

	two := Monkey{
		Items: []float64{57, 59, 97, 84, 72, 83, 56, 76},
		Operation: func(old float64) float64 {
			return old + 6
		},
		DivisibleBy: 11,
	}

	three := Monkey{
		Items: []float64{81, 78, 70, 58, 84},
		Operation: func(old float64) float64 {
			return old + 5
		},
		DivisibleBy: 17,
	}

	four := Monkey{
		Items: []float64{60},
		Operation: func(old float64) float64 {
			return old + 8
		},
		DivisibleBy: 7,
	}

	five := Monkey{
		Items: []float64{57, 69, 63, 75, 62, 77, 72},
		Operation: func(old float64) float64 {
			return old * 5
		},
		DivisibleBy: 13,
	}

	six := Monkey{
		Items: []float64{73, 66, 86, 79, 98, 87},
		Operation: func(old float64) float64 {
			return old * old
		},
		DivisibleBy: 3,
	}

	seven := Monkey{
		Items: []float64{95, 89, 63, 67},
		Operation: func(old float64) float64 {
			return old + 2
		},
		DivisibleBy: 2,
	}

	zero.AddPeers(&six, &two)
	one.AddPeers(&zero, &three)
	two.AddPeers(&five, &seven)
	three.AddPeers(&six, &zero)
	four.AddPeers(&one, &three)
	five.AddPeers(&seven, &four)
	six.AddPeers(&five, &two)
	seven.AddPeers(&one, &four)

	return []*Monkey{&zero, &one, &two, &three, &four, &five, &six, &seven}
}

func Compute(monkeys []*Monkey, postProcess func(float64) float64, rounds int) int {
	for r := 0; r < rounds; r++ {
		for m := 0; m < len(monkeys); m++ {
			monkeys[m].Round(postProcess)
		}
	}

	sort.Sort(sort.Reverse(ByInspection(monkeys)))

	return monkeys[0].Inspected * monkeys[1].Inspected
}

func SolutionA(lines []string) int {
	monkeys := Input()

	postProcess := func(new float64) float64 {
		return math.Floor(new / 3)
	}

	return Compute(monkeys, postProcess, 20)
}

func SolutionB(lines []string) int {
	monkeys := Input()

	lcm := LCM(monkeys...)
	postProcess := func(new float64) float64 {
		return math.Floor(math.Mod(new, lcm))
	}

	return Compute(monkeys, postProcess, 10_000)
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
