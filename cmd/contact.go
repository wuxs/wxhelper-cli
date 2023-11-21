/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

type ContactRequest struct {
}

//{
//    "code": 1,
//    "data": [
//        {
//           "customAccount": "",
//            "encryptName": "v3_020b3826fd03010000000000e04128fddf4d90000000501ea9a3dba12f95f6b60a0536a1adb6b40fc4086288f46c0b89e6c4eb8062bb1661b4b6fbab708dc4f89d543d7ade135b2be74c14b9cfe3accef377b9@stranger",
//            "nickname": "文件传输助手",
//            "pinyin": "WJCSZS",
//            "pinyinAll": "wenjianchuanshuzhushou",
//            "reserved1": 1,
//            "reserved2": 1,
//            "type": 3,
//            "verifyFlag": 0,
//            "wxid": "filehelper"
//        }
//    ].
//    "msg": "success"
//}

type ContactInfo struct {
	CustomAccount string `json:"customAccount"`
	EncryptName   string `json:"encryptName"`
	Nickname      string `json:"nickname"`
	Pinyin        string `json:"pinyin"`
	PinyinAll     string `json:"pinyinAll"`
	Reserved1     int    `json:"reserved1"`
	Reserved2     int    `json:"reserved2"`
	Type          int    `json:"type"`
	VerifyFlag    int    `json:"verifyFlag"`
	Wxid          string `json:"wxid"`
}
type ContactResponse struct {
	Code int           `json:"code"`
	Data []ContactInfo `json:"data"`
}

// contactCmd represents the contact command
var contactCmd = &cobra.Command{
	Use:   "contact",
	Short: "好友列表",
	Run: func(cmd *cobra.Command, args []string) {
		var res = &ContactResponse{}
		url := DefaultServer + "/api/getContactList"
		_, err := client.R().
			SetResult(res).
			Post(url)
		if err != nil {
			fmt.Println(err)
			return
		}
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Nickname", "Wxid", "Pinyin", "Type"})
		table.SetBorder(false)
		for _, info := range res.Data {
			table.Append([]string{info.Nickname, info.Wxid, info.Pinyin, strconv.Itoa(info.Type)})
		}
		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(contactCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// contactCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// contactCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
