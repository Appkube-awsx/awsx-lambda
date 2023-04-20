package controllers

import (
	"fmt"
	"github.com/Appkube-awsx/awsx-lambda/authenticater"
	"github.com/Appkube-awsx/awsx-lambda/client"
	"github.com/Appkube-awsx/awsx-lambda/services"
	"github.com/aws/aws-sdk-go/service/lambda"
)

func LambdaDetailsController(function string, vaultUrl string, accountId string, region string, acKey string, secKey string, crossAccountRoleArn string, externalId string) *lambda.GetFunctionOutput {

	// this is Api auth and compulsory for every controller
	authenticater.ApiAuth(vaultUrl, accountId, region, acKey, secKey, crossAccountRoleArn, externalId)

	lambdaClient := client.GetClient()
	lambdaDetail := services.GetLambdaDetail(lambdaClient, function)
	fmt.Println(lambdaDetail)
	return lambdaDetail
}
