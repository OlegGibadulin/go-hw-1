package calc

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"

	"task_2/stack"
)

func Calculate(expr string) (float64, error) {
	values := stack.New()
	ops := stack.New()

	for i := 0; i < len(expr); i++ {
		token := rune(expr[i])

		if unicode.IsDigit(token) {
			// Parse number
			number := ""
			for ; i < len(expr) && (unicode.IsDigit(rune(expr[i])) || expr[i] == '.'); i++ {
				number += string(expr[i])
			}
			val, err := strconv.ParseFloat(number, 64)
			if err != nil {
				return 0, err
			}
			values.Push(val)

			if i == len(expr) {
				break
			}
			token = rune(expr[i])
		}

		if unicode.IsSpace(token) {
			continue
		}
		if token == '(' {
			ops.Push(token)
		} else if token == ')' {
			// Calculate inside brackets
			for !ops.Empty() && ops.Top() != '(' {
				res, err := performCalc(values, ops)
				if err != nil {
					return 0, err
				}
				values.Push(res)
			}
			_ = ops.Pop()
		} else {
			// Current token is operator
			if values.Empty() {
				return 0, errors.New("Wrong input data")
			}
			for !ops.Empty() && priority(ops.Top().(rune)) >= priority(token) {
				res, err := performCalc(values, ops)
				if err != nil {
					return 0, err
				}
				values.Push(res)
			}
			ops.Push(token)
		}
	}

	if values.Empty() && ops.Empty() {
		// Input string is empty or contains only spaces
		return 0, nil
	}

	if values.Len() == 1 {
		// Input string contains single number
		res := values.Pop().(float64)
		return res, nil
	}

	// Calculate the ramaining part without brackets
	for !ops.Empty() {
		res, err := performCalc(values, ops)
		if err != nil {
			return 0, err
		}
		values.Push(res)
	}
	res := values.Pop().(float64)
	return res, nil
}

func performCalc(values *stack.Stack, ops *stack.Stack) (float64, error) {
	secondVal := values.Pop()
	firstVal := values.Pop()
	operator := ops.Pop()
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
