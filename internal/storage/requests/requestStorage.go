package requests

import (
	"api-calculator/internal/models"
	"sync"
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

	return requests
}

func (r *RequestStorage) Set(uID string, action models.CalcAction) {
	r.mu.Lock()
	requests := append(r.requests[uID], action)
	r.requests[uID] = requests
	r.mu.Unlock()
}
