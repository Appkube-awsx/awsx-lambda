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
	Short: "gets total number of lambdas present in aws account",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		authFlag, clientAuth, err := authenticate.CommandAuth(cmd)
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
