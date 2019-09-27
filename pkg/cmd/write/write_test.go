package write_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/cappyzawa/op-kv/pkg/cmd/write"
	"github.com/cappyzawa/op-kv/pkg/helper"
	"github.com/cappyzawa/op-kv/pkg/mock"
	"github.com/spf13/cobra"
)

func TestOptionsRun(t *testing.T) {
	r := &mock.Runner{}
	p := &mock.Params{}
	o := write.NewOptions()
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
			name:     "with key and value args",
			args:     []string{"key", "value"},
			expect:   fmt.Sprintf("success to write password to \"%s\"\n", "key"),
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
				if args[0] == "get" {
					return ioutil.ReadFile("../../../testdata/login_template.json")
				}
				return nil, nil
			}
			r.MockOutputWithIn = func(args []string, in string) ([]byte, error) {
				return []byte("XXXXXXXXXXX"), nil
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
