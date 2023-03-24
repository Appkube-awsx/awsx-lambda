package controllers

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go/service/lambda"
	"log"
	"strings"
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
		tempErrCount, tempExecutionCount := GetFunctionErrCount(cloudClient, *lambda.FunctionName)

		errCount += tempErrCount
		executionCount += tempExecutionCount

		fmt.Println("Total executions till now:", executionCount, "and errors are:", errCount)
	}

	fmt.Println("Final execution count is:", executionCount, "errors are:", errCount)
	return errCount
}

func GetFunctionErrCount(cloudClient *cloudwatchlogs.CloudWatchLogs, function string) (int, int) {
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

	errCount += errorCountInStreamList(firstStreamsList, logGroupName, cloudClient)
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
		errCount += errorCountInStreamList(tokenStreamsList, logGroupName, cloudClient)

		fmt.Println("Total executions till now:", executionCount, "and errors are:", errCount)
	}

	fmt.Println("Final execution count is:", executionCount, "errors are:", errCount)

	return errCount, executionCount
}

func errorCountInStreamList(firstStreamsList *cloudwatchlogs.DescribeLogStreamsOutput, logGroupName string, cloudClient *cloudwatchlogs.CloudWatchLogs) int {
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
				errCount++
				break
			}
		}
	}
	return errCount
}

//  GetAllFunctionsErrCount -> GetFunctionErrCount -> errorCountInStreamList
