package controllers

import (
	"github.com/Appkube-awsx/awsx-lambda/authenticater"
	"github.com/Appkube-awsx/awsx-lambda/client"
	"github.com/Appkube-awsx/awsx-lambda/services"
)

func GetLambadaLatencyTimeController(function string, startTime, endTime, vaultUrl string, accountId string, region string, acKey string, secKey string, crossAccountRoleArn string, externalId string) string {

	// this is Api auth and compulsory for every controller
	authenticater.ApiAuth(vaultUrl, accountId, region, acKey, secKey, crossAccountRoleArn, externalId)
	cloudClient := client.GetCloudWatchClient()

	if function != "" {
		result, _, _ := services.GetLambadaLatencyTime(cloudClient, function, startTime, endTime)
		return result
	}

	return "Please send function name"
}
