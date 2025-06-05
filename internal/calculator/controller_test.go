package calculator

import (
	"reflect"
	"testing"

	"github.com/LexusEgorov/api-calculator/internal/models"
	"github.com/LexusEgorov/api-calculator/internal/storage/cache"
	"github.com/LexusEgorov/api-calculator/internal/storage/requests"
	"github.com/sirupsen/logrus"
)

func Test_calcController_prepareNums(t *testing.T) {
	testController := newController(logrus.New(), cache.New(), requests.New())

	type args struct {
		input string
	}
	tests := []struct {
		name    string
		c       calcController
		args    args
		want    []float64
		wantErr bool
	}{
		{
			name: "regular case",
			c:    testController,
			args: args{
				input: "2,2,2",
			},
			want:    []float64{2, 2, 2},
			wantErr: false,
		},
		{
			name: "negative numbers",
			c:    testController,
			args: args{
				input: "-2,-2,-2",
			},
			want:    []float64{-2, -2, -2},
			wantErr: false,
		},
		{
			name: "negative numbers",
			c:    testController,
			args: args{
				input: "-2, -2, -2",
			},
			want:    []float64{-2, -2, -2},
			wantErr: false,
		},
		{
			name: "spaces",
			c:    testController,
			args: args{
				input: "2,        2,      2",
			},
			want:    []float64{2, 2, 2},
			wantErr: false,
		},
		{
			name: "float nums",
			c:    testController,
			args: args{
				input: "2.5, 2.1",
			},
			want:    []float64{2.5, 2.1},
			wantErr: false,
		},
		{
			name: "not nums",
			c:    testController,
			args: args{
				input: "abcd",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "empty string",
			c:    testController,
			args: args{
				input: "",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.prepareNums(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("calcController.prepareNums() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calcController.prepareNums() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calcController_Sum(t *testing.T) {
	testController := newController(logrus.New(), cache.New(), requests.New())

	type args struct {
		uID   string
		input models.Input
	}
	tests := []struct {
		name    string
		c       calcController
		args    args
		want    *models.CalcAction
		wantErr bool
	}{
		{
			name: "regular case",
			c:    testController,
			args: args{
				uID: "123",
				input: models.Input{
					Input: "1, 2",
				},
			},
			want: &models.CalcAction{
				Input:  "1, 2",
				Action: models.Sum,
				Result: 3,
			},
			wantErr: false,
		},
		{
			name: "negative numbers",
			c:    testController,
			args: args{
				uID: "123",
				input: models.Input{
					Input: "-1, -2",
				},
			},
			want: &models.CalcAction{
				Input:  "-1, -2",
				Action: models.Sum,
				Result: -3,
			},
			wantErr: false,
		},
		{
			name: "empty string",
			c:    testController,
			args: args{
				uID: "123",
				input: models.Input{
					Input: "",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "not numbers",
			c:    testController,
			args: args{
				uID: "123",
				input: models.Input{
					Input: "a, b, c",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Sum(tt.args.uID, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("calcController.Sum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calcController.Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calcController_Mult(t *testing.T) {
	testController := newController(logrus.New(), cache.New(), requests.New())

	type args struct {
		uID   string
		input models.Input
	}
	tests := []struct {
		name    string
		c       calcController
		args    args
		want    *models.CalcAction
		wantErr bool
	}{
		{
			name: "regular case",
			c:    testController,
			args: args{
				uID: "123",
				input: models.Input{
					Input: "1, 2",
				},
			},
			want: &models.CalcAction{
				Input:  "1, 2",
				Action: models.Mult,
				Result: 2,
			},
			wantErr: false,
		},
		{
			name: "float numbers",
			c:    testController,
			args: args{
				uID: "123",
				input: models.Input{
					Input: "1.5, 2",
				},
			},
			want: &models.CalcAction{
				Input:  "1.5, 2",
				Action: models.Mult,
				Result: 3,
			},
			wantErr: false,
		},
		{
			name: "negative numbers",
			c:    testController,
			args: args{
				uID: "123",
				input: models.Input{
					Input: "-1, 2",
				},
			},
			want: &models.CalcAction{
				Input:  "-1, 2",
				Action: models.Mult,
				Result: -2,
			},
			wantErr: false,
		},
		{
			name: "empty string",
			c:    testController,
			args: args{
				uID: "123",
				input: models.Input{
					Input: "",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "not numbers",
			c:    testController,
			args: args{
				uID: "123",
				input: models.Input{
					Input: "a, b, c",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Mult(tt.args.uID, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("calcController.Mult() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calcController.Mult() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calcController_Calculate(t *testing.T) {
	testController := newController(logrus.New(), cache.New(), requests.New())

	type args struct {
		uID   string
		input models.Input
	}
	tests := []struct {
		name    string
		c       calcController
		args    args
		want    *models.CalcAction
		wantErr bool
	}{
		{
			name: "simple",
			c:    testController,
			args: args{
				uID: "123",
				input: models.Input{
					Input: "2+2",
				},
			},
			want: &models.CalcAction{
				Input:  "2+2",
				Action: models.Calc,
				Result: 4,
			},
			wantErr: false,
		},
		{
			name: "simple spaces",
			c:    testController,
			args: args{
				uID: "123",
				input: models.Input{
					Input: "2 +   2",
				},
			},
			want: &models.CalcAction{
				Input:  "2 +   2",
				Action: models.Calc,
				Result: 4,
			},
			wantErr: false,
		},
		{
			name: "many acitons",
			c:    testController,
			args: args{
				uID: "123",
				input: models.Input{
					Input: "8*((4+3/3)/8)+15-2^4",
				},
			},
			want: &models.CalcAction{
				Input:  "8*((4+3/3)/8)+15-2^4",
				Action: models.Calc,
				Result: 4,
			},
			wantErr: false,
		},
		{
			name: "negative at start",
			c:    testController,
			args: args{
				uID: "123",
				input: models.Input{
					Input: "-8*4+3",
				},
			},
			want: &models.CalcAction{
				Input:  "-8*4+3",
				Action: models.Calc,
				Result: -29,
			},
			wantErr: false,
		},
		{
			name: "negative after action",
			c:    testController,
			args: args{
				uID: "123",
				input: models.Input{
					Input: "8*-4+3",
				},
			},
			want: &models.CalcAction{
				Input:  "8*-4+3",
				Action: models.Calc,
				Result: -29,
			},
			wantErr: false,
		},
		{
			name: "negative at brakes start",
			c:    testController,
			args: args{
				uID: "123",
				input: models.Input{
					Input: "8*(-4+3)",
				},
			},
			want: &models.CalcAction{
				Input:  "8*(-4+3)",
				Action: models.Calc,
				Result: -8,
			},
			wantErr: false,
		},
		{
			name: "bad brakes #1",
			c:    testController,
			args: args{
				uID: "123",
				input: models.Input{
					Input: "8*((4+3)",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "bad brakes #2",
			c:    testController,
			args: args{
				uID: "123",
				input: models.Input{
					Input: "8*4+3)",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "div by zero",
			c:    testController,
			args: args{
				uID: "123",
				input: models.Input{
					Input: "8/0",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "unknown action",
			c:    testController,
			args: args{
				uID: "123",
				input: models.Input{
					Input: "1b2",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "empty string",
			c:    testController,
			args: args{
				uID: "123",
				input: models.Input{
					Input: "",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "not numbers",
			c:    testController,
			args: args{
				uID: "123",
				input: models.Input{
					Input: "a, b, c",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Calculate(tt.args.uID, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("calcController.Calculate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calcController.Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calcController_History(t *testing.T) {
	type args struct {
		uID string
	}
	type earlierActions struct {
		uID    string
		input  models.Input
		action models.Action
	}
	tests := []struct {
		name    string
		c       calcController
		args    args
		actions []earlierActions
		want    []models.CalcAction
	}{
		{
			name: "empty history",
			c:    newController(logrus.New(), cache.New(), requests.New()),
			args: args{
				uID: "1",
			},
			actions: []earlierActions{},
			want:    []models.CalcAction{},
		},
		{
			name: "empty history (current user)",
			c:    newController(logrus.New(), cache.New(), requests.New()),
			args: args{
				uID: "1",
			},
			actions: []earlierActions{
				{
					uID:    "2",
					action: models.Calc,
					input: models.Input{
						Input: "2+2",
					},
				},
				{
					uID:    "3",
					action: models.Sum,
					input: models.Input{
						Input: "2, 2",
					},
				},
				{
					uID:    "3",
					action: models.Mult,
					input: models.Input{
						Input: "2, 2",
					},
				},
			},
			want: []models.CalcAction{},
		},
		{
			name: "regular history",
			c:    newController(logrus.New(), cache.New(), requests.New()),
			args: args{
				uID: "1",
			},
			actions: []earlierActions{
				{
					uID:    "1",
					action: models.Calc,
					input: models.Input{
						Input: "2+2",
					},
				},
				{
					uID:    "1",
					action: models.Sum,
					input: models.Input{
						Input: "2, 2",
					},
				},
				{
					uID:    "1",
					action: models.Mult,
					input: models.Input{
						Input: "2, 2",
					},
				},
			},
			want: []models.CalcAction{
				{
					Input:  "2+2",
					Action: models.Calc,
					Result: 4,
				},
				{
					Input:  "2, 2",
					Action: models.Sum,
					Result: 4,
				},
				{
					Input:  "2, 2",
					Action: models.Mult,
					Result: 4,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, action := range tt.actions {
				switch action.action {
				case models.Calc:
					tt.c.Calculate(action.uID, action.input)
				case models.Sum:
					tt.c.Sum(action.uID, action.input)
				case models.Mult:
					tt.c.Mult(action.uID, action.input)
				}
			}

			if got := tt.c.History(tt.args.uID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calcController.History() = %v, want %v", got, tt.want)
			}
		})
	}
}
