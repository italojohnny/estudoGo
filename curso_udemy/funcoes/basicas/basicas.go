package main

import "fmt"

func f5() (string, string) {
	return "retorno 1", "retorno 2"
}

func f4(p1, p2 string) string {
	return fmt.Sprintf("F4: %s %s", p1, p2)
}

func f3() string {
	return "F3"
}

func f2(p1 string, p2 string) {
	fmt.Printf("F2: %s %s\n", p1, p2)
}

func f1() {
	fmt.Println("F1")
}

func main() {
	f1()
	f2("parametro1", "parametro2")
	r3, r4 := f3(), f4("parametro1", "parametro2")
	fmt.Println(r3)
	fmt.Println(r4)

	r51, r52 := f5()
	fmt.Println("F5:", r51, r52)
}