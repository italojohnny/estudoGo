package main

import "fmt"
import "reflect"
import "math"

func main () {
	// numeros inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro e", reflect.TypeOf(3200))

	// sem sinal (so positivos)... uint8, uint16, uint32, uint64
	var b byte = 255
	fmt.Println("O byte e", reflect.TypeOf(b))

	// com sinal... int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("O valor maximo	do int e", i1)
	fmt.Println("O tipo de i1 e", reflect.TypeOf(i1))

	var i2 rune = 'a' // representa um pameamento da tabela Unicode (int32)
	fmt.Println("O rune e", reflect.TypeOf(i2))
	fmt.Println(i2)

	// numeros reais (float32, float64
	var x float32 = 49.99
	//var x = float32(49.99)
	fmt.Println("O tipo de x e", reflect.TypeOf(x))
	fmt.Println("O tipo literal 49.99 e", reflect.TypeOf(49.99))


	// boolean
	bo := true
	fmt.Println("O tipo de bo e", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// string
	s1 := "Ola meu nome e italo"
	fmt.Println("O tipo de s1 e", reflect.TypeOf(s1))
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string e", len(s1))

	// string com multipla linhsa
	s2 := `Ola
	meu
	nome
	e
	Italo`
	fmt.Println("O tamanho da string e", len(s2))

	// char??
	//var x char = 'b' // erro, nao exite tipo char
	char := 'a'
	fmt.Println("O tipo de char e", reflect.TypeOf(char))
	fmt.Println(char)
}
