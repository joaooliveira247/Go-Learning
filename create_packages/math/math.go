package math

// Return average of a float64 slice 
func Avg(xs []float64) float64 {
	total := float64(0)
	for _, v := range xs {
		total += v
	}

	return total
}