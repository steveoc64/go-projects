package main

import (
	"fmt"
	"math"
)

func main() {
	var E, ρ, σS, σB, D, D_trumma, L, load, nS float64
	var g = 9.82
	fmt.Print("Elasticitetsmodul (N/mm^2): ")
	fmt.Scanln(&E)
	fmt.Print("Dencitet (kg/m^3): ")
	fmt.Scanln(&ρ)
	fmt.Print("Sträckgräns (N/mm^2): ")
	fmt.Scanln(&σS)
	fmt.Print("Brottgräns (N/mm^2): ")
	fmt.Scanln(&σB)
	fmt.Print("Belastningsvikt (kg): ")
	fmt.Scanln(&load)
	fmt.Print("Diameter på tråd (mm): ")
	fmt.Scanln(&D)
	fmt.Print("Diameter på trumma (mm): ")
	fmt.Scanln(&D_trumma)
	fmt.Print("Säkerhetsfaktor: ")
	fmt.Scanln(&nS)
	fmt.Print("Längd på vajer (m): ")
	fmt.Scanln(&L)
	L = L * 1000

	σ_till := σS / nS
	A := (math.Pi * math.Pow(D, 2)) / 4
	m_lina := (A / 1000000) * ρ * (L / 1000)
	fmt.Println("\nSäkerhetsfaktor:", nS)
	fmt.Println("Diameter på lös tråd:", D, "mm")
	fmt.Println("Lös tråds egenvikt:", m_lina, "kg")

	var m_ork_nS float64 = ((σ_till * A) / g) - m_lina
	var m_ork_S float64 = ((σS * A) / g) - m_lina
	var m_ork_B float64 = ((σB * A) / g) - m_lina
	fmt.Println("\nSträckgräns för lös tråd (med säkerhetsfaktor):", m_ork_nS, "kg")
	fmt.Println("Teoretisk sträckgräns för lös tråd:", m_ork_S, "kg")
	fmt.Println("Teoretisk brottgräns för lös tråd:", m_ork_B, "kg")

	fmt.Println("\nBelastningsvikt:", load)

	var t float64 = math.Ceil(load / m_ork_nS)
	fmt.Println("Antal behöva trådar:", t, "st")

	var m_vajer float64 = t * m_lina
	var A_total float64 = A * t
	var D_vajer float64 = math.Sqrt((A_total * 4) / math.Pi)
	fmt.Println("Hela vajerns vikt:", m_vajer, "kg")

	var m_ork_total_nS float64 = t * m_ork_nS
	var m_ork_total_S float64 = t * m_ork_S
	var m_ork_total_B float64 = t * m_ork_B
	fmt.Println("\nSträckgräns för vajern (med säkerhetsfaktor):", m_ork_total_nS, "kg")
	fmt.Println("Teoretisk sträckgräns för vajer:", m_ork_total_S, "kg")
	fmt.Println("Teoretisk sbrottgräns för vajer:", m_ork_total_B, "kg")

	var Q = m_vajer * g
	var F_load = load * g
	var δF = (F_load * L) / (A_total * E)
	var δQ = (Q * L) / (2 * A_total * E)
	var δ_tot = δF + δQ
	fmt.Println("\nBelastningsförlängning:", δF/1000, "m")
	fmt.Println("Egenförlängning:", δQ/1000, "m")
	fmt.Println("Total förlängning:", δ_tot/1000, "m")

	var ε = δ_tot / L
	fmt.Println("\nProcentuell förlängning:", ε*100, "%")

	var D_varv = D_trumma + D_vajer
	var L_förlängd = L + δ_tot
	var O = D_varv * math.Pi
	var varv = L_förlängd / O
	var trumma_L = varv * D_vajer
	fmt.Println("\nDiameter på trumman:", D_trumma/1000, "m")
	fmt.Println("Längd på trumman för ett lager:", trumma_L/1000, "m")

	var D_2varv = D_trumma + (3 * D_vajer)
	var D_3varv = D_trumma + (5 * D_vajer)
	var VarvTre = L_förlängd / ((math.Pi * D_varv) + (math.Pi * D_2varv) + (math.Pi * D_3varv))
	var TrummaL3 = VarvTre * D_vajer
	fmt.Println("Längd på trumman vid tre lager:", TrummaL3/1000, "m")
}
