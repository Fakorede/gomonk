package repl

import (
	"bufio"
	"fmt"
	"gomonk/lexer"
	"gomonk/token"
	"io"
)

const PROMPT = ">> "

// Start reads an input, sends it to the interpreter for evaluation, prints the output of the interpreter.
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()

		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
