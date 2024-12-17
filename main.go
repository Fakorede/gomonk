package main

import (
	"fmt"
	"gomonk/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! Welcome to the GoMonk language!\n", user.Username)
	fmt.Printf("Type in your commands to begin\n")

	repl.Start(os.Stdin, os.Stdout)
}
