package clevercalc

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/LexusEgorov/api-calculator/internal/models"
	"github.com/sirupsen/logrus"
)

type Calculator struct {
	priorityMap *priority
	parser      *parser
	logger      *logrus.Logger
}

func New(logger *logrus.Logger) *Calculator {
	return &Calculator{
		priorityMap: newPriority(),
		parser:      newParser(),
		logger:      logger,
	}
}

func (c Calculator) Compute(input string) (float64, error) {
	parsed, err := c.parser.parse(input)
	if err != nil {
		c.logger.Errorf("clevercalc.Compute: %v", err)
		return 0, err
	}

	result, err := c.compute(parsed)

	if err != nil {
		c.logger.Errorf("clevercalc.Compute: %v", err)
		return 0, err
	}

	return result, nil
}

// Вычисляет постфиксную нотацию
func (c Calculator) compute(parsed []string) (float64, error) {
	numsStack := Stack{}

	for _, value := range parsed {
		//Кладем число в стек
		if c.priorityMap.Get(value) == models.NotOpRank {
			numsStack.Push(value)
			continue
		}

		num2, err := getNum(&numsStack)

		if err != nil {
			return 0, err
		}

		num1, err := getNum(&numsStack)

		if err != nil {
			if !errors.Is(err, ErrEmpty) {
				return 0, err
			}

			num1 = 0
		}

		//Вычисляем
		res, err := compute(num1, num2, value)

		if err != nil {
			return 0, err
		}

		//Результат в стек
		numsStack.Push(fmt.Sprintf("%f", res))
	}

	res, err := getNum(&numsStack)

	//Проверка на правильность вычислений
	if err != nil || numsStack.Size() != 0 {
		return 0, models.ErrBadInput
	}

	return res, nil
}

// математика происходит тут
func compute(a, b float64, operation string) (float64, error) {
	switch operation {
	case models.OperationSum:
		return a + b, nil
	case models.OperationSub:
		return a - b, nil
	case models.OperationMult:
		return a * b, nil
	case models.OperationDiv:
		if b == 0 {
			return 0, models.ErrBadDivide
		}

		return a / b, nil
	case models.OperationPow:
		return math.Pow(a, b), nil
	default:
		return 0, models.NewErrUnknownAction(operation)
	}
}

// Получает число из стека в виде строки, переводит во float64
func getNum(from *Stack) (float64, error) {
	stringed, err := from.Pop()

	if err != nil {
		return 0, err
	}

	num, err := strconv.ParseFloat(stringed, 64)

	if err != nil {
		return 0, err
	}

	return num, nil
}
