package main

import "fmt"

// Nao tem operador ternario
func obterResultado(nota float64) string {
	// return nota >= 6 ? "Aprovado" : "Reprovado"
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"
}

func main() {
	fmt.Println(obterResultado(6.1))
	fmt.Println(obterResultado(6.0))
	fmt.Println(obterResultado(5.9))
}
