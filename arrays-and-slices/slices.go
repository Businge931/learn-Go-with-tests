package arraysandslices

// Arrays are declared with a fixed length WHILE slices are declared with any length
func SumForSlices(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
