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

// HandleSum godoc
// @Summary      Sum numbers from request's body
// @Tags         sum
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "User id"
// @Success      200  {object} models.CalcAction
// @Failure      400 {object} models.ErrorResponse
// @Failure      401 {object} models.ErrorResponse
// @Router       /sum [post]
func (e CalcHandler) HandleSum(ctx echo.Context) error {
	uID := ctx.Request().Header.Get("Authorization")
	body, err := e.getBody(ctx.Request().Body)

	if err != nil {
		e.sendBadResponse(ctx, err)
		return err
	}

	res, err := e.controller.Sum(uID, *body)

	if err != nil {
		e.sendBadResponse(ctx, err)
		return err
	}

	return e.sendGoodResponse(ctx, res)
}

// HandleMult godoc
// @Summary      Mult numbers from request's body
// @Tags         mult
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "User id"
// @Success      200  {object} models.CalcAction
// @Failure      400 {object} models.ErrorResponse
// @Failure      401 {object} models.ErrorResponse
// @Router       /mult [post]
func (e CalcHandler) HandleMult(ctx echo.Context) error {
	uID := ctx.Request().Header.Get("Authorization")
	body, err := e.getBody(ctx.Request().Body)

	if err != nil {
		e.sendBadResponse(ctx, err)
		return err
	}

	res, err := e.controller.Mult(uID, *body)

	if err != nil {
		e.sendBadResponse(ctx, err)
		return err
	}

	return e.sendGoodResponse(ctx, res)
}

// HandleCalculate godoc
// @Summary      Calculating completed math expression
// @Tags         calculate
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "User id"
// @Success      200  {object} models.CalcAction
// @Failure      400 {object} models.ErrorResponse
// @Failure      401 {object} models.ErrorResponse
// @Router       /calc [post]
func (e CalcHandler) HandleCalculate(ctx echo.Context) error {
	uID := ctx.Request().Header.Get("Authorization")
	body, err := e.getBody(ctx.Request().Body)

	if err != nil {
		e.sendBadResponse(ctx, err)
		return err
	}

	res, err := e.controller.Calculate(uID, *body)

	if err != nil {
		e.sendBadResponse(ctx, err)
		return err
	}

	return e.sendGoodResponse(ctx, res)
}

// HandleHistory godoc
// @Summary      Show history
// @Tags         history
// @Produce      json
// @Param        Authorization header string true "User id"
// @Success      200  {array} models.CalcAction
// @Failure      401 {object} models.ErrorResponse
// @Router       /history [get]
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

func (e CalcHandler) sendBadResponse(ctx echo.Context, err error) {
	badResponse := models.ErrorResponse{
		Error: err.Error(),
	}

	ctx.JSON(echo.ErrBadRequest.Code, badResponse)
}

func (e CalcHandler) sendGoodResponse(ctx echo.Context, body any) error {
	return ctx.JSON(http.StatusOK, body)
}
