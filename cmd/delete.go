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

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a given file",
	Long: `This will delete your file. Use at your own risk, there is NOT a recover command of any type!
	Although, that does sound pretty dope, feel free to open a PR. ;)
	This command will also delete empty directories. Returns an error if the file does not exist.
	Like read and create, this only takes a single argument`,
	Run: func(cmd *cobra.Command, args []string) {
		input := args[0]
		if err := os.Remove(input); err != nil {
			fmt.Println("File does not exist, please try again.\n")
			fmt.Println("Optionally, you can create the file with the create command")

		} else {
			fmt.Printf("The file %s has been deleted!", input)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
