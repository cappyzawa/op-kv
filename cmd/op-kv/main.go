package main

import (
	"fmt"
	"os"

	"github.com/cappyzawa/op-kv/command"
)

func main() {
	cmd := command.NewCmd()
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
