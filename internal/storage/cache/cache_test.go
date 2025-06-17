package cache

import (
	"reflect"
	"sync"
	"testing"

	"github.com/LexusEgorov/api-calculator/internal/models"
)

func TestNew(t *testing.T) {
	etalonMap := cacheMap{}
	etalonMap[models.Mult] = make(actionsMap)
	etalonMap[models.Sum] = make(actionsMap)
	etalonMap[models.Calc] = make(actionsMap)

	etalonCache := Cache{
		cache: etalonMap,
		mu:    sync.Mutex{},
	}

	tests := []struct {
		name string
		want *Cache
	}{
		{
			name: "init all maps",
			want: &etalonCache,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCache_Get(t *testing.T) {
	testCache := New()

	type args struct {
		input  string
		action models.Action
	}
	tests := []struct {
		name    string
		c       *Cache
		args    args
		init    cacheMap
		want    models.CalcAction
		wantErr bool
	}{
		{
			name: "get cached action (CALC)",
			c:    testCache,
			args: args{
				input:  "2+2",
				action: models.Calc,
			},
			init: map[models.Action]actionsMap{
				models.Calc: map[string]float64{
					"2+2": 4,
				},
			},
			want: models.CalcAction{
				Input:  "2+2",
				Action: models.Calc,
				Result: 4,
			},
			wantErr: false,
		},
		{
			name: "get cached action (Mult)",
			c:    testCache,
			args: args{
				input:  "2, 2",
				action: models.Mult,
			},
			init: map[models.Action]actionsMap{
				models.Mult: map[string]float64{
					"2, 2": 4,
				},
			},
			want: models.CalcAction{
				Input:  "2, 2",
				Action: models.Mult,
				Result: 4,
			},
			wantErr: false,
		},
		{
			name: "get cached action (Sum)",
			c:    testCache,
			args: args{
				input:  "2, 2",
				action: models.Sum,
			},
			init: map[models.Action]actionsMap{
				models.Sum: map[string]float64{
					"2, 2": 4,
				},
			},
			want: models.CalcAction{
				Input:  "2, 2",
				Action: models.Sum,
				Result: 4,
			},
			wantErr: false,
		},
		{
			name: "get uncached action (Sum)",
			c:    testCache,
			args: args{
				input:  "2, 2",
				action: models.Sum,
			},
			init:    map[models.Action]actionsMap{},
			want:    models.CalcAction{},
			wantErr: true,
		},
		{
			name: "get uncached action (Mult)",
			c:    testCache,
			args: args{
				input:  "2, 2",
				action: models.Mult,
			},
			init:    map[models.Action]actionsMap{},
			want:    models.CalcAction{},
			wantErr: true,
		},
		{
			name: "get uncached action (Calc)",
			c:    testCache,
			args: args{
				input:  "2+2",
				action: models.Calc,
			},
			init:    map[models.Action]actionsMap{},
			want:    models.CalcAction{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.cache = tt.init
			got, err := tt.c.Get(tt.args.input, tt.args.action)
			if (err != nil) != tt.wantErr {
				t.Errorf("Cache.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cache.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCache_Set(t *testing.T) {
	testCache := New()

	type args struct {
		action models.CalcAction
	}
	tests := []struct {
		name    string
		c       *Cache
		args    args
		init    cacheMap
		want    cacheMap
		wantErr bool
	}{
		{
			name: "set known action",
			c:    testCache,
			args: args{
				action: models.CalcAction{
					Input:  "2, 2",
					Action: models.Sum,
					Result: 4,
				},
			},
			init: map[models.Action]actionsMap{
				models.Sum: map[string]float64{},
			},
			want: map[models.Action]actionsMap{
				models.Sum: map[string]float64{
					"2, 2": 4,
				},
			},
			wantErr: false,
		},
		{
			name: "set unknown action",
			c:    testCache,
			args: args{
				action: models.CalcAction{
					Input:  "2, 2",
					Action: models.Sum,
					Result: 4,
				},
			},
			init:    map[models.Action]actionsMap{},
			want:    map[models.Action]actionsMap{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.cache = tt.init

			if err := tt.c.Set(tt.args.action); (err != nil) != tt.wantErr {
				t.Errorf("Cache.Set(%v) error = %v, wantErr %v", tt.args.action, err, tt.wantErr)
			}

			if !reflect.DeepEqual(tt.c.cache, tt.want) {
				t.Errorf("Cache.Set(%v) want = %v got = %v", tt.args.action, tt.want, tt.c.cache)
			}
		})
	}
}
