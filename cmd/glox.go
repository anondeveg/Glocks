package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Glox struct {
	hadError bool
}

func (e Glox) runFile(path string) error {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("could not read file %s: %w", path, err)
	}

	if err := e.run(string(bytes)); err != nil {
		return fmt.Errorf("runtime error: %w", err)
	}

	if e.hadError {
		os.Exit(65)
	}
	return nil
}

func (e Glox) runPrompt() {
	reader := bufio.NewReader(os.Stdin)
	for {

		var line string
		fmt.Print("> ")
		line, _ = reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "quit" {
			os.Exit(0)
		}

		e.run(line)

	}
}

func (e Glox) run(source string) error {
	sc := ScannerInit(source, &e)
	tokens := sc.ScanTokens()
	fmt.Println(tokens)
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
