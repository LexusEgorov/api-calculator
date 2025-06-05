package clevercalc

import (
	"fmt"
	"testing"
)

var testsCount = 10

func TestPush(t *testing.T) {
	stack := Stack{}

	for i := range testsCount {
		input := fmt.Sprint(i)
		stack.Push(input)

		if len(stack.data) != i+1 {
			t.Fatalf("stack size: %d, but expected %d", len(stack.data), i+1)
		}

		if stack.data[len(stack.data)-1] != input {
			t.Errorf("not equal: got: '%v', expected: '%v'", stack.data[len(stack.data)-1], input)
		}
	}
}

func TestPeek(t *testing.T) {
	stack := Stack{}

	for i := range testsCount {
		input := fmt.Sprint(i)
		stack.Push(input)

		size := len(stack.data)
		peek, err := stack.Peek()

		if err != nil {
			t.Fatalf("got error: %v", err)
		}

		if len(stack.data) != size {
			t.Fatalf("stack size changed after peek")
		}

		if peek != input {
			t.Errorf("not equal: got: '%v', expected: '%v'", peek, input)
		}
	}
}

func TestPop_Normal(t *testing.T) {
	stack := Stack{}

	for i := range testsCount {
		stack.Push(fmt.Sprint(i))
	}

	for range testsCount {
		size := len(stack.data)
		expected := stack.data[size-1]
		res, err := stack.Pop()

		if err != nil {
			t.Fatalf("got error: %v", err)
		}

		if len(stack.data) == size {
			t.Fatal("stack size doesn't change")
		}

		if expected != res {
			t.Errorf("not equal: got: '%v', expected: '%v'", res, expected)
		}
	}
}

func TestPop_Empty(t *testing.T) {
	stack := Stack{}

	_, err := stack.Pop()

	if err == nil {
		t.Fatal("expected error, but got nothing")
	}
}

func TestSize(t *testing.T) {
	stack := Stack{}

	for i := range testsCount {
		res := stack.Size()

		if res != len(stack.data) {
			t.Errorf("stack size: %d, but expected %d", res, len(stack.data))
		}

		stack.Push(fmt.Sprint(i))
	}
}
