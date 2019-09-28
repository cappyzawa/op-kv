package read

import (
	"encoding/json"
	"fmt"

	"github.com/cappyzawa/op-kv/pkg/cli"
	"github.com/cappyzawa/op-kv/pkg/flags"
	"github.com/cappyzawa/op-kv/pkg/helper"
	"github.com/spf13/cobra"
)

// Options describes read options
type Options struct {
	SessionToken *string
	Table        bool
}

// NewOptions initializes read options
func NewOptions() *Options {
	return &Options{}
}

// NewCmd initializes read command
func NewCmd(s *cli.Stream, p cli.Params) *cobra.Command {
	o := NewOptions()
	cmd := &cobra.Command{
		Use:   "read <key>",
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

	cmd.Flags().BoolVarP(&o.Table, "table", "", false, "Print username and password of the item as Table")
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
		helper.RunnerOut(c.OutOrStdout()),
		helper.RunnerErr(c.ErrOrStderr()),
	)
	// Get Item
	opOut, err := runner.Output([]string{"get", "item", item, fmt.Sprintf("--session=%s", *o.SessionToken)})
	if err != nil {
		// op outputs err to stderr
		return fmt.Errorf("failed to execute op command")
	}

	var obj map[string]interface{}
	if err := json.Unmarshal(opOut, &obj); err != nil {
		return err
	}

	fields := obj["details"].(map[string]interface{})["fields"].([]interface{})

	printer := p.Printer(
		helper.PrinterOut(c.OutOrStdout()),
	)
	var username, password string
	for _, f := range fields {
		ff := f.(map[string]interface{})
		name, ok := ff["designation"].(string)
		if !ok {
			continue
		}

		value := ff["value"].(string)
		if name == "password" {
			password = value
		} else if name == "username" {
			username = value
		}
	}

	if username == "" && password == "" {
		return fmt.Errorf("not exist \"%s\"", item)
	}
	if o.Table {
		printer.Header()
		printer.Pair(username, password)
		return nil
	}
	c.Print(password)
	return nil
}
