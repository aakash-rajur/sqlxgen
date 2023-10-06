package cli

import (
	"github.com/spf13/cobra"
)

func versionCmd(version string) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "version",
		Aliases: []string{"v"},
		Short:   "Print the version number of sqlxgen",
		Run: func(_ *cobra.Command, _ []string) {
			println(version)
		},
	}

	return cmd
}
