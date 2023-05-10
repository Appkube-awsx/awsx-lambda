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
func AuthenticateData(vaultUrl string, accountNo string, region string, acKey string, secKey string, crossAccountRoleArn string, externalId string) bool {

	if vaultUrl != "" && accountNo != "" {
		if region == "" {
			log.Fatalln("Zone not provided. Program exit")
			return false
		}
		log.Println("Getting account details")
		data, err := vault.GetAccountDetails(vaultUrl, accountNo)
		if err != nil {
			log.Println("Error in calling the account details api. \n", err)
			return false
		}
		if data.AccessKey == "" || data.SecretKey == "" || data.CrossAccountRoleArn == "" {
			log.Println("Account details not found.")
			return false
		}
		return true

	} else if region != "" && acKey != "" && secKey != "" && crossAccountRoleArn != "" && externalId != "" {
		return true
	} else {
		log.Fatal("AWS credentials like accesskey/secretkey/region/crossAccountRoleArn/externalId not provided. Program exit")
		return false
	}
}

// ChildCommandAuth -> For validation of child command
func ChildCommandAuth(cmd *cobra.Command) bool {

	ClientAuth = client.Auth{
		cmd.Parent().PersistentFlags().Lookup("zone").Value.String(),
		cmd.Parent().PersistentFlags().Lookup("crossAccountRoleArn").Value.String(),
		cmd.Parent().PersistentFlags().Lookup("accessKey").Value.String(),
		cmd.Parent().PersistentFlags().Lookup("secretKey").Value.String(),
		cmd.Parent().PersistentFlags().Lookup("externalId").Value.String(),
	}
	authFlag := AuthenticateData("", "", ClientAuth.Region, ClientAuth.AccessKey, ClientAuth.SecretKey, ClientAuth.CrossAccountRoleArn, ClientAuth.ExternalId)

	return authFlag
}

// RootCommandAuth -> For validation of parent command
func RootCommandAuth(cmd *cobra.Command) bool {

	ClientAuth = client.Auth{
		cmd.PersistentFlags().Lookup("zone").Value.String(),
		cmd.PersistentFlags().Lookup("crossAccountRoleArn").Value.String(),
		cmd.PersistentFlags().Lookup("accessKey").Value.String(),
		cmd.PersistentFlags().Lookup("secretKey").Value.String(),
		cmd.PersistentFlags().Lookup("externalId").Value.String(),
	}

	authFlag := AuthenticateData("", "", ClientAuth.Region, ClientAuth.AccessKey, ClientAuth.SecretKey, ClientAuth.CrossAccountRoleArn, ClientAuth.ExternalId)

	return authFlag
}

// ApiAuth -> for authentication of api request
func ApiAuth(auth client.Auth) bool {

	authFlag := AuthenticateData("", "", auth.Region, auth.AccessKey, auth.SecretKey, auth.CrossAccountRoleArn, auth.ExternalId)

	if !authFlag {
		log.Fatalln("authentication error")
	}
	return authFlag
}
