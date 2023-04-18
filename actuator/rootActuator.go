package actuator

import (
	"fmt"
	"github.com/Appkube-awsx/awsx-lambda/authenticater"
	"github.com/Appkube-awsx/awsx-lambda/client"
	"github.com/Appkube-awsx/awsx-lambda/controllers"
)

func LambdaListActuator(marker string, all bool, vaultUrl string, accountId string, region string, acKey string, secKey string, crossAccountRoleArn string, externalId string) {

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

	if all {
		functionList := controllers.GetAllLambdaList(lambdaClient)
		fmt.Println("List of all lambda functions", functionList)
	} else {
		functionList := controllers.GetLambdaList(lambdaClient, marker)
		fmt.Println("List of by marker lambda functions", functionList)
	}
}
