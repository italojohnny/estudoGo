package main

import (
	"math"
	f "fmt"
)

func main () {
	const PI float64 = 3.1415
	var raio = 3.2 // tipo (float64) inferido pelo compilador

	// forma reduzida de criar uma var
	area := PI * math.Pow(raio, 2)
	f.Println("A area da circunferencia e", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)
	f.Println(a, b, c, d)

	var e, g bool = true, false
	f.Println(e, g)

	h, i, j := 2, false, "epa!"
	f.Println(h, i, j)

}
