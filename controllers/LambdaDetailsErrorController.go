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

	cloudClient := client.GetCloudWatchClient()

	detail := services.GetFunctionsErrDetail(cloudClient, function)
	fmt.Println(detail)
}
