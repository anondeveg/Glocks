package main

import (
	"fmt"
	"strings"
)

type AstPrinter struct{}

func (a *AstPrinter) log(expr Expr) (any, error) {
	return expr.accept(a)
}

func (a *AstPrinter) parenthesize(name string, exprs ...Expr) (string, error) {
	var builder strings.Builder

	builder.WriteString("( " + name)

	for _, expr := range exprs {
		value, err := expr.accept(a)
		if err != nil {
			return "", err
		}

		builder.WriteString(" ")
		builder.WriteString(fmt.Sprintf("%v", value))
	}

	builder.WriteString(")")

	return builder.String(), nil
}

func (a *AstPrinter) visitBinary(b *Binary) (any, error) {
	return a.parenthesize(
		b.operator.lexeme,
		b.left,
		b.right,
	)
}

func (a *AstPrinter) visitLiteral(l *Literal) (any, error) {
	if l.value == nil {
		return "nil", nil
	}

	return fmt.Sprintf("%v", l.value), nil
}

func (a *AstPrinter) visitUnary(u *Unary) (any, error) {
	return a.parenthesize(
		u.operator.lexeme,
		u.right,
	)
}

func (a *AstPrinter) visitGrouping(g *Grouping) (any, error) {
	return a.parenthesize(
		"Group ",
		g.groupedExpression,
	)
}
