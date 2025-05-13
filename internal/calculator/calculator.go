package calculator

func sum(nums []float64) float64 {
	var sum float64

	for _, num := range nums {
		sum += num
	}

	return sum
}

func mult(nums []float64) float64 {
	var mult float64 = 1

	for _, num := range nums {
		mult *= num
	}

	return mult
}
