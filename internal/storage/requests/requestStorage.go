package requests

import (
	"sync"

	"github.com/LexusEgorov/api-calculator/internal/models"
)

type RequestStorage struct {
	requests map[string][]models.CalcAction
	mu       sync.Mutex
}

func New() *RequestStorage {
	return &RequestStorage{
		requests: make(map[string][]models.CalcAction),
	}
}

func (r *RequestStorage) Get(uID string) []models.CalcAction {
	r.mu.Lock()
	requests := r.requests[uID]
	r.mu.Unlock()

	if requests == nil {
		return make([]models.CalcAction, 0)
	}

	return requests
}

func (r *RequestStorage) Set(uID string, action models.CalcAction) {
	r.mu.Lock()
	requests := append(r.requests[uID], action)
	r.requests[uID] = requests
	r.mu.Unlock()
}
