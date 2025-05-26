package clevercalc

import (
	"reflect"
	"testing"

	"github.com/LexusEgorov/api-calculator/internal/models"
)

var parserTest = newParser()

func TestParse(t *testing.T) {
	input := "2+2"
	expected := []string{
		"2",
		"2",
		"+",
	}

	res, err := parserTest.parse(input)

	if err != nil {
		t.Errorf("Got error: %v", err)
	}

	if !reflect.DeepEqual(res, expected) {
		t.Errorf("Not equal!\nGot: %v\nExpected: %v\n", res, expected)
	}
}

func TestParseErr(t *testing.T) {
	input := "(2+2))"

	_, err := parserTest.parse(input)

	if err == nil {
		t.Errorf("Expected error after bad input: '%s'", input)
	}
}

func TestPrepare(t *testing.T) {
	input := "-5*(-3+2)"
	expected := "(0-5)*((0-3)+2)"

	res := parserTest.prepare(input)

	if res != expected {
		t.Errorf("Not equal!\nGot: %v\nExpected: %v\n", res, expected)
	}
}

func TestPrepareSimplePart(t *testing.T) {
	input := "-5*2"
	expected := "(0-5)*2"

	res := parserTest.prepareSimplePart(input)

	if res != expected {
		t.Errorf("Not equal!\nGot: %v\nExpected: %v\n", res, expected)
	}
}

func TestRemoveSpaces(t *testing.T) {
	input := "-5 * 2"
	expected := "-5*2"

	res := removeSpaces(input)

	if res != expected {
		t.Errorf("Not equal!\nGot: %v\nExpected: %v\n", res, expected)
	}
}

func TestAddNum(t *testing.T) {
	input := "10"
	destination := make([]string, 0)

	addNum(&destination, &input)

	if len(destination) == 0 {
		t.Errorf("destination slice is empty")
	}

	if destination[0] != "10" {
		t.Errorf("Not equal!\nGot: %v\nExpected: %v\n", destination[0], "10")
	}

	if input != "" {
		t.Errorf("input num isn't empty: got '%v'", input)
	}
}

func TestGetActions(t *testing.T) {
	mockStack := Stack{}

	mockStack.Push(models.OPERATION_SUM)
	mockStack.Push(models.OPERATION_POW)
	mockStack.Push(models.OPERATION_MULT)
	mockStack.Push(models.OPERATION_DIV)

	expected := []string{
		models.OPERATION_DIV,
		models.OPERATION_MULT,
		models.OPERATION_POW,
	}

	res := parserTest.getActions(&mockStack, 2)

	if !reflect.DeepEqual(res, expected) {
		t.Errorf("Not equal!\nGot: %v\nExpected: %v\n", res, expected)
	}
}
