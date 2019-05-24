package main

import (
	"fmt"
	"math"
	"os"
)

// Beräkna arean för tråden från matematiska formeln.
func threadArea(ThreadDiameter float64) float64 {
	return (math.Pi * math.Pow(ThreadDiameter, 2)) / 4
}

// Beräkna trådens vikt utifrån matematiska formeln för vikt.
func threadMass(ThreadMass, Dencity, Length float64) float64 {
	return (threadArea(ThreadMass) / 1000000) * Dencity * (Length / 1000)
}

// Beräkning av trådens tillåtna spänning.
func allowedThreadStrenght(AllowedThreadStrength, ThreadArea, Gravity, ThreadMass float64) float64 {
	return ((AllowedThreadStrength * ThreadArea) / Gravity) - ThreadMass
}

// Beräkning av trådens sträckgräns.
func threadYieldStength(ThreadYieldStength, ThreadArea, Gravity, ThreadMass float64) float64 {
	return ((ThreadYieldStength * ThreadArea) / Gravity) - ThreadMass
}

// Beräkning av trådens brottgräns.
func threadTensileStrength(ThreadTensileStrength, ThreadArea, GravityConstant, ThreadMass float64) float64 {
	return ((ThreadTensileStrength * ThreadArea) / GravityConstant) - ThreadMass
}

// Antal trådar som krävs för att lyfta vikten.
func requiredThreads(Load, ThreadStrengthValue float64) float64 {
	return math.Ceil(Load / ThreadStrengthValue)
}

// Vikten för hela vajern.
func wireMass(RequiredThreads, ThreadMass float64) float64 {
	return RequiredThreads * ThreadMass
}

// Totala arean för vajern.
func wireArea(ThreadArea, RequiredThreads float64) float64 {
	return ThreadArea * RequiredThreads
}

// Beräkning av diameter på vajern utifrån arean.
func wireDiameter(WireArea float64) float64 {
	return math.Sqrt((WireArea * 4) / math.Pi)
}

// Totala förlängningen vid belastning.
func totalExtension(WireMass, WireArea, Gravity, Length, Load, Elasticity float64) float64 {

	// Tyngden för vajern och belastningen.
	Q := WireMass * Gravity
	LoadF := Load * Gravity

	// Förlängning för vajer på grund av vajerns vik respektive belastningens vikt.
	δQ := (Q * Length) / (2 * WireArea * Elasticity)
	δF := (LoadF * Length) / (WireArea * Elasticity)

	return δF + δQ
}

// Längden på trumman när man rullar upp vajern i ett lager.
func spoolLength(SpoolDiameter, WireDiameter, Length, TotalExtension float64) float64 {
	LengthExtended := Length + TotalExtension

	// Diametern på trumman plus diametern på vajern för att få mitten av vajern på varje sida.
	TotalDiameter := SpoolDiameter + WireDiameter

	// Omkretsen för hur mycket vajer som går på ett varv följt av antalet varv som krävs för att rulla upp vajern.
	Circumreference := TotalDiameter * math.Pi
	Laps := LengthExtended / Circumreference

	return Laps * WireDiameter
}

// Längden på trumman när man rullar upp vajern i tre lager.
func threeLapsSpoolLength(SpoolDiameter, WireDiameter, Length, TotalExtension float64) float64 {
	LengthExtended := Length + TotalExtension

	// Totala diameternara för varje varv upp till tre, varje varv ska vara två * radien av vajern större för add få mittpunkten av vajer då den är lika stor vid böjning.
	TotalDiameterLap := SpoolDiameter + WireDiameter
	TotalDiameter2Laps := SpoolDiameter + (3 * WireDiameter)
	TotalDiameter3Laps := SpoolDiameter + (5 * WireDiameter)

	Circumreference := ((math.Pi * TotalDiameterLap) + (math.Pi * TotalDiameter2Laps) + (math.Pi * TotalDiameter3Laps))

	Laps := LengthExtended / Circumreference

	return Laps * WireDiameter
}

// Definiera typen Material och egenskaperna för den.
type Material struct {
	Elasticity      float64 //N/mm²
	Dencity         float64 //kg/m3
	YieldStrength   float64 //N/mm²
	TensileStrength float64 //N/mm²
	ThreadDiameter  float64 //mm
	SpoolDiameter   float64 //mm
}

func main() {
	// Konstanter för användning i beräkningarna.
	const (
		Load         = 2000    //kg
		Length       = 1000000 //mm
		Gravity      = 9.82    //m/s²
		SafetyFactor = 2
	)

	var material string

	// Alla egenskaper för materialen utifrån Material-typen.
	titan := Material{105000, 4600, 747.5, 962.5, 2.5, 500}
	cfrp := Material{107000, 1550, 800, 800, 2.5, 500}
	nylon := Material{2910, 1130, 72.4, 127.5, 5, 2000}

	fmt.Print("\nMöjliga material är titan, cfrp eller nylon.\nVälj material att använda: ")
	fmt.Scanln(&material)

	// Skapa ett val av material där materialets egenskaper kopieras till choice.
	var choice Material
	switch material {
	case "titan":
		choice = titan
	case "cfrp":
		choice = cfrp
	case "nylon":
		choice = nylon
	default:
		fmt.Println("Error: Inte ett giltigt material!")
		os.Exit(2)
	}

	// Beräkning av tillåten spänning samt definiering av trådens area och vikt.
	σTill := choice.YieldStrength / SafetyFactor
	ThreadArea := threadArea(choice.ThreadDiameter)
	ThreadMass := threadMass(choice.ThreadDiameter, choice.Dencity, Length)

	// Skriv ut säkerhetsfaktorn och tråddiametern.
	fmt.Println("\nSäkerhetsfaktor:", SafetyFactor)
	fmt.Println("Diameter på lös tråd:", choice.ThreadDiameter, "mm")

	// Definiering av variabler för tillåten styrka, sträckgräns och brottgräns utifrån funktionerna.
	AllowedThreadStrength := allowedThreadStrenght(σTill, ThreadArea, Gravity, ThreadMass)
	ThreadYieldStength := threadYieldStength(choice.YieldStrength, ThreadArea, Gravity, ThreadMass)
	ThreadTensileStrength := threadTensileStrength(choice.TensileStrength, ThreadArea, Gravity, ThreadMass)

	// Definiering av variabel för antalet trådar utifrån funktionen.
	RequiredThreads := requiredThreads(Load, AllowedThreadStrength)

	// Skriv ut belastningsvikten och antalet trådar som krävs
	fmt.Println("\nBelastningsvikt:", Load)
	fmt.Println("Trådar som krävs:", RequiredThreads, "st")

	// Konvertera vikt, area och diameter från tråden över till hela vajern med hjälp av antalet trådar som krävs och funktionerna ovan.
	WireMass := wireMass(RequiredThreads, ThreadMass)
	WireArea := wireArea(ThreadArea, RequiredThreads)
	WireDiameter := wireDiameter(WireArea)

	// Skriv ut hela vajerns vikt.
	fmt.Printf("Vajerns vikt: %.2f kg\n", WireMass)

	// Konvertera trådens tillåtna styrka, sträckgräns och brottgräns till vajer.
	WireStrengthSafe := RequiredThreads * AllowedThreadStrength
	WireStrengthYield := RequiredThreads * ThreadYieldStength
	WireStrengthTensile := RequiredThreads * ThreadTensileStrength

	// Skriv ut värdena för tillåtna sträckgränsen, sträckgränsen och brottgräns.
	fmt.Printf("\nSäker sträckgräns: %.2f kg\n", WireStrengthSafe)
	fmt.Printf("Teoretisk sträckgräns: %.2f kg\n", WireStrengthYield)
	fmt.Printf("Teoretisk brottgräns: %.2f kg\n", WireStrengthTensile)

	// Definiera totala förlängningen och skriv ut den.
	TotalExtension := totalExtension(WireMass, WireArea, Gravity, Length, Load, choice.Elasticity)
	fmt.Printf("\nTotal förlängning: %.2f mm\n", TotalExtension)

	// Beräkna procentuella förlängningen.
	Epsilon := 100 * (TotalExtension / Length)
	fmt.Printf("Procentuell förlängning: %.5f%%\n", Epsilon)

	// Definiering av längen av trumman vid ett varv och tre varv samt utskrift av svaren, inklusive trummans diameter.
	SpoolLength := spoolLength(choice.SpoolDiameter, WireDiameter, Length, TotalExtension)
	ThreeLapsSpoolLength := threeLapsSpoolLength(choice.SpoolDiameter, WireDiameter, Length, TotalExtension)
	fmt.Printf("\nDiameter på trumman: %.2f mm\n", choice.SpoolDiameter)
	fmt.Printf("Längd på trumman (ett lager): %.2f mm\n", SpoolLength)
	fmt.Printf("Längd på trumman (tre lager): %.2f mm", ThreeLapsSpoolLength)

}
