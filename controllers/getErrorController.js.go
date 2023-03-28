package controllers

import (
	"fmt"
	"github.com/Appkube-awsx/awsx-lambda/utils"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go/service/lambda"
	"log"
	"strings"
	"time"
)

func GetFunctionsErrDetail(cloudClient *cloudwatchlogs.CloudWatchLogs, function string) int {
	log.Println("Getting execution number and errors")

	errCount := 0
	executionCount := 0

	logGroupName := fmt.Sprintf("/aws/lambda/%s", function)

	fmt.Println("log group name", logGroupName)
	input := &cloudwatchlogs.DescribeLogStreamsInput{
		LogGroupName: aws.String(logGroupName),
	}

	firstStreamsList, err := cloudClient.DescribeLogStreams(input)
	if err != nil {
		log.Println("no cloud watch log found for this function")
	}

	nextToken := firstStreamsList.NextToken
	executionCount += len(firstStreamsList.LogStreams)

	errCount += errorDetailsInStreamList(firstStreamsList, logGroupName, cloudClient)
	fmt.Println("Total executions till now:", executionCount, "and errors are:", errCount)

	for firstStreamsList.NextToken != nil {

		input = &cloudwatchlogs.DescribeLogStreamsInput{
			LogGroupName: aws.String(logGroupName),
			NextToken:    nextToken,
		}

		tokenStreamsList, err := cloudClient.DescribeLogStreams(input)
		if err != nil {
			log.Println("no cloud watch log found for this function")
		}

		executionCount += len(tokenStreamsList.LogStreams)
		errCount += errorDetailsInStreamList(tokenStreamsList, logGroupName, cloudClient)

		fmt.Println("Total executions till now:", executionCount, "and errors are:", errCount)

	}

	fmt.Println("Final execution count is:", executionCount, "errors are:", errCount)

	return errCount
}

func errorDetailsInStreamList(firstStreamsList *cloudwatchlogs.DescribeLogStreamsOutput, logGroupName string, cloudClient *cloudwatchlogs.CloudWatchLogs) int {
	errCount := 0
	for _, stream := range firstStreamsList.LogStreams {

		input := &cloudwatchlogs.GetLogEventsInput{
			LogGroupName:  aws.String(logGroupName),
			LogStreamName: aws.String(*stream.LogStreamName),
			StartFromHead: aws.Bool(true),
		}

		resp, err := cloudClient.GetLogEvents(input)

		if err != nil {
			log.Fatalln("Error: in getting event data", err)
		}

		for _, event := range resp.Events {
			//fmt.Println("tracing error in :", *stream.LogStreamName)
			if strings.Contains(*event.Message, "ERROR") {
				fmt.Println()
				fmt.Println()
				fmt.Println("error in ", *stream.LogStreamName, "is::", event)
				errCount++
			}
		}
	}
	return errCount
}

func GetAllFunctionsErrCount(cloudClient *cloudwatchlogs.CloudWatchLogs, lambdaClient *lambda.Lambda) int {
	log.Println("Getting execution number and errors for all lambda function")
	allFunctions := GetAllLambdaList(lambdaClient)

	errCount := 0
	executionCount := 0

	for _, lambda := range allFunctions {
		tempErrCount, tempExecutionCount, err := GetFunctionErrCount(cloudClient, *lambda.FunctionName)

		if err != nil {
			if strings.Contains(err.Error(), "ResourceNotFoundException") {
				continue
			}
		}

		errCount += tempErrCount
		executionCount += tempExecutionCount

		fmt.Println("Total executions till now:", executionCount, "and errors are:", errCount)
	}

	fmt.Println("Final execution count is:", executionCount, "errors are:", errCount)
	return errCount
}

func GetFunctionErrCount(cloudClient *cloudwatchlogs.CloudWatchLogs, function string) (int, int, error) {
	log.Println("Getting execution number and errors")

	startTime := time.Date(2019, 1, 1, 1, 1, 1, 1, time.UTC) // 1 hour ago
	endTime := time.Now()
	query := `filter @message like /(?i)(Exception|error|fail|5dd)/
									| stats count() as ErrorCount`

	logGroupName := fmt.Sprintf("/aws/lambda/%s", function)

	fmt.Println("log group name", logGroupName)
	result, err := utils.GetQueryData(cloudClient, logGroupName, startTime, endTime, query)

	if err != nil {
		return 0, 0, err
	}

	fmt.Println("Result json", result)
	fmt.Println("Final execution count is:", *result.Statistics.RecordsScanned, "errors are:", *result.Statistics.RecordsMatched)
	fmt.Println()

	return int(*result.Statistics.RecordsMatched), int(*result.Statistics.RecordsScanned), nil
}

//  GetAllFunctionsErrCount -> GetFunctionErrCount -> errorCountInStreamList
