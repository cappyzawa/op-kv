package read

import (
	"github.com/cappyzawa/op-kv/command/util"
	"github.com/spf13/cobra"
)

type options struct {
}

func NewOptions() *options {
	return &options{}
}

func NewCmdRead(f util.Factory) *cobra.Command {
	o := NewOptions()
	cmd := &cobra.Command{
		Use:   "read [<UUID>|<name>]",
		Short: "Display one password of specified item by UUID or name",
		RunE: func(cmd *cobra.Command, args []string) error {
			return o.Run(f, cmd, args)
		},
	}
	return cmd
}

func (o *options) Run(f util.Factory, cmd *cobra.Command, args []string) error {
	if len(args) > 0 {
		cmd.Help()
	}

	return nil
}
