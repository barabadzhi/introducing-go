package main

import "fmt"

func greatest(numbers ...float64) float64 {
	greatestNumber := numbers[0]
	for _, number := range numbers {
		if number > greatestNumber {
			greatestNumber = number
		}
	}
	return greatestNumber
}

func main() {
	fmt.Println(greatest(34, 17, 89, 18, 5, 13, 11))
}
