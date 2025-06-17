package requests

import (
	"reflect"
	"sync"
	"testing"

	"github.com/LexusEgorov/api-calculator/internal/models"
)

func TestNew(t *testing.T) {
	etalonStorage := RequestStorage{
		requests: make(map[string][]models.CalcAction),
		mu:       sync.Mutex{},
	}
	tests := []struct {
		name string
		want *RequestStorage
	}{
		{
			name: "init storage",
			want: &etalonStorage,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequestStorage_Set(t *testing.T) {
	testStorage := New()

	type args struct {
		uID    string
		action models.CalcAction
	}
	tests := []struct {
		name string
		r    *RequestStorage
		init map[string][]models.CalcAction
		args args
		want map[string][]models.CalcAction
	}{
		{
			name: "set action",
			r:    testStorage,
			init: make(map[string][]models.CalcAction),
			args: args{
				uID: "123",
				action: models.CalcAction{
					Input:  "2+2",
					Action: models.Calc,
					Result: 4,
				},
			},
			want: map[string][]models.CalcAction{
				"123": {
					{
						Input:  "2+2",
						Action: models.Calc,
						Result: 4,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.r.Set(tt.args.uID, tt.args.action)
		})
	}
}

func TestRequestStorage_Get(t *testing.T) {
	testStorage := New()

	type args struct {
		uID string
	}
	tests := []struct {
		name string
		r    *RequestStorage
		init map[string][]models.CalcAction
		args args
		want []models.CalcAction
	}{
		{
			name: "get actions",
			r:    testStorage,
			args: args{
				uID: "123",
			},
			init: map[string][]models.CalcAction{
				"123": {
					{
						Input:  "2+2",
						Action: models.Calc,
						Result: 4,
					},
				},
			},
			want: []models.CalcAction{
				{
					Input:  "2+2",
					Action: models.Calc,
					Result: 4,
				},
			},
		},
		{
			name: "get actions (empty)",
			r:    testStorage,
			args: args{
				uID: "123",
			},
			init: map[string][]models.CalcAction{},
			want: []models.CalcAction{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.r.requests = tt.init
			if got := tt.r.Get(tt.args.uID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RequestStorage.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
