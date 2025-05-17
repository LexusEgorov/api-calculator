package models

type Input struct {
	Input string `json:"input"`
}

type CalcAction struct {
	Input  string  `json:"input"`
	Action Action  `json:"action"`
	Result float64 `json:"result"`
}

type Action string

const (
	SUM  Action = "SUM"
	MULT        = "MULT"
)
