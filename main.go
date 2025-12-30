package main

import (
	"fmt"
	"os"
	"os/user"
	"text-adventure/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! You have woken up\n",
		user.Username)
	fmt.Printf("Feel free to type in commands to look around and move\n")
	repl.Start(os.Stdin, os.Stdout)
}
