package controllers

import (
	"fmt"
	"github.com/Appkube-awsx/awsx-common/client"
	"github.com/aws/aws-sdk-go/service/lambda"

	"github.com/Appkube-awsx/awsx-lambda/authenticater"
	"github.com/Appkube-awsx/awsx-lambda/services"
)

func LambdaGetTotalNumberOfLambda(auth client.Auth) int {

	// this is Api auth and compulsory for every controller
	authenticater.ApiAuth(auth)

	// Lambda client
	lambdaClient := client.GetClient(auth, client.LAMBDA_CLIENT).(*lambda.Lambda)

	totalNumber := len(services.GetAllLambdaList(lambdaClient))

	fmt.Println("total number of lambda present in aws account in", authenticater.ClientAuth.Region, "is:", totalNumber)
	return totalNumber
}
