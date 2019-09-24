package command

import (
	"flag"

	"github.com/cappyzawa/op-kv/command/list"

	"github.com/cappyzawa/op-kv/command/util"

	"github.com/cappyzawa/op-kv/command/read"
	"github.com/cappyzawa/op-kv/command/write"
	"github.com/spf13/cobra"
)

// NewCmd initializes command.
func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "op-kv",
		Short: "use \"op\" like as kv",
		Run:   runHelp,
	}

	cmd.PersistentFlags().AddGoFlagSet(flag.CommandLine)

	f := util.NewFactory()

	cmd.AddCommand(
		read.NewCmd(f),
		write.NewCmd(f),
		list.NewCmd(f),
	)

	return cmd
}

func runHelp(cmd *cobra.Command, args []string) {
	cmd.Help()
}
