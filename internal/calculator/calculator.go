package calculator

type calculator struct{}

func (c calculator) Sum(nums []float64) float64 {
	var sum float64

	for _, num := range nums {
		sum += num
	}

	return sum
}

func (c calculator) Mult(nums []float64) float64 {
	var mult float64 = 1

	for _, num := range nums {
		mult *= num
	}

	return mult
}
