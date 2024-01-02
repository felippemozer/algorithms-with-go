package module01

func Fibonacci(n int) int {
	base := []int{0, 1}
	if n <= 1 {
		return base[n]
	}
	result := 1
	for i := 1; i < n; i++ {
		result = result + base[len(base)-2]
		base = append(base, result)
	}
	return result
}
