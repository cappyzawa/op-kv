package list_test

import (
	"bytes"
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/cappyzawa/op-kv/pkg/cli"
	"github.com/cappyzawa/op-kv/pkg/cmd/list"
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
	cc := list.NewCmd(s, p)
	if !reflect.DeepEqual(cc.ErrOrStderr(), errStream) {
		t.Errorf("cmd's err is %#v, errStream is %#v", cc.ErrOrStderr(), errStream)
	}
	if cc.ErrOrStderr().(*bytes.Buffer).String() != expect {
		t.Errorf("cc.Err should be %s", expect)
	}
}

func TestOptionsRun(t *testing.T) {
	r := &mock.Runner{}
	p := &mock.Params{}
	o := list.NewOptions()
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
			expect:   "test1\ntest2\ntest3\n",
			existErr: false,
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
			r.MockOutput = func(args []string) ([]byte, error) {
				return ioutil.ReadFile("../../../testdata/list.json")
			}
			p.MockRunner = func(opts ...helper.Opts) helper.Runner {
				return r
			}
			err := o.Run(p, cc, c.args)
			if !c.existErr && err != nil {
				t.Errorf("stderr should not be occurred, but actual is %v", err)
			}
			if c.existErr && err == nil {
				t.Error("error should be occurred, but it doesn't occurred")
			}
			if outStream.String() != c.expect {
				t.Errorf("stdout should be %v, but actual is %v", c.expect, outStream.String())
			}
		})
	}
}
