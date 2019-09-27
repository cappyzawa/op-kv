package read

import (
	"encoding/json"
	"fmt"

	"github.com/cappyzawa/op-kv/pkg/cli"
	"github.com/cappyzawa/op-kv/pkg/flags"
	"github.com/cappyzawa/op-kv/pkg/helper"
	"github.com/savaki/jq"
	"github.com/spf13/cobra"
)

// Options describes read options
type Options struct {
	SessionToken *string
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
		PreRunE: func(cmd *cobra.Command, args []string) error {
			var err error
			o.SessionToken, err = p.Runner().Signin(flags.SubDomain, flags.OpPassword)
			if err != nil {
				return err
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return o.Run(p, cmd, args)
		},
		PostRun: func(cmd *cobra.Command, args []string) {
			// flush
			o.SessionToken = nil
		},
	}

	cmd.SetOutput(s.Out)
	cmd.SetErr(s.Err)
	return cmd
}

// Run runs read command
func (o *Options) Run(p cli.Params, c *cobra.Command, args []string) error {
	if len(args) != 1 {
		c.Help()
		return fmt.Errorf("see Usage")
	}
	item := args[0]
	runner := p.Runner(
		helper.Out(c.OutOrStdout()),
		helper.Err(c.ErrOrStderr()),
	)
	// Get Item
	opOut, err := runner.Output([]string{"get", "item", item, fmt.Sprintf("--session=%s", *o.SessionToken)})
	if err != nil {
		// op outputs err to stderr
		return fmt.Errorf("failed to execute op command")
	}

	filter, err := jq.Parse(".details.fields")
	if err != nil {
		return err
	}
	filtered, err := filter.Apply(opOut)
	if err != nil {
		return fmt.Errorf("failed to fileter by jq: %v", err)
	}

	var obj []map[string]string
	if err := json.Unmarshal(filtered, &obj); err != nil {
		return err
	}

	for _, o := range obj {
		name, ok := o["name"]
		if !ok || name != "password" {
			continue
		}
		value := o["value"]
		c.Printf(value)
		return nil
	}
	return fmt.Errorf("not exist %s", item)
}
