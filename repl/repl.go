package repl

import (
	"bufio"
	"fmt"
	"io"
	"text-adventure/game"
	"text-adventure/lexer"
	"text-adventure/parser"
	"text-adventure/vocabulary"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	vocab := vocabulary.DefaultVocabulary()
	world := game.NewWorld()
	cmdHandler := game.NewCmdHandler(world)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line, vocab)
		p := parser.New(l)

		cmd := p.ParseCommand()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := cmdHandler.Execute(*cmd)

		io.WriteString(out, evaluated)
		io.WriteString(out, "\n")

	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
