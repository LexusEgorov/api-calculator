package calculator

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"

	"api-calculator/internal/models"
)

type Cacher interface {
	Set(action models.CalcAction)
	Get(input string, action models.Action) (*models.CalcAction, error)
}

type Storager interface {
	Set(uID string, action models.CalcAction)
	Get(uID string) []models.CalcAction
}

type calcController struct {
	cache   Cacher
	storage Storager
	logger  *logrus.Logger
}

func newController(logger *logrus.Logger, cache Cacher, storage Storager) calcController {
	return calcController{
		cache:   cache,
		storage: storage,
		logger:  logger,
	}
}

func (c calcController) Sum(uID, input string) (*models.CalcAction, error) {
	cached, err := c.cache.Get(input, models.SUM)

	if err == nil {
		c.storage.Set(uID, *cached)

		return cached, nil
	}

	nums, err := c.prepareNums(input)

	if err != nil {
		c.logger.Errorf("calcController.Sum: %v", err)
		return nil, err
	}

	res := sum(nums)

	calcAction := models.CalcAction{
		Input:  input,
		Action: models.SUM,
		Result: res,
	}

	c.cache.Set(calcAction)
	c.storage.Set(uID, calcAction)

	return &calcAction, nil
}

func (c calcController) Mult(uID, input string) (*models.CalcAction, error) {
	cached, err := c.cache.Get(input, models.MULT)

	if err == nil {
		c.storage.Set(uID, *cached)

		return cached, nil
	}

	nums, err := c.prepareNums(input)

	if err != nil {
		c.logger.Errorf("calcController.Mult: %v", err)
		return nil, err
	}

	res := mult(nums)

	calcAction := models.CalcAction{
		Input:  input,
		Action: models.MULT,
		Result: res,
	}

	c.cache.Set(calcAction)
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
