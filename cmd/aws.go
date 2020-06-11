package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var awsCmd = &cobra.Command{
	Use:   "aws",
	Short: "Manages AWS AMIs",
	Long: `amictl aws - Manages AWS AMIs

With amictl aws it's possible to:
 - List all AMIs
 - List unused AMIs`,
	Run: WelcomeCommand(),
}

func init() {
	rootCmd.AddCommand(awsCmd)
}

func WelcomeCommand() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.Long)
	}
}
