package stats

import "math"

// Variance returns the varaiance (average of the varations from the average) of an array of numbers as a float64 number. The second input is a bool and that determens if the data is from a sample. Simples are divided with n - 1 instead of n.
func Variance(numbers []float64, sample bool) float64 {
	average := Mean(numbers)
	length := len(numbers)

	var sum float64
	for i := 0; i < length; i++ {
		sum += math.Pow(numbers[i]-average, 2)
	}

	var output float64
	if sample {
		output = sum / (float64(length) - 1)
	} else {
		output = sum / float64(length)
	}

	return output
}

// StdDeviation returns the standard deviation of an array of numbers as float64.
func StdDeviation(numbers []float64, sample bool) float64 {
	return math.Sqrt(Variance(numbers, sample))
}
