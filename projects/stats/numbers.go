package stats

// Mean takes an array of float64 and returns the average as a float64 number.
func Mean(numbers []float64) float64 {
	var sum float64
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}

	return sum / float64(len(numbers))
}

// Median returns the number in the middle of the array of float64 numbers.
func Median(numbers []float64) (output float64) {
	length := len(numbers)
	var index float64
	if length%2 == 0 {
		index = float64(length) * 0.5
		output = (numbers[int(index)] + numbers[int(index-1)]) * 0.5
	} else {
		index = 0.5*float64(length) + 0.5
		output = numbers[int(index-1)]
	}

	return output
}
