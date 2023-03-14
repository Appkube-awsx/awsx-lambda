/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package lambdacmd

import (
	"fmt"
	"github.com/Appkube-awsx/awsx-lambda/authenticater"
	"github.com/Appkube-awsx/awsx-lambda/client"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/spf13/cobra"
	"log"
)

// GetTotalNumberOfLambdaCmd represents the number command
var GetTotalNumberOfLambdaCmd = &cobra.Command{
	Use:   "totalCount",
	Short: "gets total number of lambdas present in aws account",
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
			totalNumber := GetLambdaNumbers(region, crossAccountRoleArn, acKey, secKey, externalId)
			fmt.Println("total number of lambda present in aws account in", region, "is:", totalNumber)
		}
	},
}

func GetLambdaNumbers(region string, crossAccountRoleArn string, accessKey string, secretKey string, externalId string) int {
	log.Println("Getting number of lambdas...")
	lambdaClient := client.GetClient(region, crossAccountRoleArn, accessKey, secretKey, externalId)

	// get first list of function
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

	return len(allFunctions)
}

func init() {

}
