package cli

import "github.com/cappyzawa/op-kv/pkg/helper"

// OpKvParams describes parameters
type OpKvParams struct {
}

var _ Params = (*OpKvParams)(nil)

// Runner runs ex commands
func (p *OpKvParams) Runner(opts ...helper.RunnerOpts) helper.Runner {
	return helper.NewRunner(opts...)
}

// Printer prints outputs
func (p *OpKvParams) Printer(opts ...helper.PrinterOpts) helper.Printer {
	return helper.NewPrinter(opts...)
}
