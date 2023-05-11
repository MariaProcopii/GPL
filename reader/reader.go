package reader

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/MariaProcopii/GPL/token"

	"github.com/MariaProcopii/GPL/evaluator"
	"github.com/MariaProcopii/GPL/lexer"
	"github.com/MariaProcopii/GPL/object"
	"github.com/MariaProcopii/GPL/parser"
)

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}

func ReadFile(filepath string, out io.Writer) {
	env := object.NewEnvironment()

	data, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	input := strings.Split(string(data), "\n")
	l := lexer.New(string(data))
	for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
		fmt.Printf("%+v\n", tok)
	}

	for _, i := range input {
		fmt.Println(i)
		l := lexer.New(i)
		p := parser.New(l)
		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}

}
