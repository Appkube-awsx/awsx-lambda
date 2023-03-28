![lambda marked image](./lambda-home-ui-numbering.png)

| Section no. | Source | Cli commnd | Api | Response | Description |
|-------------|--------|------------|-----|----------|------------|
| 1 | aws lambda |  lambda count --all=true | /count/all | Number | It gives total number of lambda functions present in aws account |
| 2 | cloud watch logs | lambda count --idle=true | /count/idle | Number | It gives total number of idle lambda functions present in aws account |
| 3 | cloud watch logs | lambda errorCount | /errorCount | Number | It gives total number of errors in lambda functions present in aws account |
| 4 | cloud watch logs | lambda count --throttle=true | /count/throttle | Number | It gives total number of throttle lambda functions present in aws account |
| 5 | cloud watch logs | lambda lastError --env=<environment tag> | /lastError/<environment tag> | Boolean | It gives a fucntion have error or not in last some executions |
| 6 | cloud watch logs | lambda performance | /performance | Number (Percentage) | It gives overall performance percentage |
| 7 | cloud watch logs | lambda reliability | /reliability | Number (Percentage) | It gives overall reliability percentage |
| 8 | calculated data from other services | lambda avalibility | /avalibility |  Number (Percentage) | It gives overall avalibility percentage |
| 9 | calculated data from other services | lambda endUsage |/endUsage | Number (Percentage) | It gives overall endUsage percentage |
| 10 | calculated data from other services | lambda security | /security | Number (Percentage) | It gives overall security percentage |


## -- 11
Time range params in every command and api

1. Cli command

        lambda --startDate=<> --endDate=<>
2. api

    - params - 
        - startDate = <>
        - endDate = <>
