package calculator

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
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
		c.sendBadResponse(w)
		return
	}

	res, err := c.controller.Sum(uID, body)

	if err != nil {
		c.sendBadResponse(w)
		return
	}

	c.sendGoodResponse(w, res)
}

func (c CalcHandler) HandleMult(w http.ResponseWriter, r *http.Request) {
	uID := r.Header.Get("Authorization")
	body, err := c.getBody(r.Body)

	if err != nil {
		c.sendBadResponse(w)
		return
	}

	res, err := c.controller.Mult(uID, body)

	if err != nil {
		c.sendBadResponse(w)
		return
	}

	c.sendGoodResponse(w, res)
}

func (c CalcHandler) HandleHistory(w http.ResponseWriter, r *http.Request) {
	uID := r.Header.Get("Authorization")

	c.sendGoodResponse(w, c.controller.History(uID))
}

func (c CalcHandler) getBody(body io.ReadCloser) (string, error) {
	defer body.Close()
	rawBody, err := io.ReadAll(body)

	if err != nil {
		c.logger.Errorf("calcHandler.getBody: %v", err)
		return "", err
	}

	return string(rawBody), nil
}

func (c CalcHandler) sendBadResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
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
	w.Write(res)
}
