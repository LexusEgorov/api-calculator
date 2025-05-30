package models

import (
	"errors"
	"fmt"
)

var (
	ErrBadConfigPort    = errors.New("port must be upper than 0")
	ErrCacheNotFound    = errors.New("not cached")
	ErrBadInput         = errors.New("bad input")
	ErrBadInternalInput = errors.New("bad internal input")
	ErrBadDivide        = errors.New("can't divide by zero")
)

func NewErrUnknownAction(action string) error {
	return fmt.Errorf("action '%s' doesn't supported", action)
}

func NewCacheMapErr(key string) error {
	return fmt.Errorf("cache: can't set to map with key: %s", key)
}
