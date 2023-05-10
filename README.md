# Lambda CLI Documentation

- [What is awsx-lambda](#awsx-lambda)
- [How to write plugin subcommand](#how-to-write-plugin-subcommand)
- [How to build / Test](#how-to-build--test)
- [How it works](#How-it-works)
- [command input](#command-input)
- [command output](#command-output)
- [How to run ](#how-to-run)

# awsx lambda
This is a plugin subcommand for awsx cli ( https://github.com/Appkube-awsx/awsx#awsx ) cli.

For details about awsx commands and how its used in Appkube platform , please refer to the diagram below:

![alt text](https://raw.githubusercontent.com/AppkubeCloud/appkube-architectures/main/LayeredArchitecture.svg)


# How to write plugin subcommand 
Please refer to the instaruction -
https://github.com/Appkube-awsx/awsx#how-to-write-a-plugin-subcommand

It has detailed instruction on how to write a subcommand plugin , build / test / debug  / publish and integrate into the main commmand.

# How to build / Test
            go run main.go
                - Program will print Calling aws-lambda on console 

            Another way of testing is by running go install command
            go install
            - go install command creates an exe with the name of the module (e.g. awsx-lambda) and save it in the GOPATH
            - Now we can execute this command on command prompt as below
            awsx-lambda --vaultURL=vault.dummy.net --accountId=xxxxxxxxxx --zone=us-west-2


# How it works

The `lambda` command-line interface (CLI) is a tool for managing lambda. This document provides instructions for using the `lambda` CLI to retrieve a list of lambda and retrieve the configuration details of a specific lambda.

## List lambdas

To list all the lambda in an account, run the following command:

    awsx-lambda --zone <zone> --accessKey <accessKey> --secretKey <secretKey> --crossAccountRoleArn <crossAccountRoleArn>
  
    awsx-lambda --vaultUrl <vaultUrl> --accountId <accountId> 


where:
- `--vaultUrl` specifies the URL of the AWS Key Management Service (KMS) customer master key (CMK) that you want to use to encrypt a lambda. This is an optional parameter. 
- `--accountId` specifies the AWS account ID that the lambda belongs to.
- `--zone` specifies the AWS region where the lambda is located.
- `--accessKey` specifies the AWS access key to use for authentication.
- `--secretKey` specifies the AWS secret key to use for authentication.
- `--crossAccountRoleArn` specifies the Amazon Resource Name (ARN) of the role that allows access to a lambda in another account. This is an optional parameter.

Example:

    awsx-lambda --vaultUrl https://mykms.us-west-2.amazonaws.com/123456 --accountId 123456789012 
  
    awsx-lambda --zone us-west-2 --accessKey AKIAIOSFODNN7EXAMPLE --secretKey wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY --crossAccountRoleArn arn:aws:iam::123456789012:role/crossAccountRole

## Get lambda Configuration

To retrieve the configuration details of a specific lambda, run the following command:

    awsx-lambda getConfigData -f <lambda> --zone <zone> --accessKey <accessKey> --secretKey <secretKey> --crossAccountRoleArn <crossAccountRoleArn>
   
    awsx-lambda getConfigData -f <lambda> --vaultUrl <vaultUrl> --accountId <accountId> 

where:
- `-f` or `--func` is the shorthand for specifying the name of the lambda. This parameter is mandatory.
- `--vaultUrl` specifies the URL of the AWS Key Management Service (KMS) customer master key (CMK) that you want to use to encrypt a lambda. This is an optional parameter. 
- `--accountId` specifies the AWS account ID that the lambda belongs to.
- `--zone` specifies the AWS region where the lambda is located.
- `--accessKey` specifies the AWS access key to use for authentication.
- `--secretKey` specifies the AWS secret key to use for authentication.
- `--crossAccountRoleArn` specifies the Amazon Resource Name (ARN) of the role that allows access to a lambda in another account. This is an optional parameter.

Example:

    awsx-lambda getConfigData -f my-lambda --vaultUrl https://mykms.us-west-2.amazonaws.com/123456 --accountId 123456789012 
 
    awsx-lambda getConfigData -f my-lambda --zone us-west-2 --accessKey AKIAIOSFODNN7EXAMPLE --secretKey wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY --crossAccountRoleArn arn:aws:iam::123456789012:role/crossAccountRole

This command returns the configuration details of the specified lambda in JSON format.


## commands tree:
    lambda
    subcommand:
        main: (get list of lambdas)
            flags:
                1. all (boolean) -> get all lambdas at once or get it by marker(pagination)
                2. marker (string) -> to get next page of lambdas list

        getConfigData: (get details of a lambda)
            flags:
                1. function (string) -> function name to get detail || required

        totalCount: (get number of all lambdas)
            flags: !no flags

        errorCount: (get total number of executions and errors)
            flags: 
                1. function (string) -> function name to get err count
    
         errorDetail: (get total number of executions and errors)
            flags: 
                1. function (string) -> function name to get err count || required


```TODO```


## command for run test
1. go to test file directory and run

```
go test -cover
```