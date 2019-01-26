package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2
	fmt.Println("Atribuicao:        a :=", a)
	fmt.Println("Atribuicao:        b :=", b)
	fmt.Println("Soma:           a + b =", a+b)
	fmt.Println("Subtracao:      a - b =", a-b)
	fmt.Println("Multipllicacao: a * b =", a*b)
	fmt.Println("Divisao:        a / b =", a/b)
	fmt.Println("Modulo:         a % b =", a%b)

	// bitwise (operacoes logicas)
	fmt.Println("\nOperacoes logicas:")
	fmt.Println("and:            a & b =", a&b) // 11 & 10 = 10
	fmt.Println("or:             a | b =", a|b) // 11 | 10 = 11
	fmt.Println("xor:            a ^ b =", a^b) // 11 ^ 10 = 01

	// ...algumas outras operacoes usando math
	fmt.Println("Operacoes com math")
	fmt.Println("Maior:       math.Max(a, b) =", math.Max(float64(a), float64(b)))
	fmt.Println("Menor:       math.Min(a, b) =", math.Min(float64(a), float64(b)))
	fmt.Println("Exponencial: math.Pow(a, b) =", math.Pow(float64(a), float64(b)))

}
