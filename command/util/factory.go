package util

import "github.com/cappyzawa/op-kv"

type Factory interface {
	CommandRunner(command string) *opkv.Runner
}
