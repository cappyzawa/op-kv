package opkv

import (
	"bytes"
	"io"
	"os/exec"
)

type Runner struct {
	Command   string
	OutStream io.Writer
	ErrStream io.Writer
}

func NewRunner(command string) *Runner {
	var (
		stdOut bytes.Buffer
		stdErr bytes.Buffer
	)
	return &Runner{
		Command:   command,
		OutStream: &stdOut,
		ErrStream: &stdErr,
	}
}

func (r *Runner) Run(args []string) error {
	cmd := exec.Command(r.Command, args...)
	cmd.Stdout = r.OutStream
	cmd.Stderr = r.ErrStream
	return cmd.Run()
}

func (r *Runner) Output(args []string) ([]byte, error) {
	cmd := exec.Command(r.Command, args...)
	cmd.Stdout = r.OutStream
	cmd.Stderr = r.ErrStream
	return cmd.Output()
}
