package main

import "fmt"

func main () {
	//var aprovados map[int]string
	// mapas devem ser inicializado
	aprovados := make(map[int]string)

	aprovados[123456789789] = "Maria"
	aprovados[871823123132] = "Pedro"
	aprovados[900982342342] = "Ana"
	fmt.Println(aprovados)


	for cpf, nome := range aprovados  {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[900982342342])
	delete(aprovados, 900982342342)
	fmt.Println(aprovados)
	fmt.Println(aprovados[900982342342])
}
