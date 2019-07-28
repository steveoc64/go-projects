package trianglebench

import "math"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Run runs the triangle benchmark!
func Run(calc int) int {
	c := 0
	var a, b int
	for a = 1; a < calc+1; a++ {

		for b = a; b < min(a*a+1, calc+1); b++ {
			if math.Sqrt(float64(a*a+b*b-a*b)) == float64(int64(math.Sqrt(float64(a*a+b*b-a*b)))) {
				c++
			}
		}
	}

	return c
}
