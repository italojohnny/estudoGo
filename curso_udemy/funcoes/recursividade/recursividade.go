package main

import "fmt"

func fatorial (n int) (int, error) {
	switch {
		case n < 0:
			return -1, fmt.Errorf("Numero invalido: %d", n)
		case n == 0:
			return 1, nil
		default:
			fatorialAnterior, _ := fatorial(n -1)
			return n * fatorialAnterior, nil
	}
}

func fatorial2 (valor uint) uint {
	switch {
		case valor == 0:
			return 1
		default:
			return valor * fatorial2(valor -1)
	}
}

func main () {
	resultado := fatorial2(5)
	fmt.Println(resultado)

	resultado = fatorial2(0)
	fmt.Println(resultado)
}
