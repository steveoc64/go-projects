package stats

import "testing"

func TestTableLargest(t *testing.T) {
	var tests = []struct {
		input    []float64
		expected float64
	}{
		// first, second and third are imported from numbers_test.go
		{first, 8},
		{second, 102},
		{third, 623},
	}

	for _, test := range tests {
		if output := Largest(test.input); output != test.expected {
			t.Errorf("Test Failed: %v inputed, %v expected, recieved %v", test.input, test.expected, output)
		}
	}
}

func TestTableSmallest(t *testing.T) {
	var tests = []struct {
		input    []float64
		expected float64
	}{
		// first, second and third are imported from numbers_test.go
		{first, 1},
		{second, 1},
		{third, 7},
	}

	for _, test := range tests {
		if output := Smallest(test.input); output != test.expected {
			t.Errorf("Test Failed: %v inputed, %v expected, recieved %v", test.input, test.expected, output)
		}
	}
}

func TestTableRange(t *testing.T) {
	var tests = []struct {
		input    []float64
		expected float64
	}{
		// first, second and third are imported from numbers_test.go
		{first, 7},
		{second, 101},
		{third, 616},
	}

	for _, test := range tests {
		if output := Range(test.input); output != test.expected {
			t.Errorf("Test Failed: %v inputed, %v expected, recieved %v", test.input, test.expected, output)
		}
	}
}
