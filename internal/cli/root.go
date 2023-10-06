package cli

import (
	"github.com/spf13/cobra"
)

func RootCmd(version string) *cobra.Command {
	root := &cobra.Command{
		Use:   "sqlxgen",
		Short: "sqlxgen is a code generator for sqlx",
		Long: `sqlxgen is a code generator for sqlx, with support for both postgres and mysql.
1. introspect database schema and generate models
2. introspect sql queries in your project and generates input and output models
`,
	}

	root.AddCommand(initCmd())

	root.AddCommand(generateCmd())

	root.AddCommand(versionCmd(version))

	return root
}
