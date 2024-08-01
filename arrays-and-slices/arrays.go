package arraysandslices


// Arrays are declared with a fixed length WHILE slices are declared with any length 
func SumForArrays(numbers [5]int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

