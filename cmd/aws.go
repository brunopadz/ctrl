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
 - List all AMIs -> amictl aws list-all [ACCOUNT_ID] [REGION]
 - List unused AMIs -> amictl aws list-unused [ACCOUNT_ID] [REGION]`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.Long)
	},
}

func init() {
	rootCmd.AddCommand(awsCmd)
}
