package calculator

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"

	"api-calculator/internal/models"
)

type Cacher interface {
	Save(action models.CalcAction) error
	Get(input string, action models.Action) (float64, error)
}

type Storager interface {
	Save(uID string, action models.CalcAction) error
	Get(uID string) []models.CalcAction
}

type CalcController struct {
	calc    calculator
	cache   Cacher
	storage Storager
	logger  *logrus.Logger
}

func New(logger *logrus.Logger, cache Cacher, storage Storager) *CalcController {
	return &CalcController{
		calc:    calculator{},
		cache:   cache,
		storage: storage,
		logger:  logger,
	}
}

func (c CalcController) HandleSum(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	uID := r.Header.Get("Authorization")

	if uID == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		c.logger.Errorf("calcController.HandleSum: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()

	stringedBody := string(body)
	res, err := c.cache.Get(stringedBody, models.SUM)

	if err == nil {
		c.storage.Save(uID, models.CalcAction{
			Input:  stringedBody,
			Action: models.SUM,
			Result: res,
		})

		w.Write(fmt.Appendf(nil, "%f", res))
		return
	}

	nums, err := c.prepareNums(stringedBody)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res = c.calc.Sum(nums)
	calcAction := models.CalcAction{
		Input:  stringedBody,
		Action: models.SUM,
		Result: res,
	}

	c.cache.Save(calcAction)
	c.storage.Save(uID, calcAction)

	w.Write(fmt.Appendf(nil, "%f", res))
}

func (c CalcController) HandleMult(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	uID := r.Header.Get("Authorization")

	if uID == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		c.logger.Errorf("calcController.HandleMult: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()

	stringedBody := string(body)
	res, err := c.cache.Get(stringedBody, models.MULT)

	if err == nil {
		c.storage.Save(uID, models.CalcAction{
			Input:  stringedBody,
			Action: models.MULT,
			Result: res,
		})

		w.Write(fmt.Appendf(nil, "%f", res))
		return
	}

	nums, err := c.prepareNums(stringedBody)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res = c.calc.Mult(nums)
	calcAction := models.CalcAction{
		Input:  stringedBody,
		Action: models.MULT,
		Result: res,
	}

	c.cache.Save(calcAction)
	c.storage.Save(uID, calcAction)

	w.Write(fmt.Appendf(nil, "%f", res))
}

func (c CalcController) HandleHistory(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	uID := r.Header.Get("Authorization")

	if uID == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	actions := c.storage.Get(uID)

	res, err := json.Marshal(actions)

	if err != nil {
		c.logger.Errorf("calcController.HandleHistory: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(res)
}

func (c CalcController) prepareNums(input string) ([]float64, error) {
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
