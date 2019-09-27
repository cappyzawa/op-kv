package mock

import (
	"github.com/cappyzawa/op-kv/pkg/cli"
	"github.com/cappyzawa/op-kv/pkg/helper"
)

var _ cli.Params = (*Params)(nil)
var _ helper.Runner = (*Runner)(nil)

type Params struct {
	MockRunner func(opts ...helper.Opts) helper.Runner
}

func (mp *Params) Runner(opts ...helper.Opts) helper.Runner {
	return mp.MockRunner(opts...)
}

type Runner struct {
	MockOutput       func(args []string) ([]byte, error)
	MockOutputWithIn func(args []string, in string) ([]byte, error)
	MockSignin       func(subdomain, password string) (*string, error)
}

func (mr *Runner) Output(args []string) ([]byte, error) {
	return mr.MockOutput(args)
}

func (mr *Runner) OutputWithIn(args []string, in string) ([]byte, error) {
	return mr.MockOutputWithIn(args, in)
}

func (mr *Runner) Signin(subdomain, password string) (*string, error) {
	return mr.Signin(subdomain, password)
}
