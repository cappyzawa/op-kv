package list

import (
	"encoding/json"
	"fmt"

	"github.com/cappyzawa/op-kv/pkg/cli"
	"github.com/cappyzawa/op-kv/pkg/flags"
	"github.com/cappyzawa/op-kv/pkg/helper"
	"github.com/spf13/cobra"
)

// Options describes list options
type Options struct {
	SessionToken *string
}

// NewOptions initializes list options
func NewOptions() *Options {
	return &Options{}
}

// NewCmd initializes list command
func NewCmd(s *cli.Stream, p cli.Params) *cobra.Command {
	o := NewOptions()
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Display item titles",
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

// Run runs list command
func (o *Options) Run(p cli.Params, c *cobra.Command, args []string) error {
	runner := p.Runner(
		helper.RunnerOut(c.OutOrStdout()),
		helper.RunnerErr(c.ErrOrStderr()),
	)

	opOut, err := runner.Output([]string{"list", "items", fmt.Sprintf("--session=%s", *o.SessionToken)})
	if err != nil {
		// op outputs err to stderr
		return fmt.Errorf("failed to execute op command")
	}

	var obj []interface{}
	if err := json.Unmarshal(opOut, &obj); err != nil {
		return err
	}

	for _, o := range obj {
		overview, ok := o.(map[string]interface{})["overview"].(map[string]interface{})
		if !ok {
			continue
		}
		title, ok := overview["title"].(string)
		if !ok {
			continue
		}
		c.Println(title)
	}
	return nil
}
