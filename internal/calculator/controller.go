package calculator

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/LexusEgorov/api-calculator/internal/clevercalc"
	"github.com/LexusEgorov/api-calculator/internal/models"
)

type Cacher interface {
	Set(action models.CalcAction) error
	Get(input string, action models.Action) (models.CalcAction, error)
}

type Storager interface {
	Set(uID string, action models.CalcAction)
	Get(uID string) []models.CalcAction
}

type calcController struct {
	cache   Cacher
	storage Storager
	calc    *clevercalc.Calculator
	logger  *logrus.Logger
}

func newController(logger *logrus.Logger, cache Cacher, storage Storager) calcController {
	return calcController{
		cache:   cache,
		storage: storage,
		logger:  logger,
		calc:    clevercalc.New(logger),
	}
}

func (c calcController) Sum(uID string, input models.Input) (*models.CalcAction, error) {
	cached, err := c.cache.Get(input.Input, models.Sum)

	if err == nil {
		c.storage.Set(uID, cached)

		return &cached, nil
	}

	nums, err := c.prepareNums(input.Input)

	if err != nil {
		c.logger.Errorf("calcController.Sum: %v", err)
		return nil, err
	}

	res := sumNums(nums...)

	calcAction := models.CalcAction{
		Input:  input.Input,
		Action: models.Sum,
		Result: res,
	}

	err = c.cache.Set(calcAction)

	if err != nil {
		c.logger.Errorf("cache.Set: %v", err)
	}

	c.storage.Set(uID, calcAction)

	return &calcAction, nil
}

func (c calcController) Mult(uID string, input models.Input) (*models.CalcAction, error) {
	cached, err := c.cache.Get(input.Input, models.Mult)

	if err == nil {
		c.storage.Set(uID, cached)

		return &cached, nil
	}

	nums, err := c.prepareNums(input.Input)

	if err != nil {
		c.logger.Errorf("calcController.Mult: %v", err)
		return nil, err
	}

	res := multNums(nums...)

	calcAction := models.CalcAction{
		Input:  input.Input,
		Action: models.Mult,
		Result: res,
	}

	err = c.cache.Set(calcAction)

	if err != nil {
		c.logger.Errorf("cache.Set: %v", err)
	}

	c.storage.Set(uID, calcAction)

	return &calcAction, nil
}

func (c calcController) Calculate(uID string, input models.Input) (*models.CalcAction, error) {
	cached, err := c.cache.Get(input.Input, models.Calc)

	if err == nil {
		c.storage.Set(uID, cached)

		return &cached, nil
	}

	res, err := c.calc.Compute(input.Input)

	if err != nil {
		return nil, err
	}

	calcAction := models.CalcAction{
		Input:  input.Input,
		Action: models.Calc,
		Result: res,
	}

	err = c.cache.Set(calcAction)

	if err != nil {
		c.logger.Errorf("cache.Set: %v", err)
	}

	c.storage.Set(uID, calcAction)

	return &calcAction, nil
}

func (c calcController) History(uID string) []models.CalcAction {
	return c.storage.Get(uID)
}

func (c calcController) prepareNums(input string) ([]float64, error) {
	stringedNums := strings.Split(strings.ReplaceAll(input, " ", ""), ",")
	nums := make([]float64, 0)

	for _, stringed := range stringedNums {
		num, err := strconv.ParseFloat(stringed, 64)

		if err != nil {
			c.logger.Errorf("calcController.prepareNums: can't parse '%s': %v", stringed, err)
			return nil, fmt.Errorf("calcController.prepareNums: %w", err)
		}

		nums = append(nums, num)
	}

	return nums, nil
}
