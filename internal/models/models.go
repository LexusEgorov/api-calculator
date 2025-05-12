package models

import "errors"

type CalcAction struct {
	Input  string
	Action Action
	Result float64
}

type UserRequest struct {
	uID string
	CalcAction
}

type Action int

const (
	SUM Action = iota
	MULT
)

var CacheNotFoundErr = errors.New("not cached")
