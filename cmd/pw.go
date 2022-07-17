/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"main/haveibeenpwned"

	"github.com/spf13/cobra"
)

// pwCmd represents the pw command
var pwCmd = &cobra.Command{
	Use:     "pw",
	Aliases: []string{"password"},
	Short:   "Check if a password has been exposed in a data breach.",
	Long:    `Check if a password has been exposed in a data breach, it will return the number of a times the provided password has been exposed.`,
	Example: `pwned pw <password>`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			fmt.Println("Please provide a password.")
			return
		}

		password := args[0]

		passwordCount, err := haveibeenpwned.PwnedPassword(password)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("%s has been exposed %d times\n", password, passwordCount)

	},
}

func init() {
	rootCmd.AddCommand(pwCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pwCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pwCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
