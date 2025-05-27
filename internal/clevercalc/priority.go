package clevercalc

import "github.com/LexusEgorov/api-calculator/internal/models"

type ranksMap map[string]int

type priority struct {
	ranks ranksMap
}

func newPriority() *priority {
	//Расставляем приоритеты действий
	priorityMap := make(ranksMap, 0)
	priorityMap["+"] = 1
	priorityMap["-"] = 1
	priorityMap["*"] = 2
	priorityMap["/"] = 2
	priorityMap["^"] = 3
	priorityMap["("] = 4
	priorityMap[")"] = 4

	return &priority{
		ranks: priorityMap,
	}
}

func (p priority) Get(operator string) int {
	rank, isFound := p.ranks[operator]

	if !isFound {
		return models.NotOpRank
	}

	return rank
}
