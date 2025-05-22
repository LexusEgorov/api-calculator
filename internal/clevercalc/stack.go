package clevercalc

import (
	"errors"
)

var ErrEmpty = errors.New("empty stack")

type Stack struct {
	data []string
	size int
}

func (a *Stack) Push(data string) {
	a.data = append(a.data, data)
	a.size++
}

func (a *Stack) Pop() (string, error) {
	if a.size == 0 {
		return "", ErrEmpty
	}

	a.size--
	res := a.data[a.size]

	a.data = a.data[:a.size]
	return res, nil
}

func (a Stack) Peek() (string, error) {
	if a.size == 0 {
		return "", ErrEmpty
	}

	return a.data[a.size-1], nil
}

func (a Stack) Size() int {
	return a.size
}
