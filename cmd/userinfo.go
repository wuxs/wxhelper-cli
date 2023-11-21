/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// userinfoCmd represents the userinfo command
var userinfoCmd = &cobra.Command{
	Use:   "userinfo",
	Short: "获取登录用户信息",
	Run: func(cmd *cobra.Command, args []string) {
		body := map[string]interface{}{
			"wxid":     msgToUser,
			"receiver": msgToUser,
		}
		url := DefaultServer + "/api/userInfo"
		res, err := client.R().
			SetBody(body).
			Post(url)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(res.String())
	},
}

func init() {
	rootCmd.AddCommand(userinfoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// userinfoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// userinfoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
