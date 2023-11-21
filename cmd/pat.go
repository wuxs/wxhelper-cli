/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// patCmd represents the pat command
var patCmd = &cobra.Command{
	Use:   "pat",
	Short: "拍一拍",
	Run: func(cmd *cobra.Command, args []string) {
		body := map[string]interface{}{
			"wxid":     msgToUser,
			"receiver": msgToUser,
		}
		url := DefaultServer + "/api/sendPatMsg"
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
	rootCmd.AddCommand(patCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// patCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// patCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
