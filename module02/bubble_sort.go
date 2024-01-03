package module02

// BubbleSortInt will sort a list of integers using the bubble sort algorithm.
//
// Big O: O(N^2), where N == len(list)
func BubbleSortInt(list []int) {
	n := len(list)

	for n > 0 {
		for i := 0; i < n-1; i++ {
			if list[i] > list[i+1] {
				aux := list[i]
				list[i] = list[i+1]
				list[i+1] = aux
			}
		}
		n = n - 1
	}
}
