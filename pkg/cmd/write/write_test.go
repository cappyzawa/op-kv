package write_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/cappyzawa/op-kv/pkg/cli"
	"github.com/cappyzawa/op-kv/pkg/cmd/write"
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
	cc := write.NewCmd(s, p)
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

	cases := []struct {
		name     string
		args     []string
		options  *write.Options
		expect   string
		existErr bool
	}{
		{
			name: "with zero args",
			args: []string{},
			options: &write.Options{
				Password: "password",
				Username: "",
			},
			expect:   "",
			existErr: true,
		},
		{
			name: "with key and password",
			args: []string{"key"},
			options: &write.Options{
				Password: "password",
				Username: "",
			},
			expect:   fmt.Sprintf("success to write password (%s) and username (%s) to \"%s\"\n", "password", "", "key"),
			existErr: false,
		},
		{
			name: "with key, username and password",
			args: []string{"key"},
			options: &write.Options{
				Password: "password",
				Username: "username",
			},
			expect:   fmt.Sprintf("success to write password (%s) and username (%s) to \"%s\"\n", "password", "username", "key"),
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
			p.MockRunner = func(opts ...helper.RunnerOpts) helper.Runner {
				return r
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
			if outStream.String() != c.expect {
				t.Errorf("stdout should be %v, but actual is %v", c.expect, outStream.String())
			}
		})
	}
}
