package streams

////go:generate scloudgen gen-cmd --name streams --package streams --output streams-gen.go

import (
	usageUtil "github.com/Laplace-Transformer/splunk-cloud-sdk-go/cmd/scloud/util"
	"github.com/spf13/cobra"
)

// Cmd -- used to connection to rootCmd
func Cmd() *cobra.Command {
	return streamsCmd
}

// streamsCmd represents the streams command
var streamsCmd = &cobra.Command{
	Use:   "streams",
	Short: "Streams service",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Usage()
	},
}

func init() {
	streamsCmd.SetUsageTemplate(usageUtil.UsageTemplate)
	streamsCmd.SetHelpTemplate(usageUtil.HelpTemplate)
}
