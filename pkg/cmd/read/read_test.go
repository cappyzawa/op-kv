package read_test

import (
	"bytes"
	"errors"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/cappyzawa/op-kv/pkg/cli"
	"github.com/cappyzawa/op-kv/pkg/cmd/read"
	"github.com/cappyzawa/op-kv/pkg/helper"
	"github.com/spf13/cobra"
)

func TestOptionsRun(t *testing.T) {
	r := &mockRunner{}
	p := &mockParams{}
	o := read.NewOptions()
	st := "session token"
	o.SessionToken = &st

	cases := []struct {
		name     string
		args     []string
		expect   string
		existErr bool
	}{
		{
			name:     "with zero args",
			args:     []string{},
			expect:   "",
			existErr: true,
		},
		{
			name:     "with an item arg",
			args:     []string{"test"},
			expect:   "password",
			existErr: false,
		},
		{
			name:     "item is missing",
			args:     []string{"missing item"},
			expect:   "",
			existErr: true,
		},
	}

	t.Helper()
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			cc := &cobra.Command{}
			outStream := new(bytes.Buffer)
			errStream := new(bytes.Buffer)
			cc.SetOut(outStream)
			cc.SetErr(errStream)
			r.output = func(args []string) ([]byte, error) {
				if strings.Contains(c.name, "missing") {
					return nil, errors.New("missing item")
				}
				return ioutil.ReadFile("../../../testdata/op_get.json")
			}
			p.runner = func(opts ...helper.Opts) helper.Runner {
				return r
			}
			err := o.Run(p, cc, c.args)
			if !c.existErr && err != nil {
				t.Errorf("stderr should be empty, but actual is %v", errStream.String())
			}
			if c.existErr && err == nil {
				t.Error("error should be occurred, but it doesn't occurred")
			}
			if outStream.String() != c.expect {
				t.Errorf("stdout should be password, but actual is %v", outStream.String())
			}
		})
	}
}

var _ cli.Params = (*mockParams)(nil)
var _ helper.Runner = (*mockRunner)(nil)

type mockParams struct {
	runner func(opts ...helper.Opts) helper.Runner
}

func (mp *mockParams) Runner(opts ...helper.Opts) helper.Runner {
	return mp.runner(opts...)
}

type mockRunner struct {
	output       func(args []string) ([]byte, error)
	outputWithIn func(args []string, in string) ([]byte, error)
	signin       func(subdomain, password string) (*string, error)
}

func (mr *mockRunner) Output(args []string) ([]byte, error) {
	return mr.output(args)
}

func (mr *mockRunner) OutputWithIn(args []string, in string) ([]byte, error) {
	return mr.outputWithIn(args, in)
}

func (mr *mockRunner) Signin(subdomain, password string) (*string, error) {
	return mr.Signin(subdomain, password)
}
