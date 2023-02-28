/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package lambdacmd

import (
	"fmt"
	"github.com/Appkube-awsx/awsx-lambda/authenticater"
	"github.com/Appkube-awsx/awsx-lambda/client"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/spf13/cobra"
	"log"
)

// getConfigDataCmd represents the getConfigData command
var GetConfigDataCmd = &cobra.Command{
	Use:   "getConfigData",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		vaultUrl := cmd.Parent().PersistentFlags().Lookup("vaultUrl").Value.String()
		accountNo := cmd.Parent().PersistentFlags().Lookup("accountId").Value.String()
		region := cmd.Parent().PersistentFlags().Lookup("zone").Value.String()
		acKey := cmd.Parent().PersistentFlags().Lookup("accessKey").Value.String()
		secKey := cmd.Parent().PersistentFlags().Lookup("secretKey").Value.String()
		crossAccountRoleArn := cmd.Parent().PersistentFlags().Lookup("crossAccountRoleArn").Value.String()
		externalId := cmd.Parent().PersistentFlags().Lookup("externalId").Value.String()

		authFlag := authenticater.AuthenticateData(vaultUrl, accountNo, region, acKey, secKey, crossAccountRoleArn, externalId)

		if authFlag {
			function, _ := cmd.Flags().GetString("func")
			getLambdaDetail(region, crossAccountRoleArn, acKey, secKey, function, externalId)
		}
	},
}

func getLambdaDetail(region string, crossAccountRoleArn string, accessKey string, secretKey string, function string, externalId string) (*lambda.FunctionConfiguration, error) {
	log.Println("Getting Lambda  data")
	lambdaClient := client.GetClient(region, crossAccountRoleArn, accessKey, secretKey, externalId)

	input := &lambda.GetFunctionConfigurationInput{
		FunctionName: aws.String(function),
	}

	lambdaData, err := lambdaClient.GetFunctionConfiguration(input)
	if err != nil {
		log.Fatalln("Error: in getting lambda data", err)
	}

	log.Println(lambdaData)
	return lambdaData, err
}

func init() {
	GetConfigDataCmd.Flags().StringP("func", "f", "", "lambda function name")

	if err := GetConfigDataCmd.MarkFlagRequired("func"); err != nil {
		fmt.Println("--func is required", err)
	}
}
