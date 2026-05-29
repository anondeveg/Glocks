package main

import (
	"fmt"
)

func (e *Glox) report(line int, where string, message string) {
	fmt.Printf("[line  %d ] Error  %s  : %s", line, where, message)
	e.hadError = true
}

// A helper for simple errors
func (e *Glox) error(line int, message string) {
	e.report(line, "", message)
}
