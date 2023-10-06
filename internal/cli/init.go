package cli

import (
	"os"

	i "github.com/aakash-rajur/sqlxgen/internal/init"
	"github.com/aakash-rajur/sqlxgen/internal/utils"
	"github.com/spf13/cobra"
)

func initCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "init",
		Aliases: []string{"i"},
		Short:   "Initialize a sqlxgen config",
		Run: func(_ *cobra.Command, _ []string) {
			err := runInit()

			if err == nil {
				return
			}

			utils.ExitWithError(err)
		},
	}

	return cmd
}

func runInit() error {
	workDir, err := os.Getwd()

	if err != nil {
		return err
	}

	return i.Init(workDir)
}
