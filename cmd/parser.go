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

type Parser struct {
	glox    *Glox
	current int
	tokens  []Token
}

func (p *Parser) parse() Expr {
	p.current = 0
	ret := p.expression()
	return ret
}

func (p *Parser) expression() Expr {
	return p.equality()
}

func (p *Parser) equality() Expr {
	expr := p.comparison()
	for p.match(BANG_EQUAL, BANG) {
		operator := p.previous()
		right := p.comparison()
		expr = &Binary{expr, operator, right}
	}
	return expr
}

func (p *Parser) isAtEnd() bool {
	return p.current >= len(p.tokens)
}

func (p *Parser) previous() Token {
	return p.tokens[p.current-1]
}

func (p *Parser) peek() Token {
	return p.tokens[p.current]
}

func (p *Parser) advance() {
	p.current++
}

func (p *Parser) match(tokens ...TokenType) bool {
	for _, token := range tokens {
		if p.tokens[p.current].ttype == token {
			p.advance()
			return true
		}
	}
	return false
}

func (p *Parser) comparison() Expr {
	expr := p.term()
	for p.match(GREATER, GREATER_EQUAL, LESS, LESS_EQUAL) {
		operator := p.previous()
		right := p.term()
		expr = &Binary{expr, operator, right}
	}
	return expr
}

func (p *Parser) term() Expr {
	expr := p.factor()
	for p.match(MINUS, PLUS) {
		operator := p.previous()
		right := p.factor()
		expr = &Binary{expr, operator, right}
	}
	return expr
}

func (p *Parser) factor() Expr {
	expr := p.unary()
	for p.match(STAR, SLASH) {
		operator := p.previous()
		right := p.unary()
		expr = &Binary{expr, operator, right}
	}
	return expr
}

func (p *Parser) unary() Expr {
	var expr any = false
	for p.match(BANG, MINUS) {
		operator := p.previous()
		right := p.unary()
		expr = &Unary{operator, right}
	}
	if expr != false {
		return expr.(Expr)
	}
	return p.primary()
}

func (p *Parser) primary() Expr {
	switch p.peek().ttype {
	case FALSE:
		p.advance()
		return &Literal{false}
	case TRUE:
		p.advance()
		return &Literal{true}
	case NIL:
		p.advance()
		return &Literal{nil}

	case NUMBER, STRING:
		p.advance()
		return &Literal{p.previous().literal}
	case LEFT_PAREN:
		p.advance()
		expr := p.expression()
		p.consume(RIGHT_PAREN, "Expect ')' aftrer expression")
		return &Grouping{expr}
	default:
		p.glox.error(p.peek().line, "Expected 'false','true','nil',Number,String, or '('")
	}

	return &Literal{nil} // dead code
}

func (p *Parser) consume(ttype TokenType, message string) {
	if p.tokens[p.current].ttype == ttype {
		p.advance()
	} else {
		p.glox.error(p.peek().line, message)
	}
}
