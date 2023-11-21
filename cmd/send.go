/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var msgType = "text"
var msgAT = ""

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "发送消息",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			fmt.Println("请输入要发送的消息")
			os.Exit(1)
		}
		if msgAT != "" {
			if strings.HasSuffix(msgToUser, "@chatroom") {
				sendATMessage(msgAT, msgToUser, args[0])
				return
			} else {
				fmt.Println("发送群消息必须指定群聊id")
				return
			}
		}
		sendTextMessage(msgToUser, args[0])
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sendCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sendCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	sendCmd.Flags().StringVarP(&msgType, "type", "t", "text", "消息类型.")
	sendCmd.Flags().StringVarP(&msgAT, "at", "a", "", "wxid字符串，多个用,分隔，发送所有人传值\"notify@all\".")
}

func sendATMessage(wxids, chatRoomId, msg string) {
	body := map[string]interface{}{
		"wxids":      wxids,
		"chatRoomId": chatRoomId,
		"msg":        msg,
	}
	url := DefaultServer + "/api/sendAtText"
	res, err := client.R().
		SetBody(body).
		Post(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}

func sendTextMessage(wxid, msg string) {
	body := map[string]interface{}{
		"wxid": wxid,
		"msg":  msg,
	}
	url := DefaultServer + "/api/sendTextMsg"
	res, err := client.R().
		SetBody(body).
		Post(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res.String())
}
