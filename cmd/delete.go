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
	d "journal/date"
	"journal/log"
	"journal/repository"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete - erases an entry from the journal",
	Long: `The delete command is used to erase an entry from the journal. 
It may or may not take a date parameter.
If no date parameter is passed the today's entry will be deleted.
The date parameter can be nothing, a DD/MM/YYYY date, "yesterday" or "today".
You can also specify a hash to delete a single entry.
Usage:
	journal delete 
	journal delete 14/2/2022
	journal delete today
	journal delete 2966741453`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			log.Fatal("Error! journal delete must receive 0 or 1 parameters, try journal delete --help to see more.\n")
		}
		DeleteEntry(args)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

func DeleteEntry(args []string) {
	key := args[0]
	isDate := false
	if isDate = d.IsDate(key); isDate {
		key = d.DateParse(args[0])
		log.Warningf("Deleting all entries of the date: %v\n", key)
	} else {
		log.Warningf("Deleting the entry with Hash: %v\n", key)
	}
	repository.DeleteEntry(key, isDate)
}
