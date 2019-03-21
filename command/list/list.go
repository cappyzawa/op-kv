package list

import (
	"bytes"
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

func NewCmdList(f util.Factory) *cobra.Command {
	o := NewOptions(os.Stdout, new(bytes.Buffer))
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Display item titles",
		Run: func(cmd *cobra.Command, args []string) {
			o.Run(f, cmd, args)
		},
	}
	return cmd
}

func (o *options) Run(f util.Factory, cmd *cobra.Command, args []string) {
	runner := f.CommandRunner()

	// op list items | jq -r ".[].overview.title"
	opCmd := []string{"op", "list", "items"}
	jqCmd := []string{"jq", "-r", ".[].overview.title"}

	output, err := runner.Output(opCmd, jqCmd)
	if err != nil {
		fmt.Fprint(o.errStream, err)
		return
	}

	fmt.Fprintf(o.outStream, "%s", string(output))
}
