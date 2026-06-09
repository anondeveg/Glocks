package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Glox struct {
	hadError        bool
	hadRuntimeError bool
}

func (g Glox) runFile(path string) error {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("could not read file %s: %w", path, err)
	}

	if err := g.run(string(bytes)); err != nil {
		return fmt.Errorf("runtime error: %w", err)
	}

	if g.hadError {
		os.Exit(65)
	}
	if g.hadRuntimeError {
		os.Exit(70)
	}
	return nil
}

func (g Glox) runPrompt() {
	reader := bufio.NewReader(os.Stdin)
	for {

		var line string
		fmt.Print("\n> ")
		line, _ = reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "quit" {
			os.Exit(0)
		}

		g.run(line)
	}
}

func (g *Glox) run(source string) error {
	sc := ScannerInit(source, g)
	tokens := sc.ScanTokens()
	parser := Parser{g, 0, tokens}
	AST := parser.parse()
	intrepeter := interpreter{g}
	_, err := intrepeter.interpret(AST)
	if err != nil {
		g.reportRuntime(err.(RuntimeError))
		return err
	}
	if g.hadError || g.hadRuntimeError {
		return errors.New("")
	}

	return nil
}

func GloxInit() Glox {
	glox := Glox{}
	fmt.Println(len(os.Args))
	if len(os.Args) > 2 {
		fmt.Println("Usage: Glox [script]")
		os.Exit(64) // https://gist.github.com/bojanrajkovic/831993 line 100.
	} else if len(os.Args) == 2 {
		glox.runFile(os.Args[1])
	} else {
		glox.runPrompt()
	}

	return glox
}

func main() {
	GloxInit()
}
