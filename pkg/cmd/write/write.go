package write

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/cappyzawa/op-kv/pkg/cli"
	"github.com/cappyzawa/op-kv/pkg/flags"
	"github.com/cappyzawa/op-kv/pkg/helper"
	"github.com/spf13/cobra"
)

// Options describes write options
type Options struct {
	SessionToken *string
}

// NewOptions initializes write options
func NewOptions() *Options {
	return &Options{}
}

// NewCmd initializes write command
func NewCmd(s *cli.Stream, p cli.Params) *cobra.Command {
	o := NewOptions()
	cmd := &cobra.Command{
		Use:   "write <key> <value>",
		Short: "Generate one password by specified item and password",
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

// Run runs write command
func (o *Options) Run(p cli.Params, c *cobra.Command, args []string) error {
	if len(args) != 2 {
		c.Help()
		return fmt.Errorf("see Usage")
	}
	key := args[0]
	value := args[1]

	runner := p.Runner(
		helper.RunnerOut(c.OutOrStdout()),
		helper.RunnerErr(c.ErrOrStderr()),
	)

	sessionFlag := fmt.Sprintf("--session=%s", *o.SessionToken)

	opOut, err := runner.Output([]string{"get", "template", "login", sessionFlag})
	if err != nil {
		// op outputs err to stderr
		return fmt.Errorf("failed to execute op command")
	}
	var obj map[string]interface{}
	if err := json.Unmarshal(opOut, &obj); err != nil {
		return err
	}

	fields := obj["fields"].([]interface{})
	for _, f := range fields {
		m := f.(map[string]interface{})
		if m["name"].(string) == "password" {
			m["value"] = value
		}
	}

	template, err := json.Marshal(&obj)
	if err != nil {
		return err
	}

	opEncoded, err := runner.OutputWithIn([]string{"encode"}, string(template))
	if err != nil {
		return err
	}

	if _, err := runner.Output([]string{"create", "item", "login", strings.TrimSpace(string(opEncoded)), fmt.Sprintf("--title=%s", key), sessionFlag}); err != nil {
		return err
	}

	c.Printf("success to write password to \"%s\"\n", key)
	return nil
}
