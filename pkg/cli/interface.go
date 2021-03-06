package cli

import (
	"io"

	"github.com/cappyzawa/op-kv/pkg/helper"
)

// Stream describes stream of cli
type Stream struct {
	In  io.Reader
	Out io.Writer
	Err io.Writer
}

// Params interface provides
type Params interface {
	Runner(opts ...helper.RunnerOpts) helper.Runner
	Printer(opts ...helper.PrinterOpts) helper.Printer
}
