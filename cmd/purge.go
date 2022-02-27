/*
Copyright © 2022 Victor Magalhães hello@victordias.dev

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
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
		if len(args) > 0 {
			log.Fatal("Error: journal purge must not receive any parameter!\n")
		}
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
