package controllers

import (
	"fmt"
	"github.com/Appkube-awsx/awsx-lambda/authenticater"
	"github.com/Appkube-awsx/awsx-lambda/client"
	"github.com/Appkube-awsx/awsx-lambda/services"
	"github.com/aws/aws-sdk-go/service/lambda"
)

// Controller for all lambda list
func AllLambdaListController(vaultUrl string, accountId string, region string, acKey string, secKey string, crossAccountRoleArn string, externalId string) []*lambda.FunctionConfiguration {

	// this is Api auth and compulsory for every controller
	authenticater.ApiAuth(vaultUrl, accountId, region, acKey, secKey, crossAccountRoleArn, externalId)

	lambdaClient := client.GetClient()
	functionList := services.GetAllLambdaList(lambdaClient)
	fmt.Println("List of all lambda functions", functionList)
	return functionList
}

// Controller for pagination lambda list
func LambdaListController(marker string, vaultUrl string, accountId string, region string, acKey string, secKey string, crossAccountRoleArn string, externalId string) *lambda.ListFunctionsOutput {

	authenticater.ApiAuth(vaultUrl, accountId, region, acKey, secKey, crossAccountRoleArn, externalId)

	lambdaClient := client.GetClient()

	functionList := services.GetLambdaList(lambdaClient, marker)

	fmt.Println("List of lambda functions", functionList)
	return functionList
}
