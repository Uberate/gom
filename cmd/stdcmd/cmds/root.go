package cmds

import "github.com/spf13/cobra"

func RootCmd() *cobra.Command {
	cmd := &cobra.Command{}

	cmd.AddCommand(RegexCmd())

	return cmd
}
