package read_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/cappyzawa/op-kv/command/read"
	"github.com/cappyzawa/op-kv/command/util/utilfakes"
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
		{name: "ensure StdOut", args: []string{"item"}, mockOut: []byte("test"), mockErr: nil, expect: "test", expectError: ""},
	}

	t.Helper()
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			f := new(utilfakes.FakeFactory)
			outStream := new(bytes.Buffer)
			errStream := new(bytes.Buffer)
			options := read.NewOptions(outStream, errStream)
			cmd := read.NewCmdRead(f)
			opCmd := []string{"op", "get", "item", c.args[0]}
			jqCmd := []string{"jq", "-r", ".details.fields[] | select(.designation==\"password\").value"}
			runner := new(opkvfakes.FakeRunner)
			f.CommandRunnerReturns(runner)
			runner.Output(opCmd, jqCmd)
			runner.OutputReturns(c.mockOut, c.mockErr)
			args := runner.OutputArgsForCall(0)
			if reflect.DeepEqual(args[0][0], opCmd) {
				t.Errorf("op command is passed as first arg")
			}
			if reflect.DeepEqual(args[0][1], jqCmd) {
				t.Errorf("jq command is passed as second arg")
			}
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
