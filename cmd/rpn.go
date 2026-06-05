// this was an excersice for chapter 5 and is no longer updated to the current version of the parser :) please do not run.
package main

import "fmt"

type Rpn struct{}

func (r *Rpn) visitLiteral(l *Literal) any {
	if l.value == nil {
		return "nil"
	}
	return fmt.Sprintf("%v ", l.value)
}

func (r *Rpn) visitBinary(b *Binary) any {
	return fmt.Sprintf("%v %v %v", b.left.accept(r), b.right.accept(r), b.operator.lexeme)
}

func (r *Rpn) visitUnary(u *Unary) any {
	return fmt.Sprintf("%v %v", u.right.accept(r), u.operator.lexeme)
}

func (r *Rpn) visitGrouping(g *Grouping) any {
	return fmt.Sprintf("%v", g.groupedExpression.accept(r))
}
