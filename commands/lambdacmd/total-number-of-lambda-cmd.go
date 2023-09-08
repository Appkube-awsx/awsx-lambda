/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package lambdacmd

import (
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-lambda/controllers"
	"github.com/spf13/cobra"
)

// GetTotalNumberOfLambdaCmd represents the number command
var GetTotalNumberOfLambdaCmd = &cobra.Command{
	Use:   "totalCount",
	Short: "total number of lambdas in an aws account",
	Long:  `get total number of lambdas present in aws account`,
	Run: func(cmd *cobra.Command, args []string) {

		authFlag, clientAuth, err := authenticate.SubCommandAuth(cmd)
		if err != nil {
			cmd.Help()
			return
		}

		if authFlag {
			controllers.LambdaGetTotalNumberOfLambda(*clientAuth)
		}
	},
}

func init() {}
