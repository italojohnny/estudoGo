package main

import "fmt"

func inc1 (n int) {
	n++
}

// revisao: um ponteiro e representado por um *
func inc2 (n *int) {
	// revisao: * e usado para desreferenciar, ou seja,
	// ter acesso ao valor no qual o ponteiro aponta
	*n++
}

func main () {
	n := 1

	inc1(n) // por copia de valor
	fmt.Println(n)

	// & usado para obter o endereco
	inc2(&n) // por referencia
	fmt.Println(n)
}
