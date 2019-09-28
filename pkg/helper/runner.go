package helper

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

type runner struct {
	Path string
	Out  io.Writer
	Err  io.Writer
	In   io.Reader
}

// RunnerOpts describes options for Runner
type RunnerOpts func(*runner)

// RunnerPath sets RunnerPath of Opts optionally
func RunnerPath(path string) RunnerOpts {
	return func(o *runner) {
		o.Path = path
	}
}

// RunnerOut sets RunnerOut of Opts optionally
func RunnerOut(out io.Writer) RunnerOpts {
	return func(o *runner) {
		o.Out = out
	}
}

// RunnerErr sets RunnerErr of Opts optionally
func RunnerErr(err io.Writer) RunnerOpts {
	return func(o *runner) {
		o.Err = err
	}
}

// Runner runs ex command
type Runner interface {
	Output(args []string) ([]byte, error)
	OutputWithIn(args []string, in string) ([]byte, error)
	Signin(subdomain, password string) (*string, error)
}

// NewRunner initilizes runner
func NewRunner(opts ...RunnerOpts) Runner {
	r := &runner{
		Path: "op",
		Out:  os.Stdout,
		Err:  os.Stderr,
	}
	for _, opt := range opts {
		opt(r)
	}
	return r
}

func (r *runner) Output(args []string) ([]byte, error) {
	c := exec.Command(r.Path, args...)
	c.Stderr = r.Err
	output, err := c.Output()
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (r *runner) OutputWithIn(args []string, in string) ([]byte, error) {
	c := exec.Command(r.Path, args...)
	c.Stderr = r.Err
	eS := new(bytes.Buffer)
	c.Stderr = eS
	stdin, err := c.StdinPipe()
	if err != nil {
		return nil, err
	}
	go func() {
		defer stdin.Close()
		io.WriteString(stdin, in)
	}()

	output, err := c.Output()
	if err != nil {
		fmt.Println(eS.String())
		return nil, err
	}

	return output, nil
}

// Signin gets session token of subdomain using by password
func (r *runner) Signin(subdomain, password string) (*string, error) {
	// overwride r.Path with "op"
	if r.Path != "op" {
		r.Path = "op"
	}

	// get a session token
	st, err := r.OutputWithIn([]string{"signin", subdomain, "--output=raw"}, password)
	if err != nil {
		return nil, err
	}

	trimmedSt := strings.TrimSpace(string(st))
	return &trimmedSt, nil
}
