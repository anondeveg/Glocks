package main

type Expr interface {
	accept(Visitor) any
}

type Visitor interface {
	visitBinary(b *Binary) any
	visitLiteral(l *Literal) any
	visitGrouping(g *Grouping) any
	visitUnary(u *Unary) any
}

type Binary struct {
	left     Expr
	operator Token
	right    Expr
}

func (b *Binary) accept(v Visitor) any {
	return v.visitBinary(b)
}

type Literal struct {
	value any
}

func (l *Literal) accept(v Visitor) any {
	return v.visitLiteral(l)
}

type Grouping struct {
	groupedExpression Expr
}

func (g *Grouping) accept(v Visitor) any {
	return v.visitGrouping(g)
}

type Unary struct {
	operator Token
	right    Expr
}

func (u *Unary) accept(v Visitor) any {
	return v.visitUnary(u)
}
