package repl

import (
	"Zeon/Programming-Language/evaluator"
	"Zeon/Programming-Language/lexer"
	"Zeon/Programming-Language/object"
	"Zeon/Programming-Language/parser"
	"bufio"
	"fmt"
	"io"
)

const PROMPT = ">>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParseErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}

	}
}

func printParseErrors(out io.Writer, errors []string) {

	io.WriteString(out, "Woops! 🥲, We ran into some Neon Zone of Zeon\n")
	io.WriteString(out, "Parser Errors...\n")

	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\t")
	}
}
