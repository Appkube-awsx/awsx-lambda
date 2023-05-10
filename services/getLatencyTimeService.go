package services

import (
	"fmt"
	"log"
	"time"

	"github.com/Appkube-awsx/awsx-lambda/utils"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

func GetLambadaLatencyTime(cloudClient *cloudwatchlogs.CloudWatchLogs, function string, startTime string, endTime string) (string, int, error) {
	log.Println("Getting execution number and errors")
	logGroupName := fmt.Sprintf("/aws/lambda/%s", function)

	layout := "2006-01-02 15:04:05.000000000 -0700 MST"
	fmt.Println("log group name", logGroupName)

	// startTime = time.Date(2019, 1, 1, 1, 1, 1, 1, time.UTC) // 1 hour ago
	input := &cloudwatchlogs.DescribeLogStreamsInput{
		LogGroupName: aws.String(logGroupName),
	}

	resp, err := cloudClient.DescribeLogStreams(input)

	dataCreationTime := resp.LogStreams[0].CreationTime
	durationCreationTime := int64(*dataCreationTime)

	durationCreationTimeInSeconds := durationCreationTime / 1000
	creationTime := time.Unix(durationCreationTimeInSeconds, 0)
	creationTimeString := creationTime.Format("2006-01-02 15:04:05.000000000 -0700 MST")

	if err != nil {
		log.Fatalln("Error: in getting event data", err)
	}

	var up_endTime time.Time
	var up_startTime time.Time

	if endTime == "" && startTime == "" {
		up_endTime = time.Now()
		up_startTime, err = time.Parse("2006-01-02 15:04:05.000000000 -0700 MST", creationTimeString)

		if err != nil {
			log.Fatalln("Error: in getting event data", err)
		}

	} else if startTime == "" {
		up_startTime, err = time.Parse(layout, creationTimeString)
		if err != nil {
			log.Fatalln("Error: in getting event data", err)
		}
		up_endTime, err = time.Parse(layout, endTime)
		fmt.Println("this is updt start time", up_startTime)

		if err != nil {
			log.Fatalln("Error: in getting event data", err)
		}

	} else if endTime == "" {
		up_startTime, err = time.Parse(layout, startTime)
		up_endTime = time.Now()
		fmt.Println("this is updt end time", up_endTime)
		if err != nil {
			log.Fatal("This is time", err)
		}

	} else {
		up_startTime, err = time.Parse(layout, startTime)
		if err != nil {
			log.Fatalln("Error: in getting event data", err)
		}
		up_endTime, err = time.Parse(layout, endTime)
		if err != nil {
			log.Fatalln("Error: in getting event data", err)
		}
	}

	query := `stats avg(@duration)`

	result, err := utils.GetQueryData(cloudClient, logGroupName, up_startTime, up_endTime, query)
	fmt.Println("This si result", result)

	if err != nil {
		return "", 0, err
	}

	fmt.Println("Result json", result)
	// fmt.Println("Avg latency of ", logGroupName, "is", *result.Results[0][0].Value, "ms")
	fmt.Println()

	return *result.Results[0][0].Value, int(*result.Statistics.RecordsScanned), nil
}
