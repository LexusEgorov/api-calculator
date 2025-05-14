package cache

import (
	"api-calculator/internal/models"
	"sync"
)

type actionsMap map[string]float64
type cacheMap map[models.Action]actionsMap

type Cache struct {
	cache cacheMap
	mu    sync.Mutex
}

func New() *Cache {
	cacheMap := make(cacheMap)
	cacheMap[models.MULT] = make(actionsMap)
	cacheMap[models.SUM] = make(actionsMap)

	return &Cache{
		cache: cacheMap,
	}
}

func (c *Cache) Get(input string, action models.Action) (*models.CalcAction, error) {
	c.mu.Lock()
	res, isFound := c.cache[action][input]
	c.mu.Unlock()

	if !isFound {
		return nil, models.CacheNotFoundErr
	}

	return &models.CalcAction{
		Input:  input,
		Action: action,
		Result: res,
	}, nil
}

func (c *Cache) Set(action models.CalcAction) {
	c.mu.Lock()
	actionMap := c.cache[action.Action]
	actionMap[action.Input] = action.Result
	c.cache[action.Action] = actionMap
	c.mu.Unlock()
}
