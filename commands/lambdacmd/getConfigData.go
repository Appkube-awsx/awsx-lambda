/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package lambdacmd

import (
	"fmt"

	"github.com/Appkube-awsx/awsx-lambda/authenticater"
	"github.com/Appkube-awsx/awsx-lambda/controllers"
	"github.com/spf13/cobra"
)

// GetConfigDataCmd represents the getConfigData command
var GetConfigDataCmd = &cobra.Command{
	Use:   "getConfigData",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		authFlag := authenticater.ChildCommandAuth(cmd)

		function, _ := cmd.Flags().GetString("function")

		if authFlag {
			controllers.LambdaDetails(function, authenticater.ClientAuth)
		}
	},
}

func init() {
	GetConfigDataCmd.Flags().StringP("function", "f", "", "lambda function name")

	if err := GetConfigDataCmd.MarkFlagRequired("function"); err != nil {
		fmt.Println("--function or -f is required", err)
	}
}
