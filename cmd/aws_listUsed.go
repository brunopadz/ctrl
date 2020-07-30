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
	awsCmd.AddCommand(listUsed)
}

// listAmiCmd represents the listAmi command
var listUsed = &cobra.Command{
	Use:     "list-used",
	Short:   "List used AMIs",
	Long:    `List used AMIs for a given region and account.`,
	Example: `  amictl aws list-used --account 123456789012 --region us-east-1`,
	RunE:    runUsed,
}

func runUsed(cmd *cobra.Command, args []string) error {

	c := color.New(color.FgGreen)

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

	// List used AMIs
	u := providers.AwsListUsed(a, s)

	x := commons.Deduplicate(u)

	r := strings.Join(x, "\n")

	fmt.Println(r)
	c.Println("Total of", len(x), "used AMIs")

	return nil
}
