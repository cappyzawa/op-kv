package helper

import (
	"bytes"
	"os/exec"
)

type runner struct {
	Path string
	Out  *bytes.Buffer
	Err  *bytes.Buffer
}

// Opts describes options for Runner
type Opts func(*runner)

// Path sets Path of Opts optionally
func Path(path string) Opts {
	return func(o *runner) {
		o.Path = path
	}
}

// Out sets Out of Opts optionally
func Out(out *bytes.Buffer) Opts {
	return func(o *runner) {
		o.Out = out
	}
}

// Err sets Err of Opts optionally
func Err(err *bytes.Buffer) Opts {
	return func(o *runner) {
		o.Err = err
	}
}

//go:generate counterfeiter . Runner
// Runner runs ex command
type Runner interface {
	Output(commands []string) ([]byte, error)
}

// NewRunner initilizes runner
func NewRunner(opts ...Opts) Runner {
	r := &runner{
		Path: "op",
	}
	for _, opt := range opts {
		opt(r)
	}
	return r
}

func (r *runner) Output(args []string) ([]byte, error) {
	c := exec.Command(r.Path, args...)
	c.Stdout = r.Out
	c.Stderr = r.Err
	if err := c.Run(); err != nil {
		return nil, err
	}
	return r.Out.Bytes(), nil
}
