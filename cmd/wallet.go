/*
Copyright © 2023 Peeanio

*/
package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"github.com/spf13/viper"
	"github.com/spf13/cobra"
)

// walletCmd represents the wallet command
var walletCmd = &cobra.Command{
	Use:   "wallet",
	Short: "Character wallet",
	Long: `All wallet subcommands come from here. Invoking with no subcommands returns wallet amount.`,
	Run: func(cmd *cobra.Command, args []string) {
		getWallet()
		refresh_token()
	},
}

func init() {
	rootCmd.AddCommand(walletCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// walletCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// walletCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getWallet() {
	character_id := viper.GetString("character_id")
	token := viper.GetString("access_token")
	datasource := viper.GetString("datasource")
	base_url := viper.GetString("base_url")
	full_url := base_url + "characters/" + character_id + "/wallet/"
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, full_url, nil)
	q := req.URL.Query()
	q.Add("datasource", datasource)
	q.Add("token", token)

	if err != nil {
		fmt.Println(err)
	}
	req.URL.RawQuery = q.Encode()
		resp, _ := client.Do(req)
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		log.Fatal(err)
	}
	fmt.Println(string(resBody))
}
