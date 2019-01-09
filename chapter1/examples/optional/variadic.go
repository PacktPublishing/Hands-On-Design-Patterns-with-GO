package chapter1

type stats struct {
	length  int
	sum     int
	average float64
}

func variadic(vals ...int) stats {
	length := len(vals)

	sum := 0
	for _, val := range vals {
		sum += val
	}

	denominator := float64(length)
	if denominator == 0 {
		denominator = 1
	}

	average := float64(sum) / denominator

	return stats{
		length:  length,
		sum:     sum,
		average: average,
	}
}
