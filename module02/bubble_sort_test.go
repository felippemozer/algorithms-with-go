package module02

import (
	"module02/utils"
	"testing"
)

func TestBubbleSortInt(t *testing.T) {
	utils.TestInt(t, BubbleSortInt)
}

func BenchmarkBubbleSortInt(b *testing.B) {
	utils.BenchmarkInt(b, BubbleSortInt)
}

// func TestBubbleSortString(t *testing.T) {
// 	utils.TestString(t, BubbleSortString)
// }

// func TestBubbleSortInterface(t *testing.T) {
// 	utils.TestInterface(t, BubbleSort)
// }
