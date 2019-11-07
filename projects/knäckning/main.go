package main

import . "math"
import "fmt"

// Column represent the values for the column to calculate buckling of a column.
type Column struct {
	YieldStrength  float64      // N/mm2
	ElasticModulus float64      // N/mm2
	Length         float64      // mm
	EulerCase      Fastening    // Which fastening case that is used. (Only one should be passed as true and the second one will be used if none is provided)
	ColumnType     Type         // Defines what type of column that we are working with. (Only one should be passed as true)
	CrossSection   CrossSection // Defines the dimensional data for our cross section.
}

// Fastening provides the three fastening cases.
type Fastening struct {
	First  bool // Imovable on one side and free on the other.
	Second bool // Articulated on both sides.
	Third  bool // Imovable on one side and articculated on the other.
	Forth  bool // Imovable on both sides of the column.
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

// checkValidBuckling is a private function to check if the buckling theory is valid. Make method or something later.
func checkValidBuckling(column *Column) bool {

	// Check what case should use for the free buckling length. Use case two when no case is provided.
	var BucklingLength float64
	if column.EulerCase.First {
		BucklingLength = 2 * column.Length
	} else if column.EulerCase.Third {
		BucklingLength = 0.7 * column.Length
	} else if column.EulerCase.Forth {
		BucklingLength = 0.5 * column.Length
	} else {
		BucklingLength = column.Length
	}

	var Imin float64
	if column.ColumnType.Circle {
		Imin = (Pi * Pow(column.CrossSection.CircleDiameter, 4)) / 64
	} else if column.ColumnType.CircularPipe {
		Imin = (Pi / 64) * (Pow(column.CrossSection.OuterCircleDiameter, 4) - Pow(column.CrossSection.InnerCircleDiameter, 4))
	} else if column.ColumnType.Rectangle {
		Imin = (Pow(column.CrossSection.RectSideShort, 3) * column.CrossSection.RectSideLong) / 12
	} else if column.ColumnType.RectangularPipe {
		Imin = ((Pow(column.CrossSection.RectSideShort, 3) * column.CrossSection.RectSideLong) / 12) - ((Pow(column.CrossSection.RectSideShort-2*column.CrossSection.RectWallThickness, 3)*column.CrossSection.RectSideLong - (2 * column.CrossSection.RectWallThickness)) / 12)
	}

	var Area float64
	if column.ColumnType.Circle {
		Area = Pi * Pow(0.5*column.CrossSection.CircleDiameter, 2)
	} else if column.ColumnType.CircularPipe {
		Area = Pi * (Pow(0.5*column.CrossSection.OuterCircleDiameter, 2) - Pow(0.5*column.CrossSection.InnerCircleDiameter, 2))
	} else if column.ColumnType.Rectangle {
		Area = column.CrossSection.RectSideShort * column.CrossSection.RectSideLong
	} else if column.ColumnType.RectangularPipe {
		Area = column.CrossSection.RectSideShort*column.CrossSection.RectSideLong - (column.CrossSection.RectSideShort-2*column.CrossSection.RectWallThickness)*(column.CrossSection.RectSideLong-2*column.CrossSection.RectWallThickness)
	}

	Lambda := BucklingLength / Sqrt(Imin/Area)
	LambdaZero := Sqrt((Pi * Pi * column.ElasticModulus) / (0.99 * column.YieldStrength))

	if Lambda > LambdaZero {
		return true
	}

	return false
}

func main() {
	stang := &Column{YieldStrength: 275, ElasticModulus: 105000, Length: 2400, EulerCase: Fastening{Second: true}, ColumnType: Type{Rectangle: true}, CrossSection: CrossSection{RectSideShort: 30, RectSideLong: 50, RectWallThickness: 2.6}}
	fmt.Println(checkValidBuckling(stang))
}
