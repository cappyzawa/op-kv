package flags

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	// OpPassword is password of op
	OpPassword string
)

// AddOpOptions amends command to add flags
func AddOpOptions(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(
		&OpPassword, "password", "p", os.Getenv("OP_PASSWORD"),
		"password for 1password",
	)
}
