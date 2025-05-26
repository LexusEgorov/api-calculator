package clevercalc

import (
	"fmt"
	"strings"

	"github.com/LexusEgorov/api-calculator/internal/models"
)

// слайс для разбивки входной строки для обработки кейсов X-N (X операция, N число)
var operationsSplit = []string{
	models.OPENING_BRAKE,
	models.OPERATION_DIV,
	models.OPERATION_SUM,
	models.OPERATION_MULT,
	models.OPERATION_POW,
}

type parser struct {
	priorityMap *priority
}

func newParser() *parser {
	return &parser{
		priorityMap: newPriority(),
	}
}

// Парсит входные данные в слайс постфиксной нотации для дальнейших вычислений
func (p parser) parse(input string) ([]string, error) {
	input = p.prepare(input)
	actionsStack := Stack{}

	var result []string
	//Воспомогательная переменная для обработки кейсов многозначных (в том числе дробных) чисел
	var current string

	for _, symb := range input {
		curr := string(symb)
		if curr == models.OPENING_BRAKE {
			actionsStack.Push(curr)
			continue
		}

		//Достаем все действия до открывающей скобки
		if curr == models.CLOSING_BRAKE {
			addNum(&result, &current)
			for {
				action, err := actionsStack.Pop()

				if err != nil {
					return nil, models.ErrBadInput
				}

				if action == models.OPENING_BRAKE {
					break
				}

				result = append(result, string(action))
			}

			continue
		}

		rank := p.priorityMap.Get(curr)

		//Если true, то текущий символ точно не действие
		if rank == models.NOT_OP_RANK {
			current += string(curr)
			continue
		}

		//Сохраняем число, дальше работа с действиями
		addNum(&result, &current)

		//Извлекаем по правилу приоритетов
		actions := p.getActions(&actionsStack, rank)

		//И добавляем в результат
		for _, action := range actions {
			result = append(result, action)
		}

		//Текущее действие в стек
		actionsStack.Push(curr)
	}

	addNum(&result, &current)

	//Добавляем оставшиеся действия в результат
	for {
		action, err := actionsStack.Pop()

		if action == models.OPENING_BRAKE {
			return nil, models.ErrBadInput
		}

		if err != nil {
			break
		}

		result = append(result, string(action))
	}

	return result, nil
}

// Преобразует все отрицательные числа в начале строки/скобок/после оператора в формат (0-X)
func (p parser) prepare(input string) string {
	//Убираем пробелы
	input = removeSpaces(input)

	//Подготавливаем отрицательные числа
	for _, separator := range operationsSplit {
		tmp := strings.Split(input, separator)
		for i, part := range tmp {
			tmp[i] = p.prepareSimplePart(part)
		}

		input = strings.Join(tmp, separator)
	}

	return input
}

// Подготовка частей, после разбития строки по разделителю
func (p parser) prepareSimplePart(part string) string {
	if part == "" {
		return part
	}

	if string(part[0]) != models.OPERATION_SUB {
		return part
	}

	//Может встретиться, только если исходная строка начинается на -(
	if len(part) < 2 {
		return "0-"
	}

	if string(part[1]) == models.OPENING_BRAKE {
		return fmt.Sprintf("(0%s)", part)
	}

	//Набираем первое отрицательное число и меняем его на (0-X)
	current := "-"
	for i := 1; i < len(part); i++ {
		stringedSymb := string(part[i])

		if p.priorityMap.Get(stringedSymb) == models.NOT_OP_RANK {
			current += stringedSymb
			continue
		}

		part = strings.Replace(part, current, fmt.Sprintf("(0%s)", current), 1)
		return part
	}

	//Если текущая часть просто отрицательное число
	part = strings.Replace(part, current, fmt.Sprintf("(0%s)", current), 1)
	return part
}

// Извлечение до тех пор, пока следующее действие не будет по приоритету ниже текущего
func (p parser) getActions(from *Stack, rank int) []string {
	actions := make([]string, 0)

	for {
		nextAction, err := from.Peek()

		if err != nil || nextAction == models.OPENING_BRAKE || p.priorityMap.Get(nextAction) < rank {
			return actions
		}

		from.Pop()
		actions = append(actions, nextAction)
	}
}

// добавление числа в результирующий слайс
func addNum(destination *[]string, num *string) {
	resNum := removeSpaces(*num)
	if resNum == "" {
		return
	}

	*destination = append(*destination, resNum)
	*num = ""
}

// удаляет пробелы
func removeSpaces(str string) string {
	return strings.ReplaceAll(str, " ", "")
}
