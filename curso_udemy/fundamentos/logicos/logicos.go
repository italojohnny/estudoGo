package main

import "fmt"

func compras(trampo1, trampo2 bool) (bool, bool, bool) {
	comprarTv50 := trampo1 && trampo2
	comprarTv32 := trampo1 != trampo2 // ou excluivo
	comprarSorvete := trampo1 || trampo2
	return comprarTv50, comprarTv32, comprarSorvete
}

func main() {
	var tv50, tv32, sorvete bool

	fmt.Println("\nCaso 1")
	tv50, tv32, sorvete = compras(true, true)
	fmt.Printf("tv50: %t,\t tv32: %t,\t sorvete: %t,\t\tsaudavel: %t\n", tv50, tv32, sorvete, !sorvete)

	fmt.Println("\nCaso 2")
	tv50, tv32, sorvete = compras(true, false)
	fmt.Printf("tv50: %t,\t tv32: %t,\t sorvete: %t,\t\tsaudavel: %t\n", tv50, tv32, sorvete, !sorvete)

	fmt.Println("\nCaso 2")
	tv50, tv32, sorvete = compras(false, true)
	fmt.Printf("tv50: %t,\t tv32: %t,\t sorvete: %t,\t\tsaudavel: %t\n", tv50, tv32, sorvete, !sorvete)

	fmt.Println("\nCaso 4")
	tv50, tv32, sorvete = compras(false, false)
	fmt.Printf("tv50: %t,\t tv32: %t,\t sorvete: %t,\tsaudavel: %t\n", tv50, tv32, sorvete, !sorvete)
}
