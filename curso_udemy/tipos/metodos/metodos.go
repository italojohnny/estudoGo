package main

import "fmt"
import "strings"

type pessoa struct {
	nome string
	sobrenome string
}

func (p pessoa) getNomeCompleto () string {
	return p.nome + "  " + p.sobrenome
}

//func (p pessoa) setNomeCompleto (nomeCompleto string) { // copia de valor
func (p *pessoa) setNomeCompleto (nomeCompleto string) { // referencia do valor
	partes := strings.Split(nomeCompleto, " ")
	p.nome = partes[0]
	p.sobrenome = partes[1]
}

func main () {
	p1 := pessoa{"Pedro", "Silva"}
	fmt.Println(p1.getNomeCompleto())

	p1.setNomeCompleto("Maria Pereira")
	fmt.Println(p1.getNomeCompleto())

}
