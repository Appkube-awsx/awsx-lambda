package controllers

import (
	"fmt"
	"github.com/Appkube-awsx/awsx-common/client"
	"github.com/aws/aws-sdk-go/service/lambda"

	"github.com/Appkube-awsx/awsx-lambda/services"
)

func LambdaGetTotalNumberOfLambda(auth client.Auth) int {

	// this is Api auth and compulsory for every controller
	//authenticate.ApiAuth(auth) // No need to call this again. client.Auth is already instantiated at the time of authentication

	// Lambda client
	lambdaClient := client.GetClient(auth, client.LAMBDA_CLIENT).(*lambda.Lambda)

	totalNumber := len(services.GetAllLambdaList(lambdaClient))

	//fmt.Println("total number of lambda present in aws account in", auth.Region, "is:", totalNumber)
	fmt.Println("total lambda: ", totalNumber, ", region: ", auth.Region, ", account no: ", auth.CrossAccountRoleArn)
	return totalNumber
}
