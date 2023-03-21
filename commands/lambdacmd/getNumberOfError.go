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
			function, _ := cmd.Flags().GetString("function")
			//getLambdaSigningDetail(region, crossAccountRoleArn, acKey, secKey, function, externalId)

			if function != "" {
				GetFunctionErrCount(region, crossAccountRoleArn, acKey, secKey, externalId, function)
			} else {
				GetAllFunctionsErrCount(region, crossAccountRoleArn, acKey, secKey, externalId)
			}
		}
	},
}

func GetAllFunctionsErrCount(region string, crossAccountRoleArn string, accessKey string, secretKey string, externalId string) int {
	log.Println("Getting execution number and errors for all lambda function")
	allFunctions := GetAllLambdaList(region, crossAccountRoleArn, accessKey, secretKey, externalId)

	errCount := 0
	executionCount := 0

	for _, lambda := range allFunctions {
		tempErrCount, tempExecutionCount := GetFunctionErrCount(region, crossAccountRoleArn, accessKey, secretKey, externalId, *lambda.FunctionName)

		errCount += tempErrCount
		executionCount += tempExecutionCount

		fmt.Println("Total executions till now:", executionCount, "and errors are:", errCount)
	}

	fmt.Println("Final execution count is:", executionCount, "errors are:", errCount)
	return errCount
}

func GetFunctionErrCount(region string, crossAccountRoleArn string, accessKey string, secretKey string, externalId string, function string) (int, int) {
	log.Println("Getting execution number and errors")
	cloudClient := client.GetCloudWatchClient(region, crossAccountRoleArn, accessKey, secretKey, externalId)

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

func GetAllLambdaList(region string, crossAccountRoleArn string, accessKey string, secretKey string, externalId string) []*lambda.FunctionConfiguration {
	log.Println("Getting lambda list summary")
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

	return allFunctions
}

func init() {
	GetNumberOfErrorCmd.Flags().StringP("function", "f", "", "lambda function name")
}
