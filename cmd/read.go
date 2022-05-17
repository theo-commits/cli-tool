/*
Copyright © 2022 theo-commits <me@tlindsey.cloud>
This file is part of CLI application cli-tool
*/
package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

// readCmd represents the read command
var readCmd = &cobra.Command{
	Use:   "read",
	Short: "reads the name of a given file",
	Long:  `reads the file and prints the contents to stdout. Takes one argument`,
	Run: func(cmd *cobra.Command, args []string) {
		input := args[0]
		data, err := os.ReadFile(input)
		if err != nil {
			log.Fatal(err)
		}
		os.Stdout.Write(data)
	},
}

func init() {
	rootCmd.AddCommand(readCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// readCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// readCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
