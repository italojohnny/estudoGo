package main

import "fmt"

func main () {
	numeros := [...] int {7, 11, 2, 141, 1,} // compilador conta

	for i, numero := range numeros {
		fmt.Printf("array[%d] == %d\n", i, numero)
	}

	for _, num := range numeros {
		fmt.Println(num)
	}

	fmt.Println(numeros[0])
}
