package main

import "fmt"

type imprimivel interface {
	toString () string
}

func imprimir (x imprimivel) {
	fmt.Println(x.toString())
}

type pessoa struct {
	nome string
	sobrenome string
}

type produto struct {
	nome string
	preco float64
}

// interfaces sao implementadas implicitamente
func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.nome, p.preco)
}

func main () {
	var coisa imprimivel = pessoa{"Roberto", "Silva"}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	coisa = produto{"Calca Jeans", 79.90}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	imprimir(produto{"Celular", 1500.00})
	imprimir(pessoa{"Italo", "Johnny"})
}
