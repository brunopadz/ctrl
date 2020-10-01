package cmd

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/brunopadz/amictl/commons"
	"github.com/brunopadz/amictl/providers"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	awsCmd.AddCommand(listUnused)
}

// listAmiCmd represents the listAmi command
var listUnused = &cobra.Command{
	Use:     "list-unused",
	Short:   "List unused AMIs",
	Long:    `List not used AMIs for a given region and account.`,
	Example: `  amictl aws list-unused --account 123456789012 --region us-east-1`,
	RunE:    runUnused,
}

func runUnused(cmd *cobra.Command, args []string) error {

	red := color.New(color.FgRed)
	yellow := color.New(color.FgYellow)
	green := color.New(color.FgGreen)

	// Creates a input filter to get AMIs
	f := &ec2.DescribeImagesInput{
		Owners: []*string{
			aws.String(account),
		},
	}

	// Establishes new authenticated session to AWS
	s := providers.AwsSession(region)

	// Filter AMIs based on input filter
	a, err := s.DescribeImages(f)
	if err != nil {
		fmt.Println(err)
	}

	// Compare AMI list
	l, u := providers.AwsListNotUsed(a, s)

	n := commons.Compare(l, u)
	r := strings.Join(n, "\n")

	fmt.Println(r)

	if len(n) == 0 {
		green.Println("Yay! You're already saving some money. 🎉")
	} else if len(n) >= 1 && len(n) <= 10 {
		yellow.Println("There are a total of", len(n), "not used AMIs. You could be saving some money.")
	} else {
		red.Println("There are a total of", len(n), "not used AMIs. Go ahead and delete them to save some money.")
	}

	return nil
}
