package math

func Average(xd []int) int {
	total := 0
	for _, v := range xd {
		total += v
	}
	return total / len(xd)
}

func Double(x int) int {
	return x * x
}
