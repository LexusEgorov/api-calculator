package calculator

import (
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()
var handlerCache = mockCacher{}
var handlerStorage = mockStorager{}
var handlerController = newController(logrus.New(), cache, storage)

var handler = &CalcHandler{
	logger:     logger,
	controller: controller,
}

func TestCalcHandler_HandleHistory(t *testing.T) {
	type args struct {
		ctx echo.Context
	}
	tests := []struct {
		name    string
		e       CalcHandler
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.e.HandleHistory(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("CalcHandler.HandleHistory() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCalcHandler_HandleSum(t *testing.T) {
	type args struct {
		ctx echo.Context
	}
	tests := []struct {
		name    string
		e       CalcHandler
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.e.HandleSum(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("CalcHandler.HandleSum() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCalcHandler_HandleMult(t *testing.T) {
	type args struct {
		ctx echo.Context
	}
	tests := []struct {
		name    string
		e       CalcHandler
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.e.HandleMult(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("CalcHandler.HandleMult() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCalcHandler_HandleCalculate(t *testing.T) {
	type args struct {
		ctx echo.Context
	}
	tests := []struct {
		name    string
		e       CalcHandler
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.e.HandleCalculate(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("CalcHandler.HandleCalculate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
