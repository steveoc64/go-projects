package main

import (
	"fmt"
	"log"
)

func betygsPoäng(bokstav string) (poäng float64) {

	switch bokstav {
	case "A":
		poäng = 20
	case "B":
		poäng = 17.5
	case "C":
		poäng = 15
	case "D":
		poäng = 12.5
	case "E":
		poäng = 10
	default:
		log.Fatalln("Inte ett giltigt betyg!")
	}

	return poäng
}

func merit(betyg int, extra float64) (merit float64) {
	fmt.Println("Skriv in ett betyg per rad:")

	var summa float64
	var bokstav string

	for i := 1; i <= betyg; i++ {
		fmt.Scanln(&bokstav)
		summa = summa + betygsPoäng(bokstav)
	}

	merit = (summa / float64(betyg)) + extra

	return merit
}

func main() {
	var betyg int
	var extra float64
	fmt.Print("Ange antal betyg från gymnasiet: ")
	fmt.Scanln(&betyg)

	fmt.Print("Antal extramerit från kurser: ")
	fmt.Scanln(&extra)

	if extra < 0 || extra > 2.5 {
		log.Fatalln("Inte ett giltigt antal extramerit!")
	}

	fmt.Printf("Meritpoäng: %.2f", merit(betyg, extra))
}
