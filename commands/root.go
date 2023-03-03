/*
Copyright © 2023 Manoj Sharma manoj.sharma@synectiks.com
*/
package commands

import (
	"github.com/Appkube-awsx/awsx-lambda/authenticater"
	"github.com/Appkube-awsx/awsx-lambda/client"
	"github.com/Appkube-awsx/awsx-lambda/commands/lambdacmd"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/spf13/cobra"
	"log"
)

// AwsxCloudElementsCmd represents the base command when called without any subcommands
var AwsxLambdaCmd = &cobra.Command{
	Use:   "lambda",
	Short: "get Lambda Details command gets resource counts",
	Long:  `get Lambda Details command gets resource counts details of an AWS account`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Command lambda started")
		vaultUrl := cmd.PersistentFlags().Lookup("vaultUrl").Value.String()
		accountNo := cmd.PersistentFlags().Lookup("accountId").Value.String()
		region := cmd.PersistentFlags().Lookup("zone").Value.String()
		acKey := cmd.PersistentFlags().Lookup("accessKey").Value.String()
		secKey := cmd.PersistentFlags().Lookup("secretKey").Value.String()
		crossAccountRoleArn := cmd.PersistentFlags().Lookup("crossAccountRoleArn").Value.String()
		externalId := cmd.PersistentFlags().Lookup("externalId").Value.String()

		authFlag := authenticater.AuthenticateData(vaultUrl, accountNo, region, acKey, secKey, crossAccountRoleArn, externalId)

		if authFlag {
			GetLambdaList(region, crossAccountRoleArn, acKey, secKey, externalId)
		}
	},
}

func GetLambdaList(region string, crossAccountRoleArn string, accessKey string, secretKey string, externalId string) *lambda.ListFunctionsOutput {
	log.Println("Getting lambda list summary")
	lambdaClient := client.GetClient(region, crossAccountRoleArn, accessKey, secretKey, externalId)

	input := &lambda.ListFunctionsInput{}
	functionList, err := lambdaClient.ListFunctions(input)
	if err != nil {
		log.Fatalln("Error: in getting lambda list", err)
	}
	log.Println(functionList)
	return functionList
}

//func GetConfig(region string, crossAccountRoleArn string, accessKey string, secretKey string) *configservice.GetDiscoveredResourceCountsOutput {
//	return GetLambdaList(region, crossAccountRoleArn, accessKey, secretKey)
//}

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
	AwsxLambdaCmd.AddCommand(lambdacmd.GetCostDataCmd)
	AwsxLambdaCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxLambdaCmd.PersistentFlags().String("accountId", "", "aws account number")
	AwsxLambdaCmd.PersistentFlags().String("zone", "", "aws region")
	AwsxLambdaCmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxLambdaCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxLambdaCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws cross account role arn")
	AwsxLambdaCmd.PersistentFlags().String("externalId", "", "aws external id auth")
}
