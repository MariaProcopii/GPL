package reader

import (
	"fmt"
	"os"

	"github.com/MariaProcopii/GPL/lexer"
	"github.com/MariaProcopii/GPL/token"
)

func ReadFile(filepath string) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	l := lexer.New(string(data))

	for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
		fmt.Printf("%+v\n", tok)
	}
}
