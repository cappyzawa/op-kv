package read

import (
	"github.com/cappyzawa/op-kv/pkg/cli"
	"github.com/spf13/cobra"
)

// Options describes read options
type Options struct {
}

// NewOptions initializes read options
func NewOptions() *Options {
	return &Options{}
}

// NewCmd initializes read command
func NewCmd(s *cli.Stream, p cli.Params) *cobra.Command {
	o := NewOptions()
	cmd := &cobra.Command{
		Use:   "read [<UUID>|<name>]",
		Short: "Display one password of specified item by UUID or name",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SetOut(s.Out)
			cmd.SetErr(s.Err)
			return o.Run(p, cmd, args)
		},
	}

	cmd.SetOutput(s.Out)
	cmd.SetErr(s.Err)
	return cmd
}

// Run runs read command
func (o *Options) Run(p cli.Params, c *cobra.Command, args []string) error {
	// Get Session
	// signinCmd := []string{"op", "signin", "my", "--output=raw"}
	// p.Runner()
	return nil
}
