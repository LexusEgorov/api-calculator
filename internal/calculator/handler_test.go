package calculator

import (
	"testing"

	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

var handler = &CalcHandler{
	logger:     logger,
	controller: controller,
}

func TestHandleSum(t *testing.T) {

}
