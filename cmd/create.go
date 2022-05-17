/*
Copyright © 2022 theo-commits <me@tlindsey.cloud>
This file is part of CLI application cli-tool
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create option allows you to dynamically create a file",
	Long: `This can be any type of file and also takes a path. Takes only a single argument
      For example:
"file.txt" # Saved in the current directory by default. 
"/home/theo-commits/file.txt"
WARNING ----- If the file already exists it will be overwritten with an empty file.
While we could provide an 'exists' module or simply handle this bug we won't feel free to submit a PR ;) `,
	Run: func(cmd *cobra.Command, args []string) {
		input := args[0]
		_, err := os.Create(input)
		if err != nil {
			fmt.Println("File does not exist.")
		}
		fmt.Printf("You've just created a file called, %s", input)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
