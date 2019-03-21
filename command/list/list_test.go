package list_test

import (
	"bytes"
	"reflect"
	"strings"
	"testing"

	"github.com/cappyzawa/op-kv/command/list"
	"github.com/cappyzawa/op-kv/command/util/utilfakes"
	"github.com/cappyzawa/op-kv/op-kvfakes"
)

func TestOptions_Run(t *testing.T) {
	cases := []struct {
		name        string
		args        []string
		mockOut     []byte
		mockErr     error
		expect      []string
		expectError string
	}{
		{name: "ensure StdOut", args: nil, mockOut: []byte("test1\ntest2"), mockErr: nil, expect: []string{"test1", "test2"}, expectError: ""},
	}

	t.Helper()
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			f := new(utilfakes.FakeFactory)
			outStream := new(bytes.Buffer)
			errStream := new(bytes.Buffer)
			options := list.NewOptions(outStream, errStream)
			cmd := list.NewCmdList(f)
			opCmd := []string{"op", "list", "items"}
			jqCmd := []string{"jq", "-r", ".[]overview.title"}
			runner := new(opkvfakes.FakeRunner)
			f.CommandRunnerReturns(runner)
			runner.Output(opCmd, jqCmd)
			runner.OutputReturns(c.mockOut, c.mockErr)
			options.Run(f, cmd, nil)

			actual := strings.Split(outStream.String(), "\n")
			if !reflect.DeepEqual(actual, c.expect) {
				t.Errorf("%v should be displayed, but actual is %v", c.expect, actual)
			}
			if errStream.String() != c.expectError {
				t.Errorf("%s should be displayed, but actual is %s", c.expectError, errStream.String())
			}
		})
	}
}
