package models

import (
	"errors"
	"fmt"
)

var CacheNotFoundErr = errors.New("not cached")

func NewCacheMapErr(key string) error {
	return errors.New(fmt.Sprintf("cache: can't set to map with key: %s", key))
}
