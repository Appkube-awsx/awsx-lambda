/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package lambdacmd

import (
	"fmt"

	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-lambda/controllers"
	"github.com/spf13/cobra"
)

// GetConfigDataCmd represents the getConfigData command
var GetConfigDataCmd = &cobra.Command{
	Use:   "getConfigData",
	Short: "lambda configuration",
	Long:  `get configuration of lambda function`,
	Run: func(cmd *cobra.Command, args []string) {

		authFlag, clientAuth, err := authenticate.SubCommandAuth(cmd)
		if err != nil {
			cmd.Help()
			return
		}
		function, _ := cmd.Flags().GetString("function")

		if authFlag {
			controllers.LambdaDetails(function, *clientAuth)
		}
	},
}

func init() {
	GetConfigDataCmd.Flags().StringP("function", "f", "", "lambda function name")

	if err := GetConfigDataCmd.MarkFlagRequired("function"); err != nil {
		fmt.Println("--function or -f is required", err)
	}
}
