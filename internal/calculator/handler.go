package calculator

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/LexusEgorov/api-calculator/internal/models"
)

type CalcHandler struct {
	controller calcController
	logger     *logrus.Logger
}

func New(logger *logrus.Logger, cache Cacher, storage Storager) *CalcHandler {
	return &CalcHandler{
		controller: newController(logger, cache, storage),
		logger:     logger,
	}
}

func (c CalcHandler) HandleSum(w http.ResponseWriter, r *http.Request) {
	uID := r.Header.Get("Authorization")
	body, err := c.getBody(r.Body)

	if err != nil {
		c.sendBadResponse(w, err)
		return
	}

	res, err := c.controller.Sum(uID, *body)

	if err != nil {
		c.sendBadResponse(w, err)
		return
	}

	c.sendGoodResponse(w, res)
}

func (c CalcHandler) HandleMult(w http.ResponseWriter, r *http.Request) {
	uID := r.Header.Get("Authorization")
	body, err := c.getBody(r.Body)

	if err != nil {
		c.sendBadResponse(w, err)
		return
	}

	res, err := c.controller.Mult(uID, *body)

	if err != nil {
		c.sendBadResponse(w, err)
		return
	}

	c.sendGoodResponse(w, res)
}

func (c CalcHandler) HandleHistory(w http.ResponseWriter, r *http.Request) {
	uID := r.Header.Get("Authorization")

	c.sendGoodResponse(w, c.controller.History(uID))
}

func (c CalcHandler) getBody(body io.ReadCloser) (*models.Input, error) {
	defer func() {
		err := body.Close()

		if err != nil {
			c.logger.Errorf("calcHandler.getBody: %v", err)
		}
	}()
	rawBody, err := io.ReadAll(body)

	if err != nil {
		c.logger.Errorf("calcHandler.getBody: %v", err)
		return nil, err
	}

	var inputNums models.Input

	err = json.Unmarshal(rawBody, &inputNums)

	if err != nil {
		c.logger.Errorf("calcHandler.getBody: %v", err)
		return nil, err
	}

	return &inputNums, nil
}

func (c CalcHandler) sendBadResponse(w http.ResponseWriter, err error) {
	c.logger.Errorf("badResponse: %v", err)
	body, err := json.Marshal(models.ErrorResponse{
		Error: err.Error(),
	})

	if err != nil {
		c.logger.Errorf("calcHandler.sendBadResponse: %v", err)
		c.sendInternalResponse(w)
		return
	}

	w.WriteHeader(http.StatusBadRequest)
	_, err = w.Write(body)

	if err != nil {
		c.logger.Errorf("calcHandler.sendBadResponse: %v", err)
	}
}

func (c CalcHandler) sendInternalResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
}

func (c CalcHandler) sendGoodResponse(w http.ResponseWriter, body interface{}) {
	res, err := json.Marshal(body)
	if err != nil {
		c.logger.Errorf("calcHandler.sendGoodResponse: %v", err)
		c.sendInternalResponse(w)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	_, err = w.Write(res)

	if err != nil {
		c.logger.Errorf("calcHandler.sendBadResponse: %v", err)
	}
}
