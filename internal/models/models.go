package models

type Input struct {
	Input string `json:"input"`
}

type CalcAction struct {
	Input  string  `json:"input"`
	Action Action  `json:"action"`
	Result float64 `json:"result"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type Action string

const (
	MULT           Action = "MULT"
	SUM            Action = "SUM"
	CALC           Action = "CALC"
	NOT_OP_RANK           = -1
	OPENING_BRAKE         = "("
	CLOSING_BRAKE         = ")"
	OPERATION_SUM         = "+"
	OPERATION_SUB         = "-"
	OPERATION_MULT        = "*"
	OPERATION_DIV         = "/"
	OPERATION_POW         = "^"
)
