package collect

////go:generate scloudgen gen-cmd --name collect --package collect --output collect-gen.go

import (
	usageUtil "github.com/Laplace-Transformer/splunk-cloud-sdk-go/cmd/scloud/util"
	"github.com/spf13/cobra"
)

// Cmd -- used to connection to rootCmd
func Cmd() *cobra.Command {
	return collectCmd
}

// collectCmd represents the collect command
var collectCmd = &cobra.Command{
	Use:   "collect",
	Short: "Collect service",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Usage()
	},
}

func init() {
	collectCmd.SetUsageTemplate(usageUtil.UsageTemplate)
	collectCmd.SetHelpTemplate(usageUtil.HelpTemplate)
}
