package cache

import (
	"sync"

	"github.com/LexusEgorov/api-calculator/internal/models"
)

type actionsMap map[string]float64
type cacheMap map[models.Action]actionsMap

type Cache struct {
	cache cacheMap
	mu    sync.Mutex
}

func New() *Cache {
	cacheMap := make(cacheMap)
	cacheMap[models.Mult] = make(actionsMap)
	cacheMap[models.Sum] = make(actionsMap)
	cacheMap[models.Calc] = make(actionsMap)

	return &Cache{
		cache: cacheMap,
	}
}

func (c *Cache) Get(input string, action models.Action) (models.CalcAction, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	actionsMap, isFound := c.cache[action]

	if !isFound {
		return models.CalcAction{}, models.ErrCacheNotFound
	}

	res, isFound := actionsMap[input]

	if !isFound {
		return models.CalcAction{}, models.ErrCacheNotFound
	}

	return models.CalcAction{
		Input:  input,
		Action: action,
		Result: res,
	}, nil
}

func (c *Cache) Set(action models.CalcAction) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	actions, isFound := c.cache[action.Action]

	if !isFound {
		return models.NewCacheMapErr(string(action.Action))
	}

	actions[action.Input] = action.Result
	return nil
}
