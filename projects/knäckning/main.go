package main

import (
	"errors"
	"fmt"
	"math"
)

// Column represent the values for the column to calculate buckling of a c.
type Column struct {
	// Constant types we need to have:
	YieldStrength  float64 // N/mm2
	ElasticModulus float64 // N/mm2
	Length         float64 // mm
	ColumnForce    float64 // N
	BucklingSafety float64 // No unit.

	// Handling of extra data types that we need:
	EulerCase    Fastening    // Which fastening case that is used. (Only one should be passed as true and the second one will be used if none is provided)
	ColumnType   Type         // Defines what type of column that we are working with. (Only one should be passed as true)
	CrossSection CrossSection // Defines the dimensional data for our cross section.

	// The members below won't need to be entered, they will be calculated from other data if not passed to the type:
	BucklingLength float64
	Area           float64
	Imin           float64
	BucklingForce  float64
}

// Fastening provides the three fastening cases.
type Fastening struct {
	First  bool // Imovable on one side and free on the other.
	Second bool // Articulated on both sides.
	Third  bool // Imovable on one side and articculated on the other.
	Forth  bool // Imovable on both sides of the c.
}

// Type defines what type of column are working with.
type Type struct {
	Circle          bool // Tells that we have a circle.
	CircularPipe    bool // Tells that we have a circular pipe.
	Rectangle       bool // Tells that we have a rectangle.
	RectangularPipe bool // Tells that we have a rectangular pipe.
}

// CrossSection houses all the stuff like Diameter and cross section lengths.
type CrossSection struct {
	CircleDiameter      float64
	OuterCircleDiameter float64
	InnerCircleDiameter float64
	RectSideLong        float64
	RectSideShort       float64
	RectWallThickness   float64
}

// CheckValidBuckling checks if the can use the bucklign theory.
func (c Column) CheckValidBuckling() bool {
	// Calulate the lambda number for our c.
	Lambda := c.BucklingLength / math.Sqrt(c.Imin/c.Area)

	// Calculate the lowest possible lambda number for our c.
	LambdaZero := math.Sqrt((math.Pi * math.Pi * c.ElasticModulus) / (0.99 * c.YieldStrength))

	// Return true if Lambda is bigger than LambdaZero, that means we can use Eulers buckling theories.
	if Lambda > LambdaZero {
		return true
	}

	return false
}

// Buckling calculates weather a given column will buckle and break.
func Buckling(c *Column) (bool, error) {

	if c.ColumnForce == 0 || c.ElasticModulus == 0 || c.Length == 0 {
		return false, errors.New("you need to provide all the given data")
	}

	// Check what case should use for the free buckling length. Use case two when no case is provided.
	if c.BucklingLength == 0 {
		if c.EulerCase.First {
			c.BucklingLength = 2 * c.Length
		} else if c.EulerCase.Third {
			c.BucklingLength = 0.7 * c.Length
		} else if c.EulerCase.Forth {
			c.BucklingLength = 0.5 * c.Length
		} else {
			c.BucklingLength = c.Length
		}
	}

	// Calculate the minimum value of I, called Imin.
	if c.Imin == 0 {
		if c.ColumnType.Circle {
			c.Imin = (math.Pi * math.Pow(c.CrossSection.CircleDiameter, 4)) / 64
		} else if c.ColumnType.CircularPipe {
			c.Imin = (math.Pi / 64) * (math.Pow(c.CrossSection.OuterCircleDiameter, 4) - math.Pow(c.CrossSection.InnerCircleDiameter, 4))
		} else if c.ColumnType.Rectangle {
			c.Imin = (math.Pow(c.CrossSection.RectSideShort, 3) * c.CrossSection.RectSideLong) / 12
		} else if c.ColumnType.RectangularPipe {
			c.Imin = ((math.Pow(c.CrossSection.RectSideShort, 3) * c.CrossSection.RectSideLong) - (math.Pow(c.CrossSection.RectSideShort-2*c.CrossSection.RectWallThickness, 3) * (c.CrossSection.RectSideLong - (2 * c.CrossSection.RectWallThickness)))) / 12
		}
	}

	// Calculate the Area of our column depending on which type of column we have.
	if c.Area == 0 {
		if c.ColumnType.Circle {
			c.Area = math.Pi * math.Pow(0.5*c.CrossSection.CircleDiameter, 2)
		} else if c.ColumnType.CircularPipe {
			c.Area = math.Pi * (math.Pow(0.5*c.CrossSection.OuterCircleDiameter, 2) - math.Pow(0.5*c.CrossSection.InnerCircleDiameter, 2))
		} else if c.ColumnType.Rectangle {
			c.Area = c.CrossSection.RectSideShort * c.CrossSection.RectSideLong
		} else if c.ColumnType.RectangularPipe {
			c.Area = c.CrossSection.RectSideShort*c.CrossSection.RectSideLong - (c.CrossSection.RectSideShort-2*c.CrossSection.RectWallThickness)*(c.CrossSection.RectSideLong-2*c.CrossSection.RectWallThickness)
		}
	}

	// Check if it is possible to actually use buckling theories of the great Euler.
	if !c.CheckValidBuckling() {
		return false, errors.New("we cant use Eulers theory of buckling in this example")
	}

	if c.BucklingForce == 0 {
		c.BucklingForce = math.Pow(math.Pi, 2) * ((c.ElasticModulus * c.Imin) / math.Pow(c.BucklingLength, 2))
	}

	BucklingSafetyFactor := c.BucklingForce / c.ColumnForce

	if BucklingSafetyFactor > 1 {
		return false, nil
	}

	return true, nil
}

func main() {
	horizontal := &Column{YieldStrength: 275, ElasticModulus: 105000, Length: 2400, EulerCase: Fastening{Second: true}, ColumnType: Type{RectangularPipe: true}, CrossSection: CrossSection{RectSideShort: 30, RectSideLong: 50, RectWallThickness: 2.6}, ColumnForce: 10000}
	sideways := &Column{YieldStrength: 275, ElasticModulus: 105000, Length: 1200 / math.Cos(45), EulerCase: Fastening{Second: true}, ColumnType: Type{RectangularPipe: true}, CrossSection: CrossSection{RectSideShort: 30, RectSideLong: 50, RectWallThickness: 2.6}, ColumnForce: 5 * math.Sqrt2}

	first, _ := Buckling(horizontal)
	second, _ := Buckling(sideways)

	fmt.Println("Den raka stången i mitten kommer knäckas:", first)
	fmt.Println("Den vinklade stången på sidan kommer knäckas:", second)

}
