/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package lambdacmd

import (
	"fmt"
	"github.com/Appkube-awsx/awsx-lambda/authenticater"
	"github.com/Appkube-awsx/awsx-lambda/client"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

// getConfigDataCmd represents the getConfigData command
var GetNumberOfErrorCmd = &cobra.Command{
	Use:   "errorCount",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		vaultUrl := cmd.Parent().PersistentFlags().Lookup("vaultUrl").Value.String()
		accountNo := cmd.Parent().PersistentFlags().Lookup("accountId").Value.String()
		region := cmd.Parent().PersistentFlags().Lookup("zone").Value.String()
		acKey := cmd.Parent().PersistentFlags().Lookup("accessKey").Value.String()
		secKey := cmd.Parent().PersistentFlags().Lookup("secretKey").Value.String()
		crossAccountRoleArn := cmd.Parent().PersistentFlags().Lookup("crossAccountRoleArn").Value.String()
		externalId := cmd.Parent().PersistentFlags().Lookup("externalId").Value.String()

		authFlag := authenticater.AuthenticateData(vaultUrl, accountNo, region, acKey, secKey, crossAccountRoleArn, externalId)

		if authFlag {
			//getLambdaSigningDetail(region, crossAccountRoleArn, acKey, secKey, function, externalId)
			GetStreamList(region, crossAccountRoleArn, acKey, secKey, externalId)
		}
	},
}

func getLambdaSigningDetail(region string, crossAccountRoleArn string, accessKey string, secretKey string, function string, externalId string) *cloudwatchlogs.GetLogEventsOutput {
	log.Println("Getting Lambda  data")
	cloudClient := client.GetCloudWatchClient(region, crossAccountRoleArn, accessKey, secretKey, externalId)

	logGroupName := fmt.Sprintf("/aws/lambda/%s", function)

	input := &cloudwatchlogs.GetLogEventsInput{
		LogGroupName:  aws.String(logGroupName),
		LogStreamName: aws.String("2023/03/09/[$LATEST]55eb4c092fc74d62b088ef86249353a7"),
	}

	resp, err := cloudClient.GetLogEvents(input)
	if err != nil {
		log.Fatalln("Error: in getting lambda data", err)
	}

	// If no log events were returned, return an error
	fmt.Println(resp)

	input.NextToken = resp.NextForwardToken
	resp1, err := cloudClient.GetLogEvents(input)

	if err != nil {
		log.Fatalln("Error: in getting second data", err)
	}
	fmt.Println(resp1)

	return resp
}

func GetStreamList(region string, crossAccountRoleArn string, accessKey string, secretKey string, externalId string) int {
	log.Println("Getting execution number and errors")
	cloudClient := client.GetCloudWatchClient(region, crossAccountRoleArn, accessKey, secretKey, externalId)
	lambdaClient := client.GetClient(region, crossAccountRoleArn, accessKey, secretKey, externalId)
	input := &lambda.ListFunctionsInput{}

	functionList, err := lambdaClient.ListFunctions(input)
	if err != nil {
		log.Fatalln("Error: in getting total number of lambdas", err)
	}

	allFunctions := functionList.Functions
	marker := functionList.NextMarker

	// Loop for getting all lambdas
	for marker != nil {
		input = &lambda.ListFunctionsInput{
			Marker: marker,
		}
		functionList, err = lambdaClient.ListFunctions(input)
		if err != nil {
			log.Fatalln("Error: in getting lambda numbers", err)
		}
		allFunctions = append(allFunctions, functionList.Functions...)
		marker = functionList.NextMarker
		fmt.Println("Functions got till now:: ", len(allFunctions))
	}

	errCount := 0
	executionCount := 0

	for index, lambda := range allFunctions {

		logGroupName := fmt.Sprintf("/aws/lambda/%s", *lambda.FunctionName)
		fmt.Println()
		fmt.Println(index+1, "/", len(allFunctions))
		fmt.Println("log group name", logGroupName)
		input := &cloudwatchlogs.DescribeLogStreamsInput{
			LogGroupName: aws.String(logGroupName),
		}

		allStreams, err := cloudClient.DescribeLogStreams(input)
		if err != nil {
			log.Println("no cloud watch log found for this function")
		}

		executionCount += len(allStreams.LogStreams)

		for _, stream := range allStreams.LogStreams {

			input := &cloudwatchlogs.GetLogEventsInput{
				LogGroupName:  aws.String(logGroupName),
				LogStreamName: aws.String(*stream.LogStreamName),
				StartFromHead: aws.Bool(true),
			}

			resp, err := cloudClient.GetLogEvents(input)

			if err != nil {
				log.Fatalln("Error: in getting event data", err)
			}

			//fmt.Println("events:::::::::", *stream.LogStreamName, ":::::", allStreams)

			for _, event := range resp.Events {
				//fmt.Println("message", *event.Message)

				if strings.Contains(*event.Message, "ERROR") {
					errCount++
					break
				}
			}
		}
		fmt.Println("Total executions till now:", executionCount, "and errors are:", errCount)
	}

	fmt.Println("Final execution count is:", executionCount, "errors are:", errCount)

	return errCount
}

func init() {}
