package main

import "fmt"

func main() {
	fmt.Print("Fahrenheit: ")
	var fahrenheit float64
	fmt.Scanf("%f", &fahrenheit)

	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("Celsius: %.2f", celsius)
}
