package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/MariaProcopii/GPL/reader"
	"github.com/MariaProcopii/GPL/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s!\n", user.Username)

	cmdParams := os.Args[1:]
	if len(cmdParams) > 0 {
		reader.ReadFile(cmdParams[0])
	} else {
		repl.Start(os.Stdin, os.Stdout)
	}
}
