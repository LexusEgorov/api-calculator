package models

import (
	"errors"
	"fmt"
)

var ErrCacheNotFound = errors.New("not cached")
var ErrBadInput = errors.New("bad input")
var ErrBadInternalInput = errors.New("bad internal input")
var ErrBadDivide = errors.New("can't divide by zero")

func NewErrUnknownAction(action string) error {
	return fmt.Errorf("action '%s' doesn't supported", action)
}

func NewCacheMapErr(key string) error {
	return fmt.Errorf("cache: can't set to map with key: %s", key)
}
