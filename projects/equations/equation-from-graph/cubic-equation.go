package main

import (
	"fmt"
	"math"
)

func bcdValues(Xvalue1, Xvalue2, Xvalue3, A float64) (B, C, D float64) {
	B = A * (-Xvalue3 - Xvalue2 - Xvalue1)
	C = A * ((Xvalue2 * Xvalue3) + (Xvalue1 * Xvalue3) + (Xvalue1 * Xvalue2))
	D = A * -(Xvalue1 * Xvalue2 * Xvalue3)

	return B, C, D
}

func gradient(YaxisX, YaxisY, Xvalue1, Xvalue2, Xvalue3 float64) float64 {
	// Function that calculates the gradient from one point and three differnet x-values
	// Look up how to get the k-value from a cubic equation to understand it
	return YaxisY / ((YaxisX - Xvalue1) * (YaxisX - Xvalue2) * (YaxisX - Xvalue3))
}

/*
func equationRoots(NullPoints int) (NullPoint1, NullPoint2, NullPoint3 float64) {
	switch NullPoints {
	case 3:
		NullPoint1, NullPoint2, NullPoint3 = threeRoots()
	case 2:
		NullPoint1, NullPoint2 = twoRoots()
		// The point is called a rebound point and two root points will have the same value
		NullPoint3 = NullPoint2
	case 1:
		NullPoint1 = oneRoot()
		// The point is called a rebound point and all three root points will have the same value
		NullPoint2, NullPoint3 = NullPoint1, NullPoint1
	case 0:
		log.Fatalln("The function has no real roots and can't be parsed by this program!")
	default:
		log.Fatalln("A cubic function can only have 3, 2, 1 or 0 roots!")
	}

	return NullPoint1, NullPoint2, NullPoint3
}
*/
func threeRoots() (NullPoint1, NullPoint2, NullPoint3 float64) {
	fmt.Println("\nEnter x-values for the three root points:")
	fmt.Print("First x-value: ")
	fmt.Scan(&NullPoint1)

	fmt.Print("Second x-value: ")
	fmt.Scan(&NullPoint2)

	fmt.Print("Third x-value: ")
	fmt.Scan(&NullPoint3)

	return NullPoint1, NullPoint2, NullPoint3
}

/*
func twoRoots() (NullPoint1, NullPoint2 float64) {
	fmt.Println("\nEnter x-values for the two root points:")
	fmt.Print("First x-value: ")
	fmt.Scan(&NullPoint1)

	fmt.Print("Second x-value: ")
	fmt.Scan(&NullPoint2)

	return NullPoint1, NullPoint2
}

// Get the root point for one root
func oneRoot() (NullPoint float64) {
	fmt.Println("\nEnter x-value for the root point:")
	fmt.Print("X-value: ")
	fmt.Scan(&NullPoint)

	return NullPoint
}
*/
func main() {
	var (
		Xvalue, Yvalue                     float64
		NullPoint1, NullPoint2, NullPoint3 float64
		//NullPoints                         int
	)

	fmt.Println("Enter a couple values from graph to get the cubic equation!")

	// Enter the amount of roots in the graph
	// fmt.Print("\nAmount of roots on the graph (points where y = 0): ")
	// fmt.Scanln(&NullPoints)
	NullPoint1, NullPoint2, NullPoint3 = threeRoots()

	// A random value from the graph to calculate gradient of line
	fmt.Print("\nEnter a given point on the graph (x, y): ")
	fmt.Scanf("(%v, %v)", &Xvalue, &Yvalue)

	// Caluclate gradient and the specific varaibles that define the graph
	A := gradient(Xvalue, Yvalue, NullPoint1, NullPoint2, NullPoint3)
	B, C, D := bcdValues(NullPoint1, NullPoint2, NullPoint3, A)

	if math.Round(A) == 1 {
		fmt.Printf("y = x³ + %.0fx² + %.0fx + %.0f\n", B, C, D)
	} else {
		fmt.Printf("y = %.0fx³ + %.0fx² + %.0fx + %.0f\n", A, B, C, D)
	}
}
