/*
Copyright © 2023 Peeanio

*/
package cmd

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/spf13/cobra"
	"os"
)

// setCmd represents the set command
var characterSetCmd = &cobra.Command{
	Use:   "set",
	Short: "set the active character id in the config file",
	Long: `set the active character id in the config file. this option is 'character_id'`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		viper.Set("character_id", args[0])
		err := viper.WriteConfig()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Write config failed: ", err)
		}
		fmt.Println("character_id set to ", args[0])
	},
}

func init() {
	characterCmd.AddCommand(characterSetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
