![lambda marked image](./lambda-home-ui-numbering.png)
![lambda marked image](./lambda-home-2.png)

| Section no. | Source |  Source detail |Cli commnd | Api | Response | Description |
|-------------|--------|----------------|-----------|-----|----------|-------------|
| 1 | aws lambda | `lambdaClient.ListFunctions(input)` | lambda count --all=true | /count/all | Number | It gives total number of lambda functions present in aws account |
| 2 | CWL | -- | lambda count --idle=true | /count/idle | Number | It gives total number of idle lambda functions present in aws account |
| 3 | CWL | `filter @message like /(?i) (Exception \| error \| fail \| 5dd)/ \| stats count() as ErrorCount` | lambda errorCount | /errorCount | Number | It gives total number of errors in lambda functions present in aws account |
| 4 | CWL/metrics | -- | lambda count --throttle=true | /count/throttle | Number | It gives total number of throttle lambda functions present in aws account (exceeding fixed limit) |
| 5 | CWL | `fields @message, @logStream \| sort @timestamp desc \| limit 10` | lambda lastError --env=<environment tag> | /lastError/<environment tag> | Boolean | It gives a fucntion have error or not in last some executions |
| 6 | calculated data from other services/ metric | -- | lambda performance | /performance | Number (Percentage) | It gives overall performance percentage |
| 7 | calculated data from other services/ metrics | -- | lambda reliability | /reliability | Number (Percentage) | It gives overall reliability percentage |
| 8 | calculated data from other services/ metrics | -- | lambda avalibility | /avalibility |  Number (Percentage) | It gives overall avalibility percentage |
| 9 | calculated data from other services/ metrics | -- | lambda endUsage |/endUsage | Number (Percentage) | It gives overall endUsage percentage |
| 10 | calculated data from other services/ metrics | -- | lambda security | /security | Number (Percentage) | It gives overall security percentage |
| 12 | CWL | Q => `stats avg(@duration)`  | lambda latency | /lambda/latancy | Number (mili seconds ) | It gives average latency of all lambda functions | 
| 13 | CWL / metrics | Q => `stats count() by @logStream` | lambda trends | /lambda/trends | Number | It gives top 10 most executed lambdas in given range of time
| 14 | CWL/ metrics | Q => `fields @message, @logStream \| sort @timestamp desc \| limit 10` and check it is error or not | lambda failures | /lambda/failures | It gives number of lambda that are failed in last executions  
| 15 | CWL/ metrics | -- | lambda cpuUsed | /lambda/cpuUsed | It gives average cpu used by all lambda function |
| 16 | CWL/ metrics | -- | lambda netRecieved | /lambda/netRecieved | It total number of zones that containes lambdas |
| 18 | CWL | Q => `fields @maxMemoryUsed/ 10000000 as memoryUsed_MB` | lambda memoryUsed | Number (MB) | It gives average memory used by all lambdas



## -- 11
Time range params in every command and api

1. Cli command

        lambda --startDate=<> --endDate=<>
2. api

    - params - 
        - startDate = <>
        - endDate = <>
