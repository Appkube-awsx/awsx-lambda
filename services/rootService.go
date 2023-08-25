package services

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/lambda"
	"log"
)

// GetLambdaList -> get lambda list with pagination
func GetLambdaList(lambdaClient *lambda.Lambda, marker string) *lambda.ListFunctionsOutput {
	log.Println("Getting lambda list summary")
	input := &lambda.ListFunctionsInput{}

	if marker != "" {
		input = &lambda.ListFunctionsInput{
			Marker: &marker,
		}
	}

	functionList, err := lambdaClient.ListFunctions(input)

	if err != nil {
		log.Fatalln("Error: in getting lambda list", err)
	}

	return functionList

}

// GetAllLambdaList -> get all lambdas in one go
func GetAllLambdaList(lambdaClient *lambda.Lambda) []*lambda.FunctionConfiguration {
	log.Println("Getting all the lambdas in one go")
	functionList := GetLambdaList(lambdaClient, "")

	allFunctions := functionList.Functions
	marker := functionList.NextMarker

	// Loop for getting all lambdas
	for marker != nil {
		functionList = GetLambdaList(lambdaClient, *functionList.NextMarker)
		allFunctions = append(allFunctions, functionList.Functions...)
		marker = functionList.NextMarker
		fmt.Println("Functions got till now:: ", len(allFunctions))
	}

	return allFunctions
}
