package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/Appkube-awsx/awsx-common/client"
	"github.com/Appkube-awsx/awsx-lambda/services"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/lambda"
	"log"
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

type LambdaObj struct {
	Function *lambda.FunctionConfiguration `json:"function"`
	Tags     interface{}                   `json:"tags"`
}

func AllLambdaFunctionsWithTagsController(auth client.Auth) (string, error) {
	fmt.Println("Request to get list of lambda functions with tags")
	lambdaClient := client.GetClient(auth, client.LAMBDA_CLIENT).(*lambda.Lambda)

	functionList := services.GetAllLambdaList(lambdaClient)

	allLambda := []LambdaObj{}
	for _, function := range functionList {
		input := &lambda.ListTagsInput{
			Resource: aws.String(*function.FunctionArn),
		}
		resp, err := lambdaClient.ListTags(input)
		if err != nil {
			fmt.Println("Error: in getting lambda tags", err)
			continue
		}

		allTags := []map[string]string{}
		for key, val := range resp.Tags {
			tagsMap := make(map[string]string)
			tagsMap["Key"] = key
			tagsMap["Value"] = *val

			allTags = append(allTags, tagsMap)
		}
		lambdaObj := LambdaObj{
			Function: function,
			Tags:     allTags,
		}
		allLambda = append(allLambda, lambdaObj)
	}
	jsonData, err := json.Marshal(allLambda)
	log.Println(string(jsonData))
	return string(jsonData), err
}

func LambdaFunctionWithTagsController(functionName string, auth client.Auth) (string, error) {
	fmt.Println("Request to get lambda function with tags")
	lambdaClient := client.GetClient(auth, client.LAMBDA_CLIENT).(*lambda.Lambda)

	lambdaDetail := services.GetLambdaDetail(lambdaClient, functionName)

	input := &lambda.ListTagsInput{
		Resource: aws.String(*lambdaDetail.Configuration.FunctionArn),
	}
	resp, err := lambdaClient.ListTags(input)
	if err != nil {
		fmt.Println("Error: in getting lambda tags", err)
		return "", err
	}

	allTags := []map[string]string{}
	for key, val := range resp.Tags {
		tagsMap := make(map[string]string)
		tagsMap["Key"] = key
		tagsMap["Value"] = *val

		allTags = append(allTags, tagsMap)
	}
	lambdaObj := LambdaObj{
		Function: lambdaDetail.Configuration,
		Tags:     resp.Tags,
	}

	jsonData, err := json.Marshal(lambdaObj)
	log.Println(string(jsonData))
	return string(jsonData), err
}
