package main

import "fmt"
import "runtime/debug"

func main () {
	f1()
}

func f3 () {
	debug.PrintStack()
	fmt.Println("teste")
}

func f2 () {
	f3()
}

func f1 () {
	f2()
}
