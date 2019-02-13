package write

import "github.com/spf13/cobra"

type options struct {
}

func NewOptions() *options {
	return &options{}
}

func NewCmdWrite() *cobra.Command {
	o := NewOptions()
	cmd := &cobra.Command{
		Use:   "write <item> <password>",
		Short: "Generate one password by specified item and password",
		RunE: func(cmd *cobra.Command, args []string) error {
			return o.Run(cmd, args)
		},
	}
	return cmd
}

func (o *options) Run(cmd *cobra.Command, args []string) error {
	return nil
}
