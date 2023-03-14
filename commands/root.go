/*
Copyright © 2023 Manoj Sharma manoj.sharma@synectiks.com
*/
package commands

import (
	"fmt"
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
		marker := cmd.Flags().Lookup("marker").Value.String()
		all, _ := cmd.Flags().GetBool("all")

		authFlag := authenticater.AuthenticateData(vaultUrl, accountNo, region, acKey, secKey, crossAccountRoleArn, externalId)

		if authFlag {

			if all {
				functionList := GetAllLambdaList(region, crossAccountRoleArn, acKey, secKey, externalId)
				fmt.Println("List of all lambda functions", functionList)
			} else {
				functionList := GetLambdaList(region, crossAccountRoleArn, acKey, secKey, externalId, marker)
				fmt.Println("List of by marker lambda functions", functionList)
			}
		}
	},
}

func GetLambdaList(region string, crossAccountRoleArn string, accessKey string, secretKey string, externalId string, marker string) *lambda.ListFunctionsOutput {
	log.Println("Getting lambda list summary")
	lambdaClient := client.GetClient(region, crossAccountRoleArn, accessKey, secretKey, externalId)

	input := &lambda.ListFunctionsInput{}

	if marker != "" {
		input = &lambda.ListFunctionsInput{
			Marker: &marker,
		}
	}

	functionList, err := lambdaClient.ListFunctions(input)
	if err != nil {
		log.Fatalln("Error: in getting lambda list", err)
	}

	return functionList

}

func GetAllLambdaList(region string, crossAccountRoleArn string, accessKey string, secretKey string, externalId string) []*lambda.FunctionConfiguration {
	log.Println("Getting lambda list summary")
	lambdaClient := client.GetClient(region, crossAccountRoleArn, accessKey, secretKey, externalId)

	input := &lambda.ListFunctionsInput{}
	functionList, err := lambdaClient.ListFunctions(input)
	if err != nil {
		log.Fatalln("Error: in getting total number of lambdas", err)
	}

	allFunctions := functionList.Functions
	marker := functionList.NextMarker

	// Loop for getting all lambdas
	for marker != nil {
		input = &lambda.ListFunctionsInput{
			Marker: marker,
		}
		functionList, err = lambdaClient.ListFunctions(input)
		if err != nil {
			log.Fatalln("Error: in getting lambda numbers", err)
		}
		allFunctions = append(allFunctions, functionList.Functions...)
		marker = functionList.NextMarker
		fmt.Println("Functions got till now:: ", len(allFunctions))
	}

	return allFunctions
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
	AwsxLambdaCmd.AddCommand(lambdacmd.GetNumberOfErrorCmd)
	AwsxLambdaCmd.AddCommand(lambdacmd.GetTotalNumberOfLambdaCmd)
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
