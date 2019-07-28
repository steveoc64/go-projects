package stats

// Largest returns the biggest value from the data in the array of float64 numbers.
func Largest(numbers []float64) float64 {
	// Start setting largest to the first data value.
	largest := numbers[0]

	for i := 1; i < len(numbers); i++ {
		if largest < numbers[i] {
			largest = numbers[i]
		}
	}

	return largest
}

// Smallest returns the smallest value from the data in the array of float64 numbers.
func Smallest(numbers []float64) float64 {
	smallest := numbers[0]

	for i := 1; i < len(numbers); i++ {
		if smallest > numbers[i] {
			smallest = numbers[i]
		}
	}

	return smallest
}

// Range retuns the length between the biggest and smallest values from an array of float64 numbers.
func Range(numbers []float64) float64 {
	return Largest(numbers) - Smallest(numbers)
}
