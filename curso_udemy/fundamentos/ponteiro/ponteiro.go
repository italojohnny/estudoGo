package main

import "fmt"

func main() {
	var p *int
	i := 1

	p = &i // pegando o endereco da variavel
	*p++   // desreferenciando (pegando o valor)
	i++

	//p++ // Erro. Nao eh permitido operacao aritmetica com ponteiro
	fmt.Println(*p, &p, p, i, &i)
}
