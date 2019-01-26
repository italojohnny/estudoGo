package main

import "time"
import "fmt"

func main() {
	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("Bom dia!")
	case t.Hour() < 18:
		fmt.Println("Bom tarde!")
	default:
		fmt.Println("Bom tarde!")
	}
}
