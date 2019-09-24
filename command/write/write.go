package write

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

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

func NewCmdWrite(f util.Factory) *cobra.Command {
	o := NewOptions(os.Stdout, new(bytes.Buffer))
	cmd := &cobra.Command{
		Use:   "write <item> <password>",
		Short: "Generate one password by specified item and password",
		Run: func(cmd *cobra.Command, args []string) {
			o.Run(f, cmd, args)
		},
	}
	return cmd
}

func (o *options) Run(f util.Factory, cmd *cobra.Command, args []string) {
	item := args[0]
	password := args[1]

	runner := f.CommandRunner()

	opTmpCmd := []string{"op", "get", "template", "login"}
	jqCmd := []string{"jq", "-c", fmt.Sprintf(".fields[1].value = \"%s\"", password)}
	opEncCmd := []string{"op", "encode"}
	output, err := runner.Output(opTmpCmd, jqCmd, opEncCmd)
	if err != nil {
		fmt.Fprint(o.errStream, err.Error())
		return
	}

	opCmd := []string{"op", "create", "item", "login", strings.TrimRight(string(output), "\n"), fmt.Sprintf("--title=%s", item)}
	if _, err := runner.Output(opCmd); err != nil {
		fmt.Fprint(o.errStream, err.Error())
		return
	}
	fmt.Fprint(o.outStream, "success to write password!!")
	return
}