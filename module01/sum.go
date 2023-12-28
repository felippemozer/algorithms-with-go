package module01

// Sum will sum up all of the numbers passed
// in and return the result
func Sum(list []int) int {
	var sum int
	for _, v := range list {
		sum = sum + v
	}

	return sum
}
