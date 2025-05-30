package clevercalc

import "testing"

func TestPush(t *testing.T) {
	stack := Stack{}

	input := "abc"
	stack.Push(input)

	if len(stack.data) == 0 {
		t.Error("stack is empty")
	}

	if stack.data[0] != input {
		t.Errorf("not equal: got '%v', expected '%v'", stack.data[0], input)
	}
}

func TestPeek(t *testing.T) {
	stack := Stack{}

	input := "abc"
	stack.Push(input)
	res, err := stack.Peek()

	if err != nil {
		t.Errorf("got error: %v", err)
	}

	if res != input {
		t.Errorf("not equal: got '%v', expected '%v'", res, input)
	}
}

func TestPop(t *testing.T) {
	stack := Stack{}

	input := "abc"
	stack.Push(input)
	res, err := stack.Pop()

	if err != nil {
		t.Errorf("got error: %v", err)
	}

	if res != input {
		t.Errorf("not equal: got '%v', expected '%v'", res, input)
	}
}

func TestSize(t *testing.T) {
	stack := Stack{}

	input := "abc"
	res := stack.Size()

	if res != len(stack.data) {
		t.Errorf("not equal: got '%v', expected '%v'", res, input)
	}
}
