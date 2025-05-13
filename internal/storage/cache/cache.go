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

// Get implements calculator.Cacher.
func (c *Cache) Get(input string, action models.Action) (float64, error) {
	c.mu.Lock()
	res, isFound := c.cache[action][input]
	c.mu.Unlock()

	if !isFound {
		return 0, models.CacheNotFoundErr
	}

	return res, nil
}

// Save implements calculator.Cacher.
func (c *Cache) Save(action models.CalcAction) error {
	c.mu.Lock()
	actionMap := c.cache[action.Action]
	actionMap[action.Input] = action.Result
	c.cache[action.Action] = actionMap
	c.mu.Unlock()

	return nil
}

func New() *Cache {
	cacheMap := make(cacheMap)
	cacheMap[models.MULT] = make(actionsMap)
	cacheMap[models.SUM] = make(actionsMap)

	return &Cache{
		cache: cacheMap,
	}
}
