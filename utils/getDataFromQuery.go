package utils

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"log"
	"time"
)

func GetQueryData(cloudClient *cloudwatchlogs.CloudWatchLogs, logGroupName string, startTime time.Time, endTime time.Time, query string) (*cloudwatchlogs.GetQueryResultsOutput, error) {
	input := &cloudwatchlogs.StartQueryInput{
		LogGroupName: aws.String(logGroupName),
		QueryString:  aws.String(query),
		StartTime:    aws.Int64(startTime.UnixNano() / int64(time.Millisecond)),
		EndTime:      aws.Int64(endTime.UnixNano() / int64(time.Millisecond)),
	}

	data, err := cloudClient.StartQuery(input)

	if err != nil {
		log.Println("error in getting queryId", err)
		return nil, err
	}

	// get result
	input1 := &cloudwatchlogs.GetQueryResultsInput{
		QueryId: data.QueryId,
	}

	result, err := cloudClient.GetQueryResults(input1)

	if err != nil {
		log.Fatalln("error in getting number of errors of lambda", err)
	}

	scanningTime := 0
	// Wait for Complete
	for *result.Status != "Complete" {
		fmt.Println(*result.Status, "...", scanningTime, "seconds")
		result, err = cloudClient.GetQueryResults(input1)
		if err != nil {
			log.Fatalln("error in waiting", err)
		}

		scanningTime += 2
		time.Sleep(time.Second * 2)
	}

	return result, nil
}
