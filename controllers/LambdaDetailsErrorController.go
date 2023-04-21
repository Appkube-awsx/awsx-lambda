package controllers

import (
	"fmt"

	"github.com/Appkube-awsx/awsx-lambda/authenticater"
	"github.com/Appkube-awsx/awsx-lambda/client"
	"github.com/Appkube-awsx/awsx-lambda/services"
)

func LambdaDetailsErrorController(function string, vaultUrl string, accountId string, region string, acKey string, secKey string, crossAccountRoleArn string, externalId string) {

	// this is Api auth and compulsory for every controller
	authenticater.ApiAuth(vaultUrl, accountId, region, acKey, secKey, crossAccountRoleArn, externalId)

	lambdaClient := client.GetClient()

	totalNumber := len(services.GetAllLambdaList(lambdaClient))
	fmt.Println("total number of lambda present in aws account in", authenticater.Region, "is:", totalNumber)

}
