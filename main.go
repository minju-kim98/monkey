package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/minju-kim98/monkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n",
		user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Repl(os.Stdin, os.Stdout)
}
