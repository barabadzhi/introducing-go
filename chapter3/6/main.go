package main

import "fmt"

func main() {
	fmt.Print("Feet: ")
	var feet float64
	fmt.Scanf("%f", &feet)

	meters := feet * 0.3048
	fmt.Printf("Meters: %.2f", meters)
}
