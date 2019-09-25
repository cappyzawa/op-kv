package list

import (
	"bytes"
	"io"
	"os"

	"github.com/cappyzawa/op-kv/pkg/cli"
	"github.com/spf13/cobra"
)

// Options describes list options
type Options struct {
}

// NewOptions initializes list options
func NewOptions(outStream, errStream io.Writer) *Options {
	return &Options{}
}

// NewCmd initializes list command
func NewCmd(s *cli.Stream, p cli.Params) *cobra.Command {
	o := NewOptions(os.Stdout, new(bytes.Buffer))
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Display item titles",
		Run: func(cmd *cobra.Command, args []string) {
			o.Run(cmd, args)
		},
	}
	cmd.SetOutput(s.Out)
	cmd.SetErr(s.Err)
	return cmd
}

// Run runs list command
func (o *Options) Run(cmd *cobra.Command, args []string) {
}
