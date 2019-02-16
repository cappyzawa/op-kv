package read

import (
	"fmt"
	"io"
	"os"

	"github.com/cappyzawa/op-kv/command/util"
	"github.com/spf13/cobra"
)

type options struct {
	outStream io.Writer
	errStream io.Writer
}

func NewOptions(outStream, errStream io.Writer) *options {
	return &options{
		outStream: outStream,
		errStream: errStream,
	}
}

func NewCmdRead(f util.Factory) *cobra.Command {
	o := NewOptions(os.Stdout, os.Stderr)
	cmd := &cobra.Command{
		Use:   "read [<UUID>|<name>]",
		Short: "Display one password of specified item by UUID or name",
		RunE: func(cmd *cobra.Command, args []string) error {
			return o.Run(f, cmd, args)
		},
	}
	return cmd
}

func (o *options) Run(f util.Factory, cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		cmd.Help()
		return nil
	}

	item := args[0]
	runner := f.CommandRunner()

	// op get item GitHub | jq -r '.details.fields[] | select(.designation=="password").value'
	opCmd := []string{"op", "get", "item", item}
	jqCmd := []string{"jq", "-r", ".details.fields[] | select(.designation==\"password\").value"}

	output, err := runner.Output(opCmd, jqCmd)
	if err != nil {
		fmt.Fprint(o.errStream, err)
		return err
	}

	fmt.Fprintf(o.outStream, "%s", string(output))
	return nil
}
