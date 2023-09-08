/*
Copyright © 2023 Peeanio

*/
package cmd

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/spf13/cobra"
)

// scopeCmd is for setting api scope
var scopeCmd = &cobra.Command{
	Use:   "scope",
	Short: "shows the scope login will use",
	Long: `Scope is important to let the api know what access level to give the tool`,
	Run: func(cmd *cobra.Command, args []string) {
		scope := viper.GetStringSlice("scope")
		fmt.Println(scope)
	},
}

func init() {
	rootCmd.AddCommand(scopeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// scopeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// scopeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
