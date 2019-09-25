package cli

import "io"

// Stream describes stream of cli
type Stream struct {
	In  io.Reader
	Out io.Writer
	Err io.Writer
}

// Params interface provides
type Params interface {
	Runner()
}
