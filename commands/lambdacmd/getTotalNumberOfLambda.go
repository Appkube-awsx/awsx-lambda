/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package lambdacmd

import (
	"fmt"
	"github.com/Appkube-awsx/awsx-lambda/authenticater"
	"github.com/Appkube-awsx/awsx-lambda/client"
	"github.com/Appkube-awsx/awsx-lambda/services"
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
			lambdaClient := client.GetClient()

			totalNumber := len(services.GetAllLambdaList(lambdaClient))
			fmt.Println("total number of lambda present in aws account in", authenticater.Region, "is:", totalNumber)
		}
	},
}

func init() {}
