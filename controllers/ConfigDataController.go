package controllers

import (
	"fmt"
	"github.com/Appkube-awsx/awsx-common/client"
	"github.com/Appkube-awsx/awsx-lambda/services"
	"github.com/aws/aws-sdk-go/service/lambda"
)

func LambdaDetails(function string, auth client.Auth) *lambda.GetFunctionOutput {

	// this is Api auth and compulsory for every controller
	//authenticater.ApiAuth(auth) // No need to call this again. client.Auth is already instantiated at the time of authentication

	// Lambda client
	lambdaClient := client.GetClient(auth, client.LAMBDA_CLIENT).(*lambda.Lambda)

	lambdaDetail := services.GetLambdaDetail(lambdaClient, function)
	fmt.Println(lambdaDetail)
	return lambdaDetail
}
