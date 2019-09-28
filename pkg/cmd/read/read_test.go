package read_test

import (
	"bytes"
	"errors"
	"io/ioutil"
	"reflect"
	"strings"
	"testing"

	"github.com/cappyzawa/op-kv/pkg/cli"
	"github.com/cappyzawa/op-kv/pkg/cmd/read"
	"github.com/cappyzawa/op-kv/pkg/helper"
	"github.com/cappyzawa/op-kv/pkg/mock"
	"github.com/spf13/cobra"
)

func TestNewCmd(t *testing.T) {
	expect := "test"
	errStream := new(bytes.Buffer)
	errStream.Write([]byte(expect))
	s := &cli.Stream{
		Err: errStream,
	}
	p := &mock.Params{}
	cc := read.NewCmd(s, p)
	if !reflect.DeepEqual(cc.ErrOrStderr(), errStream) {
		t.Errorf("cmd's err is %#v, errStream is %#v", cc.ErrOrStderr(), errStream)
	}
	if cc.ErrOrStderr().(*bytes.Buffer).String() != expect {
		t.Errorf("cc.Err should be %s", expect)
	}
}

func TestOptionsRun(t *testing.T) {
	runner := &mock.Runner{}
	p := &mock.Params{}

	cases := []struct {
		name     string
		args     []string
		options  *read.Options
		expect   string
		existErr bool
	}{
		{
			name:     "with zero args",
			args:     []string{},
			options:  read.NewOptions(),
			expect:   "",
			existErr: true,
		},
		{
			name:     "with an item arg",
			args:     []string{"test"},
			options:  read.NewOptions(),
			expect:   "testPassword",
			existErr: false,
		},
		{
			name: "with an item arg and table flag",
			args: []string{"test"},
			options: &read.Options{
				Table: true,
			},
			expect:   "email@com",
			existErr: false,
		},
		{
			name:     "item is missing",
			args:     []string{"missing item"},
			options:  read.NewOptions(),
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
			runner.MockOutput = func(args []string) ([]byte, error) {
				if strings.Contains(c.name, "missing") {
					return nil, errors.New("missing item")
				}
				return ioutil.ReadFile("../../../testdata/op_get.json")
			}
			p.MockRunner = func(opts ...helper.RunnerOpts) helper.Runner {
				return runner
			}
			p.MockPrinter = func(opts ...helper.PrinterOpts) helper.Printer {
				return helper.NewPrinter(helper.PrinterOut(outStream))
			}
			st := "session token"
			c.options.SessionToken = &st
			err := c.options.Run(p, cc, c.args)
			if !c.existErr && err != nil {
				t.Errorf("stderr should not be occurred, but actual is %v", err)
			}
			if c.existErr && err == nil {
				t.Error("error should be occurred, but it doesn't occurred")
			}
			if !strings.Contains(outStream.String(), c.expect) {
				t.Errorf("stdout should contain %v, but actual is %v", c.expect, outStream.String())
			}
		})
	}
}
