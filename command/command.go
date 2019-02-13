package command

import (
	"flag"

	"github.com/cappyzawa/op-kv/command/read"
	"github.com/cappyzawa/op-kv/command/write"
	"github.com/spf13/cobra"
)

type Group struct {
	Message  string
	Commands []*cobra.Command
}

func (cg Group) Add(c *cobra.Command) {
	c.AddCommand(cg.Commands...)
}

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "op-kv",
		Short: "use \"op\" like as kv",
		Run:   runHelp,
	}

	cmd.PersistentFlags().AddGoFlagSet(flag.CommandLine)
	group := Group{
		Message: "Basic Commands",
		Commands: []*cobra.Command{
			read.NewCmdRead(),
			write.NewCmdWrite(),
		},
	}
	group.Add(cmd)
	return cmd
}

func runHelp(cmd *cobra.Command, args []string) {
	cmd.Help()
}
