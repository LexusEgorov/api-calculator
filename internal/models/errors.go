package models

import (
	"errors"
	"fmt"
)

var ErrCacheNotFound = errors.New("not cached")

func NewCacheMapErr(key string) error {
	return fmt.Errorf("cache: can't set to map with key: %s", key)
}
