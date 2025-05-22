package calculator

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/LexusEgorov/api-calculator/internal/models"
	"github.com/labstack/echo/v4"
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

func (e CalcHandler) HandleSum(ctx echo.Context) error {
	uID := ctx.Request().Header.Get("Authorization")
	body, err := e.getBody(ctx.Request().Body)

	if err != nil {
		e.sendBadResponse(ctx)
		return err
	}

	res, err := e.controller.Sum(uID, *body)

	if err != nil {
		e.sendBadResponse(ctx)
		return err
	}

	return e.sendGoodResponse(ctx, res)
}

func (e CalcHandler) HandleMult(ctx echo.Context) error {
	uID := ctx.Request().Header.Get("Authorization")
	body, err := e.getBody(ctx.Request().Body)

	if err != nil {
		e.sendBadResponse(ctx)
		return err
	}

	res, err := e.controller.Mult(uID, *body)

	if err != nil {
		e.sendBadResponse(ctx)
		return err
	}

	return e.sendGoodResponse(ctx, res)
}

func (e CalcHandler) HandleHistory(ctx echo.Context) error {
	uID := ctx.Request().Header.Get("Authorization")

	return e.sendGoodResponse(ctx, e.controller.History(uID))
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

func (e CalcHandler) sendBadResponse(ctx echo.Context) {
	ctx.Response().WriteHeader(echo.ErrBadRequest.Code)
}

func (e CalcHandler) sendGoodResponse(ctx echo.Context, body any) error {
	return ctx.JSON(http.StatusOK, body)
}
