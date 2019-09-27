package mock_test

import (
	"github.com/cappyzawa/op-kv/pkg/cli"
	"github.com/cappyzawa/op-kv/pkg/helper"
)

var _ cli.Params = (*Params)(nil)
var _ helper.Runner = (*Runner)(nil)

// Parmas mocks cli.Params
type Params struct {
	MockRunner func(opts ...helper.Opts) helper.Runner
}

// Runner mocks Params.Runner
func (mp *Params) Runner(opts ...helper.Opts) helper.Runner {
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
