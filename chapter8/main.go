package main

import (
	"./math"
	"fmt"
)

func main() {
	xs := []float64{1, 2, 3, 4}
	fmt.Println("Slice:", xs)
	fmt.Println("Average:", math.Average(xs))
	fmt.Println("Min:", math.Min(xs))
	fmt.Println("Max:", math.Max(xs))
}
