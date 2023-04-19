package actuator

import (
	"fmt"
	"github.com/Appkube-awsx/awsx-lambda/authenticater"
	"github.com/Appkube-awsx/awsx-lambda/client"
	"github.com/Appkube-awsx/awsx-lambda/controllers"
	"github.com/aws/aws-sdk-go/service/lambda"
)

func LambdaListActuator(marker string, all bool, vaultUrl string, accountId string, region string, acKey string, secKey string, crossAccountRoleArn string, externalId string) []*lambda.FunctionConfiguration {

	// For request from API
	authFlag := authenticater.AuthenticateData(vaultUrl, accountId, region, acKey, secKey, crossAccountRoleArn, externalId)
	if authFlag {
		authenticater.VaultUrl = vaultUrl
		authenticater.AccountId = accountId
		authenticater.Region = region
		authenticater.AcKey = acKey
		authenticater.SecKey = secKey
		authenticater.CrossAccountRoleArn = crossAccountRoleArn
		authenticater.ExternalId = externalId
	}

	lambdaClient := client.GetClient()

	functionList := controllers.GetAllLambdaList(lambdaClient)
	fmt.Println("List of all lambda functions", functionList)
	return functionList
}
