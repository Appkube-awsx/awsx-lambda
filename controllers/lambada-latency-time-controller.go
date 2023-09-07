package controllers

import (
	"github.com/Appkube-awsx/awsx-common/client"
	"github.com/Appkube-awsx/awsx-lambda/services"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

func GetLambadaLatencyTimeController(function string, startTime string, endTime string, auth client.Auth) string {

	// this is Api auth and compulsory for every controller
	//authenticate.ApiAuth(auth) // No need to call this again. client.Auth is already instantiated at the time of authentication

	cloudClient := client.GetClient(auth, client.CLOUDWATCH_LOG).(*cloudwatchlogs.CloudWatchLogs)

	if function != "" {
		result, _, _ := services.GetLambadaLatencyTime(cloudClient, function, startTime, endTime)
		return result
	}

	return "Please send function name"
}
