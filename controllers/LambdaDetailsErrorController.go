package controllers

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"

	"github.com/Appkube-awsx/awsx-common/client"
	"github.com/Appkube-awsx/awsx-lambda/services"
)

func LambdaDetailsErrorController(function string, auth client.Auth) {

	// this is Api auth and compulsory for every controller
	//authenticater.ApiAuth(auth) // No need to call this again. client.Auth is already instantiated at the time of authentication

	// Cloud client
	cloudClient := client.GetClient(auth, client.CLOUDWATCH_LOG).(*cloudwatchlogs.CloudWatchLogs)

	detail := services.GetFunctionsErrDetail(cloudClient, function)
	fmt.Println(detail)
}
