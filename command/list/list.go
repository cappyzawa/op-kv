package list

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/cappyzawa/op-kv/command/util"
	"github.com/spf13/cobra"
)

// Options describes list options
type Options struct {
	outStream io.Writer
	errStream io.Writer
}

// NewOptions initializes list options
func NewOptions(outStream, errStream io.Writer) *Options {
	return &Options{
		outStream: outStream,
		errStream: errStream,
	}
}

// NewCmd initializes list command
func NewCmd(f util.Factory) *cobra.Command {
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

// Run runs list command
func (o *Options) Run(f util.Factory, cmd *cobra.Command, args []string) {
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
