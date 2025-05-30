package cache

import (
	"testing"

	"github.com/LexusEgorov/api-calculator/internal/models"
)

func TestSetGet(t *testing.T) {
	cache := New()

	input := "1, 1"
	action := models.CalcAction{
		Input:  input,
		Action: models.Sum,
		Result: 2,
	}

	cache.Set(action)

	result, err := cache.Get(input, models.Sum)

	if err != nil {
		t.Errorf("Got error: %v", err)
	}

	if result != action {
		t.Errorf("Not equal!\nGot: %v\nExpected: %v\n", result, action)
	}
}
