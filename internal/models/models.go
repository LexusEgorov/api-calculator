package models

type CalcAction struct {
	Input  string  `json:"input"`
	Action Action  `json:"action"`
	Result float64 `json:"result"`
}

type Action int

const (
	SUM Action = iota
	MULT
)
