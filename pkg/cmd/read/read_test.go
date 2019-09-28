package read_test

import (
	"bytes"
	"errors"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/cappyzawa/op-kv/pkg/cmd/read"
	"github.com/cappyzawa/op-kv/pkg/helper"
	"github.com/cappyzawa/op-kv/pkg/mock"
	"github.com/spf13/cobra"
)

func TestOptionsRun(t *testing.T) {
	r := &mock.Runner{}
	p := &mock.Params{}
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
			r.MockOutput = func(args []string) ([]byte, error) {
				if strings.Contains(c.name, "missing") {
					return nil, errors.New("missing item")
				}
				return ioutil.ReadFile("../../../testdata/op_get.json")
			}
			p.MockRunner = func(opts ...helper.RunnerOpts) helper.Runner {
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
