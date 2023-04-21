/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package lambdacmd

import (
	"github.com/Appkube-awsx/awsx-lambda/authenticater"
	"github.com/Appkube-awsx/awsx-lambda/controllers"
	"github.com/spf13/cobra"
)

// GetTotalNumberOfLambdaCmd represents the number command
var GetTotalNumberOfLambdaCmd = &cobra.Command{
	Use:   "totalCount",
	Short: "gets total number of lambdas present in aws account",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		authFlag := authenticater.ChildCommandAuth(cmd)

		if authFlag {
			controllers.LambdaGetTotalNumberOfLambda(authenticater.VaultUrl, authenticater.AccountId, authenticater.Region, authenticater.AcKey, authenticater.SecKey, authenticater.CrossAccountRoleArn, authenticater.ExternalId)
		}
	},
}

func init() {}
