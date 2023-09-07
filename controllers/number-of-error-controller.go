package controllers

import (
	"github.com/Appkube-awsx/awsx-common/client"
	"github.com/Appkube-awsx/awsx-lambda/services"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

func LambdaGetNumberOfErrorController(function string, auth client.Auth) int {

	// this is Api auth and compulsory for every controller
	//authenticate.ApiAuth(auth) // No need to call this again. client.Auth is already instantiated at the time of authentication

	cloudClient := client.GetClient(auth, client.CLOUDWATCH_LOG).(*cloudwatchlogs.CloudWatchLogs)

	services.GetFunctionErrCount(cloudClient, function)
	return -1
}
