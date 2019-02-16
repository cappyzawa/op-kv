package write

import (
	"github.com/cappyzawa/op-kv/command/util"
	"github.com/spf13/cobra"
)

type options struct {
}

func NewOptions() *options {
	return &options{}
}

func NewCmdWrite(f util.Factory) *cobra.Command {
	o := NewOptions()
	cmd := &cobra.Command{
		Use:   "write <item> <password>",
		Short: "Generate one password by specified item and password",
		RunE: func(cmd *cobra.Command, args []string) error {
			return o.Run(f, cmd, args)
		},
	}
	return cmd
}

func (o *options) Run(f util.Factory, cmd *cobra.Command, args []string) error {
	return nil
}
