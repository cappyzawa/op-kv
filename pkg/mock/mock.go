package mock

import (
	"github.com/cappyzawa/op-kv/pkg/cli"
	"github.com/cappyzawa/op-kv/pkg/helper"
)

var _ cli.Params = (*Params)(nil)
var _ helper.Runner = (*Runner)(nil)

// Parmas mocks cli.Params
type Params struct {
	MockRunner  func(opts ...helper.RunnerOpts) helper.Runner
	MockPrinter func(opts ...helper.PrinterOpts) helper.Printer
}

// Runner mocks Params.Runner
func (mp *Params) Runner(opts ...helper.RunnerOpts) helper.Runner {
	return mp.MockRunner(opts...)
}

// Runner mocks helper.Runner
type Runner struct {
	MockOutput       func(args []string) ([]byte, error)
	MockOutputWithIn func(args []string, in string) ([]byte, error)
	MockSignin       func(subdomain, password string) (*string, error)
}

// Output mocks Runner.Output
func (mr *Runner) Output(args []string) ([]byte, error) {
	return mr.MockOutput(args)
}

// OutputWithIn mocks Runner.OutputWithIn
func (mr *Runner) OutputWithIn(args []string, in string) ([]byte, error) {
	return mr.MockOutputWithIn(args, in)
}

// Signin mocks Runner.Signin
func (mr *Runner) Signin(subdomain, password string) (*string, error) {
	return mr.Signin(subdomain, password)
}

// Printer mocks Params.Printer
func (mp *Params) Printer(opts ...helper.PrinterOpts) helper.Printer {
	return mp.MockPrinter(opts...)
}

// Printer mocks helper.Printer
type Printer struct {
	MockPair   func(username, password string) error
	MockHeader func() (err error)
}

// Pair mocks Printer.Pair
func (mp *Printer) Pair(username, password string) error {
	return mp.MockPair(username, password)
}

// Header mocks Printer.Header
func (mp *Printer) Header() (err error) {
	return mp.Header()
}
