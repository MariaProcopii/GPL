package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/MariaProcopii/GPL/reader"
	"github.com/MariaProcopii/GPL/repl"
)

func main() {
	cmdParams := os.Args[1:]
	if len(cmdParams) > 0 {
		reader.ReadFile(cmdParams[0], os.Stdout)
	} else {
		user, err := user.Current()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Hello %s!\n", user.Username)

		repl.Start(os.Stdin, os.Stdout)
	}
}
