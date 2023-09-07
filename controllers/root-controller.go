package controllers

import (
	"fmt"
	"github.com/Appkube-awsx/awsx-common/client"
	"github.com/Appkube-awsx/awsx-lambda/services"
	"github.com/aws/aws-sdk-go/service/lambda"
)

// AllLambdaListController for all lambda list
func AllLambdaListController(auth client.Auth) []*lambda.FunctionConfiguration {

	// This is Api auth and compulsory for every controller
	//authenticate.ApiAuth(auth) // No need to call this again. client.Auth is already instantiated at the time of authentication

	// Lambda client from awsx-common repo
	lambdaClient := client.GetClient(auth, client.LAMBDA_CLIENT).(*lambda.Lambda)

	functionList := services.GetAllLambdaList(lambdaClient)
	fmt.Println("List of all lambda functions", functionList)
	return functionList
}

// LambdaListController for pagination lambda list
func LambdaListController(marker string, auth client.Auth) *lambda.ListFunctionsOutput {

	// This is Api auth and compulsory for every controller
	//authenticate.ApiAuth(auth) // No need to call this again. client.Auth is already instantiated at the time of authentication

	lambdaClient := client.GetClient(auth, client.LAMBDA_CLIENT).(*lambda.Lambda)

	functionList := services.GetLambdaList(lambdaClient, marker)

	fmt.Println("List of lambda functions", functionList)
	return functionList
}
