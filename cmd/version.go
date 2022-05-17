/*
Copyright © 2022 theo-commits <me@tlindsey.cloud>
This file is part of CLI application cli-tool
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the current version of the cli-tool",
	Long: `My software has the best versions, behold it now.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cli-tool v0.00001 -- HEAD")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
