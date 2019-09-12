package main

import (
	"fmt"
	"log"
	"math"
)

// Gotta tell it which point rebounds twice. Thanks Anna! The rebind point is the one that doesn't go through!!!

// Function to take text to print and then scan input rigth after. Only for ints.
func printScanInt(output string) (input int) {
	fmt.Print(output)
	fmt.Scanln(&input)

	return input
}

// Function to take text to print and then scan input rigth after. Only for float64s.
func printScanFloat(output string) (input float64) {
	fmt.Print(output)
	fmt.Scanln(&input)

	return input
}

// Function to calculate all the values of B, C and D in the equation. Look up maths if you don't understand.
func bcdValues(Xvalue1, Xvalue2, Xvalue3, A float64) (B, C, D float64) {
	B = A * (-Xvalue3 - Xvalue2 - Xvalue1)
	C = A * ((Xvalue2 * Xvalue3) + (Xvalue1 * Xvalue3) + (Xvalue1 * Xvalue2))
	D = A * -(Xvalue1 * Xvalue2 * Xvalue3)

	return B, C, D
}

// Gradient calculates the A part of the equation which is the gradient of the line. Look up maths if you don't understand.
func gradient(YaxisX, YaxisY, Xvalue1, Xvalue2, Xvalue3 float64) float64 {
	// Function that calculates the gradient from one point and three differnet x-values
	// Look up how to get the k-value from a cubic equation to understand it
	return YaxisY / ((YaxisX - Xvalue1) * (YaxisX - Xvalue2) * (YaxisX - Xvalue3))
}

// Common inputs for the three different amount of root points.
func equationRoots(NullPoints int) (NullPoint1, NullPoint2, NullPoint3 float64) {

	// Handle common inputs for differrent amount of root points.
	if NullPoints == 3 || NullPoints == 2 {
		NullPoint1 = printScanFloat("First x-value: ")
		NullPoint2 = printScanFloat("Second x-value: ")
	} else if NullPoints == 1 {
		NullPoint1 = printScanFloat("X-value: ")
		NullPoint2, NullPoint3 = NullPoint1, NullPoint1
	} else {
		log.Fatalln("A cubic function can only have 3, 2 or 1 roots!")
	}

	// Ask for third point only if we have three points.
	if NullPoints == 3 {
		NullPoint3 = printScanFloat("Third x-value: ")
	}

	return NullPoint1, NullPoint2, NullPoint3
}

func main() {
	var (
		Xvalue, Yvalue                     float64
		NullPoint1, NullPoint2, NullPoint3 float64
		A, B, C, D                         float64
		NullPoints                         int
	)

	fmt.Println("Enter a couple values from graph to get the cubic equation!")

	// Enter the amount of roots in the graph
	NullPoints = printScanInt("\nAmount of roots on the graph (points where y = 0): ")

	// Calculate null points according to amound of nullpoints.
	NullPoint1, NullPoint2, NullPoint3 = equationRoots(NullPoints)

	// A random value from the graph to calculate gradient of line
	fmt.Print("\nEnter a given point on the graph (x, y): ")
	fmt.Scanf("(%v, %v)", &Xvalue, &Yvalue)

	// Caluclate gradient and the specific varaibles that define the graph
	if NullPoints == 1 || NullPoints == 3 {
		A = gradient(Xvalue, Yvalue, NullPoint1, NullPoint2, NullPoint3)
		B, C, D = bcdValues(NullPoint1, NullPoint2, NullPoint3, A)
	} else {
		// We get the right result for the function 3(x-2)(x-2)(x+1) using this block. -1, 2 (0, 12)
		A = gradient(Xvalue, Yvalue, NullPoint1, NullPoint2, NullPoint2)
		B, C, D = bcdValues(NullPoint1, NullPoint2, NullPoint2, A)

		HL := A * (Xvalue - NullPoint2) * (Xvalue - NullPoint2) * (Xvalue - NullPoint1)

		// Needs to do this for 3(x-2)(x+1)(x+1) to get the rigth result, but I can't get it to work. -1, 2 (1, -12)
		if Yvalue != HL {

			A = gradient(Xvalue, Yvalue, NullPoint2, NullPoint1, NullPoint1)
			B, C, D = bcdValues(NullPoint2, NullPoint1, NullPoint1, A)
		}

	}

	// Make sure that we prettyPrint everything to avoid printing "+ -5x" for example.
	prettyPrint(A, B, C, D)
}

// This is an utter abommination. Don't judge me please...
func prettyPrint(A, B, C, D float64) {
	if math.Round(A) == 1 {
		if B < 0 && C >= 0 && D >= 0 {
			fmt.Printf("y = x³ - %.0fx² + %.0fx + %.0f\n", -B, C, D)
		} else if B > 0 && C <= 0 && D >= 0 {
			fmt.Printf("y = x³ + %.0fx² - %.0fx + %.0f\n", B, -C, D)
		} else if B > 0 && C >= 0 && D <= 0 {
			fmt.Printf("y = x³ + %.0fx² + %.0fx - %.0f\n", B, C, -D)
		} else if B < 0 && C <= 0 && D >= 0 {
			fmt.Printf("y = x³ - %.0fx² - %.0fx + %.0f\n", -B, -C, D)
		} else if B > 0 && C <= 0 && D <= 0 {
			fmt.Printf("y = x³ + %.0fx² - %.0fx - %.0f\n", B, -C, -D)
		} else if B < 0 && C >= 0 && D <= 0 {
			fmt.Printf("y = x³ - %.0fx² + %.0fx - %.0f\n", -B, C, -D)
		} else if B > 0 && C <= 0 && D >= 0 {
			fmt.Printf("y = x³ + %.0fx² - %.0fx + %.0f\n", B, +C, D)
		} else if B < 0 && C <= 0 && D <= 0 {
			fmt.Printf("y = x³ - %.0fx² + %.0fx - %.0f\n", -B, -C, -D)
		} else {
			fmt.Printf("y = x³ + %.0fx² + %.0fx + %.0f\n", B, C, D)
		}

	} else {
		if B < 0 && C >= 0 && D >= 0 {
			fmt.Printf("y = %.0fx³ - %.0fx² + %.0fx + %.0f\n", A, -B, C, D)
		} else if B > 0 && C <= 0 && D >= 0 {
			fmt.Printf("y = %.0fx³ + %.0fx² - %.0fx + %.0f\n", A, B, -C, D)
		} else if B > 0 && C >= 0 && D <= 0 {
			fmt.Printf("y = %.0fx³ + %.0fx² + %.0fx - %.0f\n", A, B, C, -D)
		} else if B < 0 && C <= 0 && D >= 0 {
			fmt.Printf("y = %.0fx³ - %.0fx² - %.0fx + %.0f\n", A, -B, -C, D)
		} else if B > 0 && C <= 0 && D <= 0 {
			fmt.Printf("y = %.0fx³ + %.0fx² - %.0fx - %.0f\n", A, B, -C, -D)
		} else if B < 0 && C >= 0 && D <= 0 {
			fmt.Printf("y = %.0fx³ - %.0fx² + %.0fx - %.0f\n", A, -B, C, -D)
		} else if B > 0 && C <= 0 && D >= 0 {
			fmt.Printf("y = %.0fx³ + %.0fx² - %.0fx + %.0f\n", A, B, +C, D)
		} else if B < 0 && C <= 0 && D <= 0 {
			fmt.Printf("y = %.0fx³ - %.0fx² + %.0fx - %.0f\n", A, -B, -C, -D)
		} else {
			fmt.Printf("y = %.0fx³ + %.0fx² + %.0fx + %.0f\n", A, B, C, D)
		}
	}
}
