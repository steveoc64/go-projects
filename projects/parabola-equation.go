package main

import "fmt"

// Calculate the values for A, B and C from spcific points on graph
func quadraticEquation(YaxisX, YaxisY, NullPoint1, NullPoint2 float64) (A, B, C float64) {
	A = YaxisY / ((YaxisX - NullPoint1) * (YaxisX - NullPoint2))
  B = A * (-NullPoint1 + -NullPoint2)
  C = A * (NullPoint1 * NullPoint2)
  return A, B, C
}

func main() {
	var YaxisX, YaxisY, NullPoint1, NullPoint2 float64
	fmt.Println("Enter a couple values from graph to get quadratic equation!")

  // Values for point on the line where it goes through the y-axis
	fmt.Print("\nEnter point where line goes through y-axis (x, y): ")
	fmt.Scanf("(%v, %v)", &YaxisX, &YaxisY)

  // Get values x-values for parts where y equals to zero
	fmt.Println("\nEnter x-values for points where y-value is zero:")
	fmt.Print("First x-value: ")
	fmt.Scan(&NullPoint1)

  fmt.Print("Second x-value: ")
	fmt.Scan(&NullPoint2)

  // Enter each of the outputed values in to respective variables
  A, B, C := quadraticEquation(YaxisX, YaxisY, NullPoint1, NullPoint2)

  // Make sure that it doesn't print a plus sign before a negative number
	fmt.Println("\nYou equation is:")
	if B > 0 && C > 0 {
		fmt.Printf("%vx² + %vx + %v = 0\n", A, B, C)
	} else if B < 0 && C < 0 {
		B, C = -B, -C
		fmt.Printf("%vx² - %vx - %v = 0\n", A, B, C)
	} else if B > 0 && C < 0 {
		C = -C
		fmt.Printf("%vx² + %vx - %v = 0\n", A, B, C)
	} else if B < 0 && C > 0 {
		B = -B
		fmt.Printf("%vx² - %vx + %v = 0\n", A, B, C)
	}

}
