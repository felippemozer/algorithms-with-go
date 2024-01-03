package bigo

import (
	"fmt"
	"strings"
)

// All of the code in this file is used when explaining Big O and viewing
// benchmarks to help illustrate a few points.

// Add is O(1) - constant time.
func Add(a, b int) int {
	return a + b
}

// SumToMax is O(N) where N is the value of mazx.
func SumToMax(max int) int {
	sum := 0
	for i := 1; i <= max; i++ {
		sum += i
	}
	return sum
}

// SumToMaxV2 illustrates that some functions can do the same thing but much
// faster. V2 is O(1) - constante time.
func SumToMaxV2(max int) int {
	return (max * (max + 1)) / 2
}

// SumVals is O(N) where N is the size of the vals slice. ie N == len(vals).
func SumVals(vals []int) int {
	var sum int
	for _, val := range vals {
		sum += val
	}
	return sum
}

// Find has two inputs, but only one affects Big O. This is O(N) where N is the
// size of the list.
func Find(list []int, x int) int {
	for i, val := range list {
		if val == x {
			return i
		}
	}
	return -1
}

// Grid is an example of function that's Big O is determined by two inputs - x
// and y. Grid is O(XY) where X and Y are the values of x and y, respectively.
func Grid(x, y int) string {
	var sb strings.Builder
	chars := []byte{'x', 'o'}
	charIdx := 0
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			sb.WriteByte(chars[charIdx])
			charIdx = 1 - charIdx // alternantes between 0 and 1
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

func Cube(n int) string {
	var sb strings.Builder
	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			for z := 0; z < n; z++ {
				sb.WriteString(fmt.Sprintf("(%d %d %d)\n", x, y, z))
			}
		}
	}
	return sb.String()
}
