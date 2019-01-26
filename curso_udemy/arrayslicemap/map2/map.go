package main

import "fmt"

func main () {
	funcsESalarios := map[string]float64 {
		"Jose Joao": 11325.45,
		"Gabriela Silva": 15434.78,
		"Pedro Junior": 1200.0,
	}

	funcsESalarios["Rafael Filho"] = 1350.0
	delete(funcsESalarios, "chaveInexistente") // nao gera erro

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
