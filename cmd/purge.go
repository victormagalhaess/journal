/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"journal/log"
	"journal/repository"
	"journal/stream"
	"os"

	"github.com/spf13/cobra"
)

// purgeCmd represents the purge command
var purgeCmd = &cobra.Command{
	Use:   "purge",
	Short: "purge - cleans all entries on your journal",
	Long: `The purge command is used to delete all entries on your journal.
It does not receive any parameter.
Usage: 
journal purge`,
	Run: func(cmd *cobra.Command, args []string) {
		purgeEntries()
	},
}

func init() {
	rootCmd.AddCommand(purgeCmd)
}

func purgeEntries() {
	message := "Are you sure you want to purge all your journal?"
	if shouldPurge := stream.AskPermission(message, os.Stdin); shouldPurge {
		repository.CleanFile()
		log.Success("All your entries were deleted succesfully!\n")
	} else {
		log.Error("Error: operation aborted by user!\n")
	}
}
