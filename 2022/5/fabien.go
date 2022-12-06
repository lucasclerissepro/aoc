package main

import (
	"strings"
)

func reverse(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {

		// swap the letters of the string,
		// like first with last and so on.
		s[i], s[j] = s[j], s[i]
	}

	// return the reversed string.
	return s
}
func GetBoxes(boxes [10][]byte, ss []string) [10][]byte {
	for idx, s := range ss {
		if idx == 8 || len(s) == 0 {
			break
		}
		if s[0] == '[' || s[0] == ' ' {
			var idx2 = 1
			for i := 1; i < len(s); i += 4 {
				if s[i] != ' ' {
					boxes[idx2] = append(boxes[idx2], s[i])
				}
				idx2++
			}
		}
	}
	for idx, val := range boxes {
		boxes[idx] = reverse(val)
	}
	return boxes
}

func moveBoxes(boxes [10][]byte, ss []string) [10][]byte {
	for i := 10; i < len(ss); i++ {
		split := strings.Split(ss[i], " ")
		n := FastAtoi(split[1])
		n1 := FastAtoi(split[3])
		n2 := FastAtoi(split[5])

		for i2 := 0; i2 < n; i2++ {
			boxes[n2] = append(boxes[n2], boxes[n1][len(boxes[n1])-1:]...)
			boxes[n1] = boxes[n1][:len(boxes[n1])-1]
		}
	}
	return boxes
}

func ex1(ss []string) string {
	var boxes [10][]byte
	boxes = GetBoxes(boxes, ss)
	boxes = moveBoxes(boxes, ss)
	var res string
	for idx := 1; idx < 10; idx++ {
		res += string(boxes[idx][len(boxes[idx])-1])
	}
	return res
}

func moveBoxes2(boxes [10][]byte, ss []string) [10][]byte {
	var split []string
	for i := 10; i < len(ss); i++ {
		split = strings.Split(ss[i], " ")
		n := FastAtoi(split[1])
		n1 := FastAtoi(split[3])
		n2 := FastAtoi(split[5])
		boxes[n2] = append(boxes[n2], boxes[n1][len(boxes[n1])-n:]...)
		boxes[n1] = boxes[n1][:len(boxes[n1])-n]
	}
	return boxes
}

func ex2(ss []string) string {
	var boxes [10][]byte
	boxes = GetBoxes(boxes, ss)
	boxes = moveBoxes2(boxes, ss)
	var res string
	for idx := 1; idx < 10; idx++ {
		res += string(boxes[idx][len(boxes[idx])-1])
	}
	return res
}
