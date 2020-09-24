package calc

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"

	"task_2/stack"
)

func Calculate(expr string) (float64, error) {
	numbers := stack.New()
	operators := stack.New()

	for i := 0; i < len(expr); i++ {
		token := rune(expr[i])

		if unicode.IsDigit(token) {
			// Parse number
			val := ""
			for ; i < len(expr) && (unicode.IsDigit(rune(expr[i])) || expr[i] == '.'); i++ {
				val += string(expr[i])
			}
			number, err := strconv.ParseFloat(val, 64)
			if err != nil {
				return 0, err
			}
			numbers.Push(number)

			if i == len(expr) {
				break
			}
			token = rune(expr[i])
		}

		if unicode.IsSpace(token) {
			continue
		}
		if token == '(' {
			operators.Push(token)
		} else if token == ')' {
			// Calculate inside brackets
			for !operators.Empty() && operators.Top() != '(' {
				res, err := performCalculation(numbers, operators)
				if err != nil {
					return 0, err
				}
				numbers.Push(res)
			}
			_ = operators.Pop()
		} else {
			// Current token is operator
			if numbers.Empty() {
				return 0, errors.New("Wrong input data")
			}
			for !operators.Empty() && priority(operators.Top().(rune)) >= priority(token) {
				res, err := performCalculation(numbers, operators)
				if err != nil {
					return 0, err
				}
				numbers.Push(res)
			}
			operators.Push(token)
		}
	}

	if operators.Empty() {
		if numbers.Empty() {
			// Input string is empty or contains only spaces
			return 0, nil
		} else if numbers.Len() == 1 {
			// Input string contains only single number
			res := numbers.Pop().(float64)
			return res, nil
		}
	}

	if numbers.Len() != operators.Len()+1 {
		// There are extra numbers
		return 0, errors.New("Wrong input data")
	}

	// Calculate the ramaining part without brackets
	for !operators.Empty() {
		res, err := performCalculation(numbers, operators)
		if err != nil {
			return 0, err
		}
		numbers.Push(res)
	}
	res := numbers.Pop().(float64)
	return res, nil
}

func performCalculation(numbers *stack.Stack, operators *stack.Stack) (float64, error) {
	secondVal := numbers.Pop()
	firstVal := numbers.Pop()
	operator := operators.Pop()
	if secondVal == nil || firstVal == nil || operator == nil {
		return 0, errors.New("Wrong input data")
	}
	res, err := calcOperation(firstVal.(float64), secondVal.(float64), operator.(rune))
	if err != nil {
		return 0, err
	}
	return res, err
}

func calcOperation(firstVal float64, secondVal float64, operator rune) (res float64, err error) {
	err = nil
	switch operator {
	case '+':
		res = firstVal + secondVal
	case '-':
		res = firstVal - secondVal
	case '*':
		res = firstVal * secondVal
	case '/':
		res = firstVal / secondVal
	default:
		err = fmt.Errorf("Unknown operator %#U", operator)
	}
	return
}

func priority(operator rune) int {
	if operator == '+' || operator == '-' {
		return 1
	}
	if operator == '*' || operator == '/' {
		return 2
	}
	return 0
}
