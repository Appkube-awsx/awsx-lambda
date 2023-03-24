package controllers

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/lambda"
	"log"
)

func GetLambdaDetail(lambdaClient *lambda.Lambda, function string) *lambda.GetFunctionOutput {
	log.Println("Getting Lambda  data")

	input := &lambda.GetFunctionInput{
		FunctionName: aws.String(function),
	}

	lambdaData, err := lambdaClient.GetFunction(input)
	if err != nil {
		log.Fatalln("Error: in getting lambda data", err)
	}

	return lambdaData
}
