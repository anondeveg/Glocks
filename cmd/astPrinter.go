package main

import (
	"fmt"
	"strings"
)

type AstPrinter struct{}

func (a *AstPrinter) log(expr Expr) any { // not print because golang uses it.
	return expr.accept(a)
}

func (a *AstPrinter) parenthesize(name string, exprs ...Expr) string {
	var builder strings.Builder
	builder.WriteString("( " + name)
	for _, expr := range exprs {
		builder.WriteString("")
		builder.WriteString(fmt.Sprintf("%v ", expr.accept(a)))
	}
	builder.WriteString(")")
	return builder.String()
}

func (a *AstPrinter) visitBinary(b *Binary) any {
	return a.parenthesize(b.operator.lexeme, b.left, b.right)
}

func (a *AstPrinter) visitLiteral(l *Literal) any {
	if l.value == nil {
		return "nil"
	}
	return fmt.Sprintf("%v", l.value)
}

func (a *AstPrinter) visitUnary(u *Unary) any {
	return a.parenthesize(u.operator.lexeme, u.right)
}

func (a *AstPrinter) visitGrouping(g *Grouping) any {
	return a.parenthesize("Group ", g.groupedExpression)
}
