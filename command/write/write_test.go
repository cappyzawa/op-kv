package write_test

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/cappyzawa/op-kv/command/util/utilfakes"
	"github.com/cappyzawa/op-kv/command/write"
	"github.com/cappyzawa/op-kv/op-kvfakes"
)

func TestOptions_Run(t *testing.T) {
	cases := []struct {
		name        string
		args        []string
		mockOut     []byte
		mockErr     error
		expect      string
		expectError string
	}{
		{name: "ensure StdOut", args: []string{"item", "password"}, mockOut: nil, mockErr: nil, expect: "success to write password!!", expectError: ""},
		{name: "ensure StdErr when failing to write", args: []string{"item", "password"}, mockOut: nil, mockErr: errors.New("some error"), expect: "", expectError: "some error"},
	}

	t.Helper()
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			item := c.args[0]
			password := c.args[1]
			runner := new(opkvfakes.FakeRunner)
			f := new(utilfakes.FakeFactory)
			outStream := new(bytes.Buffer)
			errStream := new(bytes.Buffer)
			options := write.NewOptions(outStream, errStream)
			cmd := write.NewCmdWrite(f)
			opTmpCmd := []string{"op", "get", "template", "login"}
			jqCmd := []string{"jq", "-c", fmt.Sprintf(".fields[1].value = \"%s\"", password)}
			opEncCmd := []string{"op", "encode"}
			f.CommandRunnerReturns(runner)
			output, _ := runner.Output(opTmpCmd, jqCmd, opEncCmd)
			runner.OutputReturns([]byte("encoded"), nil)
			opCmd := []string{"op", "create", "item", "login", strings.TrimRight(string(output), "\n"), fmt.Sprintf("--title=%s", item)}
			runner.Output(opCmd)
			runner.OutputReturns(c.mockOut, c.mockErr)
			options.Run(f, cmd, c.args)
			if outStream.String() != c.expect {
				t.Errorf("%s should be displayed, but actual is %s", c.expect, outStream.String())
			}
			if errStream.String() != c.expectError {
				t.Errorf("%s should be displayed, but actual is %s", c.expectError, errStream.String())
			}
		})
	}
}
