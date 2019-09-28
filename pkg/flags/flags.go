package flags

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	// OpPassword is password of op
	OpPassword string
	// SubDomain is subdomain of 1password.com(e.g. my)
	SubDomain string
)

// AddOpOptions amends command to add flags
func AddOpOptions(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(
		&OpPassword, "op-password", "op", os.Getenv("OP_PASSWORD"),
		"password for 1password",
	)

	subdomain, err := getSubDomain()
	if err != nil {
		fmt.Println(err)
	}
	cmd.PersistentFlags().StringVarP(
		&SubDomain, "subdomain", "d", subdomain,
		"subdomain of 1password",
	)
}

func getSubDomain() (string, error) {
	baseDir := os.Getenv("XDG_CONFIG_HOME")
	if baseDir == "" {
		baseDir = os.Getenv("HOME")
	}
	configPath := filepath.Join(baseDir, ".op", "config")
	config, err := ioutil.ReadFile(configPath)
	if err != nil {
		return "", err
	}
	var obj map[string]interface{}
	if err := json.Unmarshal(config, &obj); err != nil {
		return "", err
	}
	subdomain, ok := obj["latest_signin"].(string)
	if !ok {
		return "", fmt.Errorf("you shoud signin to 1passord once")
	}
	return subdomain, nil
}
