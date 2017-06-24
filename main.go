package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/gideondsouza/monkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s, this is MONKEY! the language.\n", user.Username)
	fmt.Printf("Go ahead and slap in some commands!\n")
	repl.Start(os.Stdin, os.Stdout)

}
