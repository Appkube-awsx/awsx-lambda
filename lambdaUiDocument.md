![lambda marked image](./lambda-home-ui-numbering.png)


## -- 1
It gives total number of lambda functions present in aws account

1. Source

        aws lambda
2. Cli command

        :root count --all=true
3. Api

        /count/all
    - param - ! no param
4. Response

        Number

## -- 2  
It gives total number of idle lambda functions present in aws account

idle function: not run from 1 month

1. Source

        cloud watch logs
2. Cli command

        :root count --idle=true
3. api

        /count/idle
    - param - ! no param
4. Response

        Number

## -- 3
It gives total number of errors in  lambda functions present in aws account

idle function: not run from 1 month

1. Source

        cloud watch logs
2. Cli command

        :root errorCount
3. api

        /errorCount
    - param - ! no param
4. Response

        Number

## -- 4
It gives total number of throttle lambda functions present in aws account

throttle function: top execution functions

1. Source

        cloud watch logs
2. Cli command

        :root count --throttle=true
3. api

        /count/throttle
    - param - ! no param
4. Response

        Number

## -- 5
It gives a fucntion have error or not in last some executions

1. Source

        cloud watch logs
2. Cli command

        :root lastError --env=<environment tag>
3. api

        /lastError/<environment tag>
    - param - ! no param
4. Response

        Boolean

## -- 6
It gives overall performance percentage

1. Source

        cloud watch logs
2. Cli command

        :root performance
3. api

        /performance
    - param - ! no param
4. Response

        Number (Percentage)

## -- 7
It gives overall reliability percentage

1. Source

        calculated data from other services
2. Cli command

        :root reliability
3. api

        /reliability
    - param - ! no param
4. Response

        Number (Percentage)

## -- 8
It gives overall avalibility percentage

1. Source

        calculated data from other services
2. Cli command

        :root avalibility
3. api

        /avalibility
    - param - ! no param
4. Response

        Number (Percentage)

## -- 9
It gives overall endUsage percentage

1. Source

        calculated data from other services
2. Cli command

        :root endUsage
3. api

        /endUsage
    - param - ! no param
4. Response

        Number (Percentage)

## -- 10
It gives overall security percentage

1. Source

        calculated data from other services
2. Cli command

        :root security
3. api

        /security
    - param - ! no param
4. Response

        Number (Percentage)

## -- 11
Time range params in every command and api

1. Source

        --
2. Cli command

        :root --startDate=<> --endDate=<>
3. api

    - params - 
        - startDate = <>
        - endDate = <>
4. Response

        --