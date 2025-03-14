package kvstore

////go:generate scloudgen gen-cmd --name action --package action --output action-gen.go

import (
	usageUtil "github.com/Laplace-Transformer/splunk-cloud-sdk-go/cmd/scloud/util"
	"github.com/spf13/cobra"
)

// Cmd -- used to connection to rootCmd
func Cmd() *cobra.Command {
	return kvstoreCmd
}

// kvstoreCmd represents the KVStore command
var kvstoreCmd = &cobra.Command{
	Use:   "kvstore",
	Short: "KV Store service",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Usage()
	},
}

func init() {
	kvstoreCmd.SetUsageTemplate(usageUtil.UsageTemplate)
	kvstoreCmd.SetHelpTemplate(usageUtil.HelpTemplate)
}
