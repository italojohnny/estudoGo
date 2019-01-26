package main

import "fmt"

func trocar(p1, p2 int) (segundo, primeiro int) {
	segundo = p2
	primeiro = p1
	return // retorno limpo  (nos nomes das variaveis de retorno ja estao definidos)
}

func main () {
	x, y := trocar(2, 3)
	fmt.Println(x, y)

	segundo, primeiro := trocar(7, 1)
	fmt.Println(segundo, primeiro)
}
