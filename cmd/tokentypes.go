package main

type TokenType int

const (
	// Single-character tokens.
	LEFT_PAREN TokenType = iota
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	SLASH
	STAR
	// One or two character tokens.
	EQUAL
	BANG
	BANG_EQUAL
	EQUAL_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL
	// Literals.
	IDENTIFIER
	STRING
	NUMBER
	// Keywords.
	IF
	AND
	CLASS
	ELSE
	FALSE
	FUN
	FOR
	NIL
	OR
	PRINT
	RETURN
	SUPER
	THIS
	TRUE
	VAR
	WHILE
	EOF
)
