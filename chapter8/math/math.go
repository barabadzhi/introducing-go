package math

// Finds the average of a series of numbers
func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	if len(xs) != 0 {
		return total / float64(len(xs))
	} else {
		return float64(0)
	}
}

// Finds the minimum number of a series of numbers
func Min(xs []float64) float64 {
	var min float64
	if len(xs) != 0 {
		min = xs[0]
	}
	for _, x := range xs {
		if min > x {
			min = x
		}
	}
	return min
}

// Finds the maximum number of a series of numbers
func Max(xs []float64) float64 {
	var max float64
	if len(xs) != 0 {
		max = xs[0]
	}
	for _, x := range xs {
		if max < x {
			max = x
		}
	}
	return max
}
