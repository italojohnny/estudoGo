package main

import (
	"fmt"
	"time"
)

func main() {
	//a, b := "banana", "banana"
	a, b := 2, 3

	fmt.Println("igualdade:       a == b", a == b)
	fmt.Println("diferenca:       a != b", a != b)
	fmt.Println("maior que:       a  > b", a > b)
	fmt.Println("menor que:       a  < b", a < b)
	fmt.Println("maior igual que: a >= b", a >= b)
	fmt.Println("menor igual que: a <= b", a <= b)

	d1, d2 := time.Unix(0, 0), time.Unix(0, 0)
	fmt.Println("\nComparacao de datas...")
	fmt.Println("igualdade entre datas    d1 == d2 ", d1 == d2)
	fmt.Println("igualdade entre datas d1.Equal(d2)", d1.Equal(d2))

	type Pessoa struct {
		Nome string
	}
	p1 := Pessoa{"italo"}
	p2 := Pessoa{"italo"}
	fmt.Println("\nComparacao de structs")
	fmt.Println("igualdade entre structs p1 == p2", p1 == p2)

}
