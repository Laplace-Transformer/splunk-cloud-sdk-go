package forwarders

////go:generate scloudgen gen-cmd --name forwarders --package forwarders --output forwarders-gen.go

import (
	usageUtil "github.com/Laplace-Transformer/splunk-cloud-sdk-go/cmd/scloud/util"
	"github.com/spf13/cobra"
)

// Cmd -- used to connection to rootCmd
func Cmd() *cobra.Command {
	return forwardersCmd
}

var forwardersCmd = &cobra.Command{
	Use:   "forwarders",
	Short: "Forwarders service",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Usage()
	},
}

func init() {
	forwardersCmd.SetUsageTemplate(usageUtil.UsageTemplate)
	forwardersCmd.SetHelpTemplate(usageUtil.HelpTemplate)
}
