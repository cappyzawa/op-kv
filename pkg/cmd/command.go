package cmd

import (
	"github.com/cappyzawa/op-kv/pkg/cli"
	"github.com/cappyzawa/op-kv/pkg/cmd/list"
	"github.com/cappyzawa/op-kv/pkg/cmd/read"
	"github.com/cappyzawa/op-kv/pkg/cmd/write"
	"github.com/cappyzawa/op-kv/pkg/flags"
	"github.com/spf13/cobra"
)

// NewCmd initializes command.
func NewCmd(s *cli.Stream) *cobra.Command {
	cmd := &cobra.Command{
		Use:          "op-kv",
		Short:        "use \"op\" like as kv",
		Run:          runHelp,
		SilenceUsage: true,
	}

	flags.AddOpOptions(cmd)

	p := &cli.OpKvParams{}
	cmd.AddCommand(
		read.NewCmd(s, p),
		write.NewCmd(s, p),
		list.NewCmd(s, p),
	)

	return cmd
}

func runHelp(cmd *cobra.Command, args []string) {
	cmd.Help()
}
