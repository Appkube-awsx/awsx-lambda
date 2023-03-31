# lambda home
![lambda marked image](./img/lambda%20home.png)
![lambda marked image](./img/lambda-home-2.png)

# Go To Events
![lambda marked image](./img/Screenshot%202023-03-31%20173321.png)
![lambda marked image](./img/Screenshot%202023-03-31%20174517.png)
![lambda marked image](./img/Screenshot%202023-03-31%20181533.png)

| Section no. | Data Format(Api/Metrics/Log/Trace) |  Source detail | Description | Logic |
|-------------|------------------------------------|----------------|-------------|-------|
| 1 | Custom API | API => `/count/all`, CLI => `lambda count --all=true` |  It gives total number of lambda functions present in aws account | -- |
| 2 | CWL | Q => `limit 1` | It gives total number of idle lambda functions present in aws account | run query for every lambda in 1 month period and if result is 0 then it is idle function |
| 3 | CWL | Q => `filter @message like /ERROR/ \| stats count() as ErrorCount` | It gives total number of errors in lambda functions present in aws account | Run query for all lambdas and it will give total number of errors |
| 4 | Metrics | Q => `--` | It gives total number of throttle lambda functions present in aws account (exceeding fixed limit) | run this query for all lambdas and get total throttels for last 24 hours |
| 5 | CWL | Q => `fields @message, @logStream \| sort @timestamp desc \| limit 5 \| filter @message like /ERROR/`  | It gives a fucntion have error or not in last some executions | If it gives data means error was present other wise null result returns |
| 6 | API/ metric | API => `/performance`, CLI => `lambda performance` |  It gives overall performance percentage | -- |
| 7 | API/ metrics | API => `/reliability`, CLI => `lambda reliability`  | It gives overall reliability percentage | -- |
| 8 | API/ metrics | API => `/avalibility`, CLI => `lambda avalibility`   | It gives overall avalibility percentage |  -- |
| 9 | API/ metrics | API => `/endUsage`, CLI => ` lambda endUsage`  | It gives overall endUsage percentage | -- |
| 10 | API/ metrics | API => ` /security`, CLI => `lambda security` |   It gives overall security percentage | -- |
| 12 | CWL | Q => `stats avg(@duration)`  | It gives average latency of all lambda functions | Run this query for all fucntions to get average latency for all function |
| 13 | CWL | Q => `stats count() by @logStream` |  It gives top 10 most executed lambdas in given range of time | Run query for all functions and get top executed fucntions |
| 14 | CWL | Q => `fields @message, @logStream \| sort @timestamp desc \| limit 5 \| filter @message like /ERROR/` | It gives number of lambda that are failed in last executions | If it gives data means error was present other wise null result returns |
| 15 | metrics | -- | It gives average cpu used by all lambda function | -- |
| 16 | metrics | -- | It total number of zones that containes lambdas | -- |
| 18 | CWL | Q => `fields @maxMemoryUsed/ 10000000 as memoryUsed_MB` | It gives average memory used by all lambdas | Run this for all lambdas and get average memory used |
| 19 | metrics | [Click here](#General-query-for-getting-data-from-metrics)  | It gives concurrency data for lambdas | -- |
| 20 | metrics | [Click here](#General-query-for-getting-data-from-metrics) | It gives response time data for lambdas | -- |
| 21 | metrics | [Click here](#General-query-for-getting-data-from-metrics) | It gives invocations data for lambdas | -- |
| 22 | metrics | [Click here](#General-query-for-getting-data-from-metrics) | It gives top failure data for lambdas | -- |
| 23 | metrics | [Click here](#General-query-for-getting-data-from-metrics) | It gives latency data for lambdas | -- |
| 24 | CWL/metrics | [Click here](#General-query-for-getting-data-from-metrics) | It gives error data for lambdas | -- |
| 25 | metrics | [Click here](#General-query-for-getting-data-from-metrics) | It gives throttle data for lambdas | -- |
| 26 | CWL/metrics | [Click here](#General-query-for-getting-data-from-metrics) | It gives trends data for lambdas | -- |

## -- 11
Time range params in every command and api

1. Cli command

        lambda --startDate=<> --endDate=<>
2. api

    - params - 
        - startDate = <>
        - endDate = <>


## General query for getting data from metrics
```json
aws cloudwatch get-metric-data --metric-data-queries '[{
    "Id": "m1",
    "MetricStat": {
    "Metric": {
        "Namespace": "<name space>",
        "MetricName": "<metrics name>",
        "Dimensions": [
            {
                "Name": "InstanceId",
                "Value": "i-0123456789abcdef"
            }
        ]
    },
    "Period": 300,
    "Stat": "Average"
    },
    "ReturnData": true
}]' --start-time 2022-01-01T00:00:00Z --end-time 2022-03-31T23:59:59Z --region us-east-1 --query 'MetricDataResults[].Values[]'
```