package ml

////go:generate scloudgen gen-cmd --name ml --package ml --output ml-gen.go | gofmt

import (
	usageUtil "github.com/Laplace-Transformer/splunk-cloud-sdk-go/cmd/scloud/util"
	"github.com/spf13/cobra"
)

// Cmd -- used to connection to rootCmd
func Cmd() *cobra.Command {
	return mlCmd
}

var mlCmd = &cobra.Command{
	Use:   "ml",
	Short: "Machine Learning service",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Usage()
	},
}

func init() {
	mlCmd.SetUsageTemplate(usageUtil.UsageTemplate)
	mlCmd.SetHelpTemplate(usageUtil.HelpTemplate)
}
