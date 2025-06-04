package clevercalc

import (
	"testing"

	"github.com/LexusEgorov/api-calculator/internal/models"
)

func TestNew(t *testing.T) {
	testPriority := newPriority()

	tests := []struct {
		input    string
		expected int
	}{
		{
			input:    "+",
			expected: 1,
		},
		{
			input:    "-",
			expected: 1,
		},
		{
			input:    "*",
			expected: 2,
		},
		{
			input:    "/",
			expected: 2,
		},
		{
			input:    "^",
			expected: 3,
		},
		{
			input:    "(",
			expected: 4,
		},
		{
			input:    ")",
			expected: 4,
		},
	}

	for _, test := range tests {
		res := testPriority.ranks[test.input]

		if res != test.expected {
			t.Errorf("priority.New: field: '%s': expected: %d, got: %d", test.input, test.expected, res)
		}
	}
}

func TestGet(t *testing.T) {
	testPriority := newPriority()

	tests := []struct {
		input    string
		expected int
	}{
		{
			input:    "+",
			expected: 1,
		},
		{
			input:    "-",
			expected: 1,
		},
		{
			input:    "*",
			expected: 2,
		},
		{
			input:    "/",
			expected: 2,
		},
		{
			input:    "^",
			expected: 3,
		},
		{
			input:    "(",
			expected: 4,
		},
		{
			input:    ")",
			expected: 4,
		},
		{
			input:    "",
			expected: models.NotOpRank,
		},
		{
			input:    "5",
			expected: models.NotOpRank,
		},
	}

	for _, test := range tests {
		res := testPriority.Get(test.input)

		if res != test.expected {
			t.Errorf("priority.Get(%s): expected: %d, got: %d", test.input, test.expected, res)
		}
	}
}
