package authenticate

//
//import (
//	"fmt"
//	"github.com/Appkube-awsx/awsx-common/client"
//	"github.com/Appkube-awsx/awsx-lambda/vault"
//	"github.com/spf13/cobra"
//	"log"
//)
//
//// ClientAuth for storing auth data
////var ClientAuth client.Auth
//
//// AuthenticateData -> For account validation
//func AuthenticateData(vaultUrl string, vaultToken string, accountNo string, region string, acKey string, secKey string, crossAccountRoleArn string, externalId string) (bool, *client.Auth, error) {
//	//if region == "" {
//	//	log.Fatalln("Region not provided. Program exit")
//	//	return false, nil
//	//}
//	if vaultUrl != "" {
//		log.Println("vault url provided. getting user credentials from vault")
//		if vaultToken == "" {
//			log.Println("vault token missing")
//			return false, nil, fmt.Errorf("vault token missing")
//		}
//		if accountNo == "" {
//			log.Println("account no missing")
//			return false, nil, fmt.Errorf("account no missing")
//		}
//		log.Println("Getting account details from vault")
//		vaultResp, err := vault.GetAccountDetails(vaultUrl, vaultToken, accountNo)
//		if err != nil {
//			log.Println("Error in calling vault api to get account details. \n", err)
//			return false, nil, err
//		}
//		if vaultResp.Data.AccessKey == "" || vaultResp.Data.SecretKey == "" || vaultResp.Data.CrossAccountRoleArn == "" || vaultResp.Data.ExternalId == "" {
//			log.Println("account details not found in vault")
//			return false, nil, fmt.Errorf("account details not found in vault")
//		}
//
//		clientAuth := client.Auth{
//			CrossAccountRoleArn: vaultResp.Data.CrossAccountRoleArn,
//			AccessKey:           vaultResp.Data.AccessKey,
//			SecretKey:           vaultResp.Data.SecretKey,
//			ExternalId:          vaultResp.Data.ExternalId,
//		}
//		if region != "" {
//			clientAuth.Region = region
//		} else {
//			log.Println("region not provided. default region will be used")
//			clientAuth.Region = vaultResp.Data.Region
//		}
//		return true, &clientAuth, nil
//
//	}
//	log.Println("vault url not provided. validating provided user credentials")
//	if region == "" {
//		log.Println("region missing")
//		return false, nil, fmt.Errorf("region missing")
//	}
//	if acKey == "" {
//		log.Println("access key missing")
//		return false, nil, fmt.Errorf("access key missing")
//	}
//	if secKey == "" {
//		log.Println("secret key missing")
//		return false, nil, fmt.Errorf("secret key missing")
//	}
//	if crossAccountRoleArn == "" {
//		log.Println("cross account role arn missing")
//		return false, nil, fmt.Errorf("cross account role arn missing")
//	}
//	if externalId == "" {
//		log.Println("external id missing")
//		return false, nil, fmt.Errorf("external id missing")
//	}
//	clientAuth := client.Auth{
//		Region:              region,
//		CrossAccountRoleArn: crossAccountRoleArn,
//		AccessKey:           acKey,
//		SecretKey:           secKey,
//		ExternalId:          externalId,
//	}
//	return true, &clientAuth, nil
//}
//
//// CommandAuth -> For validation of child command
//func CommandAuth(cmd *cobra.Command) (bool, *client.Auth, error) {
//	vaultUrl := cmd.Parent().PersistentFlags().Lookup("vaultUrl").Value.String()
//	vaultToken := cmd.Parent().PersistentFlags().Lookup("vaultToken").Value.String()
//	accountNo := cmd.Parent().PersistentFlags().Lookup("accountId").Value.String()
//	region := cmd.Parent().PersistentFlags().Lookup("zone").Value.String()
//	acKey := cmd.Parent().PersistentFlags().Lookup("accessKey").Value.String()
//	secKey := cmd.Parent().PersistentFlags().Lookup("secretKey").Value.String()
//	crossAccountRoleArn := cmd.Parent().PersistentFlags().Lookup("crossAccountRoleArn").Value.String()
//	externalId := cmd.Parent().PersistentFlags().Lookup("externalId").Value.String()
//	authFlag, clientAuth, err := AuthenticateData(vaultUrl, vaultToken, accountNo, region, acKey, secKey, crossAccountRoleArn, externalId)
//
//	return authFlag, clientAuth, err
//}
//
//// RootCommandAuth -> For validation of parent command
////func RootCommandAuth(cmd *cobra.Command) (bool, *client.Auth) {
////
////	vaultUrl := cmd.PersistentFlags().Lookup("vaultUrl").Value.String()
////	vaultToken := cmd.PersistentFlags().Lookup("vaultToken").Value.String()
////	accountId := cmd.PersistentFlags().Lookup("accountId").Value.String()
////	zone := cmd.PersistentFlags().Lookup("zone").Value.String()
////	accessKey := cmd.PersistentFlags().Lookup("accessKey").Value.String()
////	secretKey := cmd.PersistentFlags().Lookup("secretKey").Value.String()
////	crossAccountRoleArn := cmd.PersistentFlags().Lookup("crossAccountRoleArn").Value.String()
////	externalId := cmd.PersistentFlags().Lookup("externalId").Value.String()
////	authFlag, clientAuth := AuthenticateData(vaultUrl, vaultToken, accountId, zone, accessKey, secretKey, crossAccountRoleArn, externalId)
////
////	return authFlag, clientAuth
////}
//
//// ApiAuth -> for authentication of rest-api request
////func ApiAuth(vaultUrl string, vaultToken string, accountNo string, region string, acKey string, secKey string, crossAccountRoleArn string, externalId string) (bool, *client.Auth) {
////	authFlag, clientAuth := AuthenticateData(vaultUrl, vaultToken, accountNo, region, acKey, secKey, crossAccountRoleArn, externalId)
////	return authFlag, clientAuth
////}
