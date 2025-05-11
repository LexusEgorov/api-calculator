package calculator

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

type Storager interface {
	SaveSum(input string, result float64) error
	SaveMult(input string, result float64) error
	GetSum(input string) (float64, error)
	GetMult(input string) (float64, error)
}

type calculator struct {
	logger  *logrus.Logger
	storage Storager
}

func (c calculator) Sum(input string) (float64, error) {
	res, err := c.storage.GetSum(input)

	if err == nil {
		return res, nil
	}

	var sum float64

	nums := strings.Split(input, ",")

	for _, num := range nums {
		parsedNum, err := strconv.ParseFloat(num, 64)

		if err != nil {
			c.logger.Errorf("calculator.sum: can't parse '%s': %v", num, err)
			return 0, fmt.Errorf("failed to parse number '%s': %w", num, err)
		}

		sum += parsedNum
	}

	c.storage.SaveSum(input, sum)
	return sum, nil
}

func (c calculator) Mult(input string) (float64, error) {
	res, err := c.storage.GetMult(input)

	if err == nil {
		return res, nil
	}

	var mult float64 = 1

	nums := strings.Split(input, ",")

	for _, num := range nums {
		parsedNum, err := strconv.ParseFloat(num, 64)

		if err != nil {
			c.logger.Errorf("calculator.mult: can't parse '%s': %v", num, err)
			return 0, fmt.Errorf("failed to parse number '%s': %w", num, err)
		}

		mult *= parsedNum
	}

	c.storage.SaveMult(input, mult)
	return mult, nil
}
