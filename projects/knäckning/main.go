package main

import . "math"

// Column represent the values for the column to calculate buckling of a column.
type Column struct {
	YieldStrength    float64    // N/mm2
	ElasticModulus   float64    // N/mm2
	Length           float64    // mm
	FasteningCase    Fastening  // only one value should be positive.
	ColumnType       ColumnType // the type of column.
	CrossSectionData CrossSectionData
}

// Fastening provides the three fastening cases.
type Fastening struct {
	First  bool // Imovable on one side and free on the other.
	Second bool // Articulated on both sides.
	Third  bool // Imovable on one side and articculated on the other.
	Forth  bool // Imovable on both sides of the column.
}

// ColumnType type defines what type of columnwe are working with.
type ColumnType struct {
	Circle          bool
	CircularPipe    bool
	Rectangle       bool
	RectangularPipe bool
}

// CrossSectionData houses all the stuff like Diameter and cross section lengths.
type CrossSectionData struct {
	CircleDiameter         float64
	OuterCircleDiamater    float64
	InnerCircleDiameter    float64
	RectangleSideLong      float64
	RectangleSideShort     float64
	RectangleWallThickness float64
}

// checkValidBuckling is a private function to check if the buckling theory is valid. Make method or something later.
func checkValidBuckling(column *Column) bool {

	// Check what case should use for the free buckling length.
	var BucklingLength float64
	if column.FasteningCase.First {
		BucklingLength = 2 * column.Length
	} else if column.FasteningCase.Second {
		BucklingLength = column.Length
	} else if column.FasteningCase.Third {
		BucklingLength = 0.7 * column.Length
	} else if column.FasteningCase.Forth {
		BucklingLength = 0.5 * column.Length
	}

	var Imin float64
	if column.ColumnType.Circle {
		Imin = (Pi * Pow(column.CrossSectionData.CircleDiameter, 4)) / 64
	} else if column.ColumnType.CircularPipe {
		Imin = (Pi / 64) * (Pow(column.CrossSectionData.OuterCircleDiamater, 4) - Pow(column.CrossSectionData.InnerCircleDiameter, 4))
	} else if column.ColumnType.Rectangle {
		Imin = (Pow(column.CrossSectionData.RectangleSideShort, 3) * column.CrossSectionData.RectangleSideLong) / 12
	} else if column.ColumnType.RectangularPipe {
		Imin = ((Pow(column.CrossSectionData.RectangleSideShort, 3) * column.CrossSectionData.RectangleSideLong) / 12) - ((Pow(column.CrossSectionData.RectangleSideShort-2*column.CrossSectionData.RectangleWallThickness, 3)*column.CrossSectionData.RectangleSideLong - (2 * column.CrossSectionData.RectangleWallThickness)) / 12)
	}

	var Area float64
	if column.ColumnType.Circle {
		Area = Pi * Pow(0.5*column.CrossSectionData.CircleDiameter, 2)
	} else if column.ColumnType.CircularPipe {
		Area = Pi * (Pow(0.5*column.CrossSectionData.OuterCircleDiamater, 2) - Pow(0.5*column.CrossSectionData.OuterCircleDiamater, 2))
	} else if column.ColumnType.Rectangle {
		Area = column.CrossSectionData.RectangleSideShort * column.CrossSectionData.RectangleSideLong
	} else if column.ColumnType.RectangularPipe {
		Area = column.CrossSectionData.RectangleSideShort*column.CrossSectionData.RectangleSideLong - (column.CrossSectionData.RectangleSideShort-2*column.CrossSectionData.RectangleWallThickness)*(column.CrossSectionData.RectangleSideLong-2*column.CrossSectionData.RectangleWallThickness)
	}

	Lambda := BucklingLength / Sqrt(Imin/Area)
	LambdaZero := Sqrt((Pi * Pi * column.ElasticModulus) / (0.99 * column.YieldStrength))

	if Lambda > LambdaZero {
		return true
	}

	return false
}

func main() {
	stang := Column{YieldStrength: 275, ElasticModulus: 105000, Length: 2400}
}
