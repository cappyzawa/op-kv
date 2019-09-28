package helper

import (
	"fmt"
	"io"
	"os"
)

const (
	UsernameLen = 20
	PasswordLen = 60
)

type printer struct {
	Out io.Writer
}

// PrinterOpts describes options for Printer
type PrinterOpts func(*printer)

// PrinterOut sets PrinterOut of Opts optionally
func PrinterOut(out io.Writer) PrinterOpts {
	return func(o *printer) {
		o.Out = out
	}
}

// Printer prints command's output
type Printer interface {
	Pair(username, password string) error
	Header() (err error)
}

func NewPrinter(opts ...PrinterOpts) Printer {
	p := &printer{
		Out: os.Stdout,
	}
	for _, opt := range opts {
		opt(p)
	}
	return p
}

// Pair prints username & password
func (p *printer) Pair(username, password string) error {
	_, err := fmt.Fprintf(p.Out, "| %-20s| %-60s|\n", username, password)
	return err
}

// Header prints "USERNAME" and "PASSWORD" as header
func (p *printer) Header() (err error) {
	_, err = fmt.Fprintf(p.Out, "| %-20s| %-60s|\n", "USERNAME", "PASSWORD")
	for i := 0; i < UsernameLen+PasswordLen+3+2; i++ {
		_, err = fmt.Fprint(p.Out, "-")
	}
	_, err = fmt.Fprintf(p.Out, "\n")
	return
}
