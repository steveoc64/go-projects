package stats

// Finds if the number exists in the array and returns a bool along with what index that number exists in.
func findNumIndex(value float64, data []float64) (exists bool, index int) {
	for i := 0; i < len(data); i++ {
		if value == data[i] {
			exists, index = true, i
			break
		} else {
			continue
		}
	}

	return exists, index
}

// LargestIndex returns the index for the largest number in an array of ints.
func largestPosition(numbers []int) (index int) {

	// Start setting largest to the first data value.
	largest := numbers[0]
	for i := 1; i < len(numbers); i++ {
		if largest < numbers[i] {
			largest = numbers[i]
			index = i
		}
	}

	return index
}

// Mode returns the number that exists the most times in the array of float64 numbers. The first value in the array will be outputed if every number exists once. Currently only work when there is's only one value that appears in most places.
func Mode(numbers []float64) float64 {
	data, count := []float64{}, []int{}
	var exists bool
	var index int

	for i := 0; i < len(numbers); i++ {
		exists, index = findNumIndex(numbers[i], data)
		if i == 0 || !exists {
			data = append(data, numbers[i])
			count = append(count, 1)
		} else {
			count[index]++
		}
	}

	return data[largestPosition(count)]
}
