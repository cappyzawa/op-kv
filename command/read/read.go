package read

import "github.com/spf13/cobra"

type options struct {
}

func NewOptions() *options {
	return &options{}
}

func NewCmdRead() *cobra.Command {
	o := NewOptions()
	cmd := &cobra.Command{
		Use:   "read [<UUID>|<name>]",
		Short: "Display one password of specified item by UUID or name",
		RunE: func(cmd *cobra.Command, args []string) error {
			return o.Run(cmd, args)
		},
	}
	return cmd
}

func (o *options) Run(cmd *cobra.Command, args []string) error {
	return nil
}
