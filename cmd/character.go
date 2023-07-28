/*
Copyright © 2023 Peeanio

*/
package cmd

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/spf13/cobra"
)

// characterCmd represents the character command
var characterCmd = &cobra.Command{
	Use:   "character",
	Short: "interacts with the selected character",
	Long: `Many API actions require setting a character id to based information off of. This command deals with that. Invoking with no subcommands returns the selected character id`,
	Run: func(cmd *cobra.Command, args []string) {
		character_id := viper.GetString("character_id")
		fmt.Println(character_id)
	},
}

func init() {
	rootCmd.AddCommand(characterCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// characterCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// characterCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
