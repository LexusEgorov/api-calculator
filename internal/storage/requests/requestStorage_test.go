package requests

import (
	"testing"

	"github.com/LexusEgorov/api-calculator/internal/models"
)

func TestSetGet(t *testing.T) {
	storage := New()

	input := "1, 1"
	action := models.CalcAction{
		Input:  input,
		Action: models.SUM,
		Result: 2,
	}

	storage.Set(input, action)

	result := storage.Get(input)

	if len(result) != 1 {
		t.Errorf("len not equa! got: %d expected: %d\n", len(result), 1)
	}

	if result[0] != action {
		t.Errorf("Not equal!\nGot: %v\nExpected: %v\n", result[0], action)
	}
}
