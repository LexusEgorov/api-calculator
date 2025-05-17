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
	actionsMap, isFound := c.cache[action]

	if !isFound {
		return nil, models.CacheNotFoundErr
	}

	res, isFound := actionsMap[input]
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
	actionsMap, isFound := c.cache[action.Action]

	if !isFound {
		c.cache[action.Action] = make(map[string]float64)
		c.Set(action)
		return
	}

	actionsMap[action.Input] = action.Result
	c.cache[action.Action] = actionsMap
	c.mu.Unlock()
}
