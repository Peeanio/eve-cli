/*
Copyright © 2023 Peeanio

*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/spf13/viper"
	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "searches for a [1] eve object of the name provided",
	Long: "When trying to find data or specific information about an object, use search to find it. single search item",
	Run: func(cmd *cobra.Command, args []string) {
		getSearch(args[0])
		refresh_token()
	},
}

var category string

func init() {
	rootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")
	searchCmd.Flags().StringVarP(&category, "category", "c", "character",
	`Allowed values: "agent", "alliance", "character", "constellation", "corporation", "faction", "inventory_type", "region", "solar_system", "station", "structure"`)
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
func getSearch(search_string string) {
	character_id := viper.GetString("character_id")
	token := viper.GetString("access_token")
	datasource := viper.GetString("datasource")
	base_url := viper.GetString("base_url")
	full_url := base_url + "characters/" + character_id + "/search/" + "?categories=" + category + "&search=" + search_string
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
	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		log.Fatal(err)
	}
	fmt.Println(string(resBody))
}
