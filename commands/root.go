/*
Copyright © 2023 Manoj Sharma manoj.sharma@synectiks.com
*/
package commands

import (
	"log"

	"github.com/Appkube-awsx/awsx-lambda/authenticater"
	"github.com/Appkube-awsx/awsx-lambda/commands/lambdacmd"
	"github.com/Appkube-awsx/awsx-lambda/controllers"
	"github.com/spf13/cobra"
)

// AwsxCloudElementsCmd represents the base command when called without any subcommands
var AwsxLambdaCmd = &cobra.Command{
	Use:   "lambda",
	Short: "get Lambda Details command gets resource counts",
	Long:  `get Lambda Details command gets resource counts details of an AWS account`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Command lambda started")

		// check for cli flags
		authFlag := authenticater.RootCommandAuth(cmd)

		marker := cmd.Flags().Lookup("marker").Value.String()
		all, _ := cmd.Flags().GetBool("all")

		if authFlag {
			if all {
				controllers.AllLambdaListController(authenticater.VaultUrl, authenticater.AccountId, authenticater.Region, authenticater.AcKey, authenticater.SecKey, authenticater.CrossAccountRoleArn, authenticater.ExternalId)
			} else {
				controllers.LambdaListController(marker, authenticater.VaultUrl, authenticater.AccountId, authenticater.Region, authenticater.AcKey, authenticater.SecKey, authenticater.CrossAccountRoleArn, authenticater.ExternalId)
			}
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := AwsxLambdaCmd.Execute()
	if err != nil {
		log.Fatal("There was some error while executing the CLI: ", err)
		return
	}
}

func init() {
	AwsxLambdaCmd.AddCommand(lambdacmd.GetConfigDataCmd)
	AwsxLambdaCmd.AddCommand(lambdacmd.GetNumberOfErrorCmd)
	AwsxLambdaCmd.AddCommand(lambdacmd.GetTotalNumberOfLambdaCmd)
	AwsxLambdaCmd.AddCommand(lambdacmd.GetDetailOfErrorCmd)
	AwsxLambdaCmd.AddCommand(lambdacmd.GetLatencyCmd)

	AwsxLambdaCmd.Flags().String("marker", "", "marker for next list")
	AwsxLambdaCmd.Flags().Bool("all", false, "to get all lambdas at once")

	AwsxLambdaCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxLambdaCmd.PersistentFlags().String("accountId", "", "aws account number")
	AwsxLambdaCmd.PersistentFlags().String("zone", "", "aws region")
	AwsxLambdaCmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxLambdaCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxLambdaCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws cross account role arn")
	AwsxLambdaCmd.PersistentFlags().String("externalId", "", "aws external id auth")
}
