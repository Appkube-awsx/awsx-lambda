package authenticater

import (
	"github.com/Appkube-awsx/awsx-common/client"
	"github.com/Appkube-awsx/awsx-lambda/vault"
	"github.com/spf13/cobra"
	"log"
)

// ClientAuth for storing auth data
var ClientAuth client.Auth

// AuthenticateData -> For account validation
func AuthenticateData(vaultUrl string, vaultToken string, accountNo string, region string, acKey string, secKey string, crossAccountRoleArn string, externalId string) (bool, *client.Auth) {
	if region == "" {
		log.Fatalln("Region not provided. Program exit")
		return false, nil
	}
	if vaultUrl != "" {
		if accountNo == "" {
			log.Fatalln("Account key not provided. Program exit")
			return false, nil
		}
		if vaultToken == "" {
			log.Println("Vault token not provided. Program exit")
			return false, nil
		}
		log.Println("Getting account details from vault")
		vaultResp, err := vault.GetAccountDetails(vaultUrl, vaultToken, accountNo)
		if err != nil {
			log.Println("Error in calling vault api to get account details. \n", err)
			return false, nil
		}
		if vaultResp.Data.AccessKey == "" || vaultResp.Data.SecretKey == "" || vaultResp.Data.CrossAccountRoleArn == "" || vaultResp.Data.ExternalId == "" {
			log.Println("Account details not found in vault.")
			return false, nil
		}

		ClientAuth := client.Auth{
			Region:              region,
			CrossAccountRoleArn: vaultResp.Data.CrossAccountRoleArn,
			AccessKey:           vaultResp.Data.AccessKey,
			SecretKey:           vaultResp.Data.SecretKey,
			ExternalId:          vaultResp.Data.ExternalId,
		}

		return true, &ClientAuth

	} else if acKey != "" && secKey != "" && crossAccountRoleArn != "" && externalId != "" {
		ClientAuth := client.Auth{
			Region:              region,
			CrossAccountRoleArn: crossAccountRoleArn,
			AccessKey:           acKey,
			SecretKey:           secKey,
			ExternalId:          externalId,
		}
		return true, &ClientAuth
	} else {
		log.Fatal("AWS credentials like accessKey/secretKey/region/crossAccountRoleArn/externalId not provided. Program exit")
		return false, nil
	}
}

// ChildCommandAuth -> For validation of child command
func ChildCommandAuth(cmd *cobra.Command) (bool, *client.Auth) {

	//ClientAuth = client.Auth{
	//	cmd.Parent().PersistentFlags().Lookup("vaultUrl").Value.String(),
	//	cmd.Parent().PersistentFlags().Lookup("vaultToken").Value.String(),
	//	cmd.Parent().PersistentFlags().Lookup("accountId").Value.String(),
	//	cmd.Parent().PersistentFlags().Lookup("zone").Value.String(),
	//	cmd.Parent().PersistentFlags().Lookup("crossAccountRoleArn").Value.String(),
	//	cmd.Parent().PersistentFlags().Lookup("accessKey").Value.String(),
	//	cmd.Parent().PersistentFlags().Lookup("secretKey").Value.String(),
	//	cmd.Parent().PersistentFlags().Lookup("externalId").Value.String(),
	//}
	vaultUrl := cmd.Parent().PersistentFlags().Lookup("vaultUrl").Value.String()
	vaultToken := cmd.Parent().PersistentFlags().Lookup("vaultToken").Value.String()
	accountId := cmd.Parent().PersistentFlags().Lookup("accountId").Value.String()
	zone := cmd.Parent().PersistentFlags().Lookup("zone").Value.String()
	accessKey := cmd.Parent().PersistentFlags().Lookup("accessKey").Value.String()
	secretKey := cmd.Parent().PersistentFlags().Lookup("secretKey").Value.String()
	crossAccountRoleArn := cmd.Parent().PersistentFlags().Lookup("crossAccountRoleArn").Value.String()
	externalId := cmd.Parent().PersistentFlags().Lookup("externalId").Value.String()
	authFlag, clientAuth := AuthenticateData(vaultUrl, vaultToken, accountId, zone, accessKey, secretKey, crossAccountRoleArn, externalId)

	return authFlag, clientAuth
}

// RootCommandAuth -> For validation of parent command
func RootCommandAuth(cmd *cobra.Command) (bool, *client.Auth) {

	//ClientAuth = client.Auth{
	//	cmd.PersistentFlags().Lookup("vaultUrl").Value.String(),
	//	cmd.PersistentFlags().Lookup("vaultToken").Value.String(),
	//	cmd.PersistentFlags().Lookup("accountId").Value.String(),
	//	cmd.PersistentFlags().Lookup("zone").Value.String(),
	//	cmd.PersistentFlags().Lookup("crossAccountRoleArn").Value.String(),
	//	cmd.PersistentFlags().Lookup("accessKey").Value.String(),
	//	cmd.PersistentFlags().Lookup("secretKey").Value.String(),
	//	cmd.PersistentFlags().Lookup("externalId").Value.String(),
	//}

	vaultUrl := cmd.PersistentFlags().Lookup("vaultUrl").Value.String()
	vaultToken := cmd.PersistentFlags().Lookup("vaultToken").Value.String()
	accountId := cmd.PersistentFlags().Lookup("accountId").Value.String()
	zone := cmd.PersistentFlags().Lookup("zone").Value.String()
	accessKey := cmd.PersistentFlags().Lookup("accessKey").Value.String()
	secretKey := cmd.PersistentFlags().Lookup("secretKey").Value.String()
	crossAccountRoleArn := cmd.PersistentFlags().Lookup("crossAccountRoleArn").Value.String()
	externalId := cmd.PersistentFlags().Lookup("externalId").Value.String()
	authFlag, clientAuth := AuthenticateData(vaultUrl, vaultToken, accountId, zone, accessKey, secretKey, crossAccountRoleArn, externalId)

	//authFlag := AuthenticateData(ClientAuth.VaultUrl, ClientAuth.VaultToken, ClientAuth.VaultKey, ClientAuth.Region, ClientAuth.AccessKey, ClientAuth.SecretKey, ClientAuth.CrossAccountRoleArn, ClientAuth.ExternalId)

	return authFlag, clientAuth
}

// ApiAuth -> for authentication of rest-api request
func ApiAuth(vaultUrl string, vaultToken string, accountNo string, region string, acKey string, secKey string, crossAccountRoleArn string, externalId string) (bool, *client.Auth) {

	authFlag, clientAuth := AuthenticateData(vaultUrl, vaultToken, accountNo, region, acKey, secKey, crossAccountRoleArn, externalId)

	if !authFlag {
		log.Fatalln("authentication error")
	}
	return authFlag, clientAuth
}
