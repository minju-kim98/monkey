package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/minju-kim98/monkey/lexer"
	"github.com/minju-kim98/monkey/token"
)

const PROMPT = ">>"

func Repl(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		input := scanner.Scan()
		if !input {
			return
		}

		scanned := scanner.Text()

		if scanned == "exit" {
			break
		}

		l := lexer.New(scanned)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
