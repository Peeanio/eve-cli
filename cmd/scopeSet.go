/*
Copyright © 2023 Peeanio

*/
package cmd

import (
	"fmt"
	"strings"
	"github.com/spf13/viper"
	"github.com/spf13/cobra"
)

// setCmd represents the set command
var scopeSetCmd = &cobra.Command{
	Use:   "set",
	Short: "comma separated valid scopes",
	Long: `set the active scope in the config file. this option is 'scope'`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		split_scopes := strings.Split(args[0], ",")
		viper.Set("scope", split_scopes)
		viper.WriteConfig()
		fmt.Println("scope set to ", viper.Get("scope"))
	},
}

func init() {
	scopeCmd.AddCommand(scopeSetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
