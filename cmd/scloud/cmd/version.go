package cmd

import (
	"fmt"

	"github.com/Laplace-Transformer/splunk-cloud-sdk-go/cmd/scloud/cmd/scloud/version"
	"github.com/Laplace-Transformer/splunk-cloud-sdk-go/cmd/scloud/jsonx"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Version of Splunk Cloud Services CLI",
	RunE:  execVersionCmd,
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func execVersionCmd(cmd *cobra.Command, args []string) error {

	err := fmt.Sprintf("scloud version %s\n", version.Version)
	jsonx.Pprint(cmd, err)
	return nil
}
