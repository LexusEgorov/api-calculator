package calculator

import (
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

type Cacher interface {
	Save(input string, result float64) error
	Get(input string) (float64, error)
}

type Storager interface {
	Save(uID, input string) error
	Get(uID string) []string
}

type CalcController struct {
	calc    calculator
	cache   Cacher
	storage Storager
	logger  *logrus.Logger
}

func New(logger *logrus.Logger, cache Cacher, storage Storager) *CalcController {
	return &CalcController{
		calc:    calculator{},
		cache:   cache,
		storage: storage,
		logger:  logger,
	}
}

func (c CalcController) HandleSum(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		c.logger.Errorf("calcController.HandleSum: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()

	c.logger.Info(string(body))
}

func (c CalcController) HandleMult(w http.ResponseWriter, r *http.Request)    {}
func (c CalcController) HandleHistory(w http.ResponseWriter, r *http.Request) {}
