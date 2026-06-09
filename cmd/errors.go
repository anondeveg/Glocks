package main

import (
	"fmt"
)

type RuntimeError struct {
	token   Token
	message string
}

var (
	RED          = "\033[31;1m" // Changes text color to Bold Red
	WHITE        = "\033[0m"    // Resets all formatting back to default
	ERRORMESSAGE = "\n\t%v%s%v\n\t on line: %d \n"
)

func (e RuntimeError) Error() string {
	return fmt.Sprintf(ERRORMESSAGE, RED, e.message, WHITE, e.token.line)
}

func (g *Glox) reportRuntime(e RuntimeError) {
	fmt.Println(e.Error())
	g.hadRuntimeError = true
}

func (g *Glox) report(line int, where string, message string) {
	fmt.Printf(ERRORMESSAGE, RED, message, WHITE, line)
	g.hadError = true
}

// A helper for simple errors
func (g *Glox) error(line int, message string) {
	g.report(line, "", message)
}
