package util

import "github.com/cappyzawa/op-kv"

//go:generate counterfeiter . Factory
type Factory interface {
	CommandRunner() opkv.Runner
}
