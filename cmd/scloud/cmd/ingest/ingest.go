package ingest

////go:generate scloudgen gen-cmd --name ingest --package ingest --output ingest-gen.go

import (
	usageUtil "github.com/Laplace-Transformer/splunk-cloud-sdk-go/cmd/scloud/util"
	"github.com/spf13/cobra"
)

// Cmd -- used to connection to rootCmd
func Cmd() *cobra.Command {
	return ingestCmd
}

var ingestCmd = &cobra.Command{
	Use:   "ingest",
	Short: "Ingest service",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Usage()
	},
}

func init() {
	ingestCmd.SetUsageTemplate(usageUtil.UsageTemplate)
	ingestCmd.SetHelpTemplate(usageUtil.HelpTemplate)
}
