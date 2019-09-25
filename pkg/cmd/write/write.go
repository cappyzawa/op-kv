package write

import (
	"github.com/cappyzawa/op-kv/pkg/cli"
	"github.com/spf13/cobra"
)

// Options describes write options
type Options struct {
}

// NewOptions initializes write options
func NewOptions() *Options {
	return &Options{}
}

// NewCmd initializes write command
func NewCmd(s *cli.Stream, p cli.Params) *cobra.Command {
	o := NewOptions()
	cmd := &cobra.Command{
		Use:   "write <item> <password>",
		Short: "Generate one password by specified item and password",
		Run: func(cmd *cobra.Command, args []string) {
			o.Run(cmd, args)
		},
	}
	cmd.SetOutput(s.Out)
	cmd.SetErr(s.Err)
	return cmd
}

// Run runs write command
func (o *Options) Run(cmd *cobra.Command, args []string) {
	return
}
