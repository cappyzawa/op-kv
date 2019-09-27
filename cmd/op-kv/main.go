package main

import (
	"os"

	"github.com/cappyzawa/op-kv/pkg/cli"
	"github.com/cappyzawa/op-kv/pkg/cmd"
)

func main() {
	s := &cli.Stream{
		In:  os.Stdin,
		Out: os.Stdout,
		Err: os.Stderr,
	}
	c := cmd.NewCmd(s)
	c.Execute()
}
