package main

import "fmt"

func swap(x, y *int) {
	*x, *y = *y, *x
}

func main() {
	x, y := 1, 2
	fmt.Println("x:", x, "y:", y)
	swap(&x, &y)
	fmt.Println("x:", x, "y:", y)
}
