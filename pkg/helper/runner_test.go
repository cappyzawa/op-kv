package helper_test

import (
	"bytes"
	"testing"

	"github.com/cappyzawa/op-kv/pkg/helper"
)

func TestRunnerOutput(t *testing.T) {
	var (
		outStream *bytes.Buffer
		errStream *bytes.Buffer
	)
	cases := []struct {
		name     string
		path     helper.RunnerOpts
		args     []string
		expect   string
		existErr bool
	}{
		{
			name:     "echo",
			path:     helper.RunnerPath("echo"),
			args:     []string{"-n", "test"},
			expect:   "test",
			existErr: false,
		},
		{
			name:     "invalid path",
			path:     helper.RunnerPath("invalid"),
			args:     []string{},
			expect:   "",
			existErr: true,
		},
		{
			name:     "date with invalid args",
			path:     helper.RunnerPath("date"),
			args:     []string{"--invalid"},
			expect:   "",
			existErr: true,
		},
	}

	t.Helper()
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			outStream = new(bytes.Buffer)
			errStream = new(bytes.Buffer)
			runner := helper.NewRunner(c.path,
				helper.RunnerOut(outStream),
				helper.RunnerErr(errStream))
			out, err := runner.Output(c.args)

			if !c.existErr && err != nil {
				t.Errorf("error should not be occurred, but actual is %v", err)
			}
			if c.existErr && err == nil {
				t.Error("error should be occurred, but it doesn't occurred")
			}
			if string(out) != c.expect {
				t.Errorf("output should be %s, but actual is %s", c.expect, string(out))
			}
		})
	}
}

func TestRunnerOutputWithIn(t *testing.T) {
	var (
		outStream *bytes.Buffer
		errStream *bytes.Buffer
	)
	cases := []struct {
		name     string
		path     helper.RunnerOpts
		args     []string
		in       string
		expect   string
		existErr bool
	}{
		{
			name:     "cat",
			path:     helper.RunnerPath("cat"),
			args:     []string{},
			in:       "test",
			expect:   "test",
			existErr: false,
		},
		{
			name:     "cat with empty string",
			path:     helper.RunnerPath("cat"),
			args:     []string{},
			in:       "",
			expect:   "",
			existErr: false,
		},
		{
			name:     "invalid path",
			path:     helper.RunnerPath("invalid"),
			args:     []string{},
			in:       "",
			expect:   "",
			existErr: true,
		},
	}

	t.Helper()
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			outStream = new(bytes.Buffer)
			errStream = new(bytes.Buffer)
			runner := helper.NewRunner(c.path,
				helper.RunnerOut(outStream),
				helper.RunnerErr(errStream))
			out, err := runner.OutputWithIn(c.args, c.in)
			if !c.existErr && err != nil {
				t.Errorf("error should not be occurred, but actual is %v", err)
			}
			if c.existErr && err == nil {
				t.Error("error should be occurred, but it doesn't occurred")
			}
			if string(out) != c.expect {
				t.Errorf("output should be %s, but actual is %s", c.expect, string(out))
			}
		})
	}
}
