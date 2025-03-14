package action

////go:generate scloudgen gen-cmd --name action --package action --output action-gen.go

import (
	usageUtil "github.com/Laplace-Transformer/splunk-cloud-sdk-go/cmd/scloud/util"
	"github.com/spf13/cobra"
)

// Cmd -- used to connection to rootCmd
func Cmd() *cobra.Command {
	return actionCmd
}

// actionCmd represents the Action command
var actionCmd = &cobra.Command{
	Use:   "action",
	Short: "Action service",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Usage()
	},
}

func init() {
	actionCmd.SetUsageTemplate(usageUtil.UsageTemplate)
	actionCmd.SetHelpTemplate(usageUtil.HelpTemplate)
}
