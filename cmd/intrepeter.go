package main

import (
	"fmt"
	"strconv"
)

type interpreter struct {
	glox *Glox
}

func (i *interpreter) interpret(expr Expr) (any, error) {
	val, err := i.evaluate(expr)
	if err != nil || i.glox.hadError {
		return nil, err
	} else {
		fmt.Printf("%v", val)
	}
	return val, err
}

func (i *interpreter) visitLiteral(l *Literal) (any, error) {
	return l.value, nil
}

func (i *interpreter) visitGrouping(g *Grouping) (any, error) {
	return i.evaluate(g.groupedExpression)
}

func (i *interpreter) evaluate(e Expr) (any, error) {
	return e.accept(i)
}

func (i *interpreter) visitUnary(u *Unary) (any, error) {
	expr, err := i.evaluate(u.right)
	if err != nil {
		return nil, err
	}

	switch u.operator.ttype {
	case MINUS:
		if err := i.checkNumberOperand(u.operator, expr); err != nil {
			return nil, err
		}

		return -expr.(float64), nil

	case BANG:
		return !i.isTruthy(expr), nil
	}

	return nil, nil
}

func (i *interpreter) visitBinary(b *Binary) (any, error) {
	left, err := i.evaluate(b.left)
	if err != nil {
		return nil, err
	}

	right, err := i.evaluate(b.right)
	if err != nil {
		return nil, err
	}

	switch b.operator.ttype {

	case MINUS:
		if err := i.checkNumberOperands(b.operator, left, right); err != nil {
			return nil, err
		}
		return left.(float64) - right.(float64), nil

	case SLASH:
		if err := i.checkNumberOperands(b.operator, left, right); err != nil {
			return nil, err
		}
		if right.(float64) == 0 {
			err := RuntimeError{b.operator, "MathmaticalHorror Division by zero"}
			return nil, err
		}
		return left.(float64) / right.(float64), nil

	case STAR:
		if err := i.checkNumberOperands(b.operator, left, right); err != nil {
			return nil, err
		}
		return left.(float64) * right.(float64), nil

	case PLUS:
		if IsInstanceOf[float64](left) && IsInstanceOf[float64](right) {
			return left.(float64) + right.(float64), nil
		}

		if IsInstanceOf[string](left) && IsInstanceOf[string](right) {
			return left.(string) + right.(string), nil
		}

		if IsInstanceOf[float64](left) {
			return strconv.FormatFloat(left.(float64), 'f', -1, 64) + right.(string), nil
		} else {
			return left.(string) + strconv.FormatFloat(right.(float64), 'f', -1, 64), nil
		}
	case GREATER:
		if err := i.checkNumberOperands(b.operator, left, right); err != nil {
			return nil, err
		}
		return left.(float64) > right.(float64), nil

	case GREATER_EQUAL:
		if err := i.checkNumberOperands(b.operator, left, right); err != nil {
			return nil, err
		}
		return left.(float64) >= right.(float64), nil

	case LESS:
		if err := i.checkNumberOperands(b.operator, left, right); err != nil {
			return nil, err
		}
		return left.(float64) < right.(float64), nil

	case LESS_EQUAL:
		if err := i.checkNumberOperands(b.operator, left, right); err != nil {
			return nil, err
		}
		return left.(float64) <= right.(float64), nil

	case BANG_EQUAL:
		return !i.isEqual(left, right), nil

	case EQUAL_EQUAL:
		return i.isEqual(left, right), nil
	}

	return nil, nil
}

func (i *interpreter) isTruthy(object any) bool {
	if object == nil {
		return false
	}

	if boolean, ok := object.(bool); ok {
		return boolean
	}

	return true
}

func (i *interpreter) isEqual(a any, b any) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil {
		return false
	}

	return a == b
}

func (i *interpreter) checkNumberOperand(operator Token, operand any) error {
	if IsInstanceOf[float64](operand) {
		return nil
	}

	return RuntimeError{
		operator,
		fmt.Sprintf(
			"Operand for operator {%v} must be a number",
			operator.lexeme,
		),
	}
}

func (i *interpreter) checkNumberOperands(operator Token, a any, b any) error {
	if IsInstanceOf[float64](a) &&
		IsInstanceOf[float64](b) {
		return nil
	}

	return RuntimeError{
		operator,
		fmt.Sprintf(
			"Operands for operator {%v} must be numbers",
			operator.lexeme,
		),
	}
}

func IsInstanceOf[T any](obj any) bool {
	_, ok := obj.(T)
	return ok
}
