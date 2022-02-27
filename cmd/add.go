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

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add - creates an entry in the journal",
	Long: `The add command is used to create a new entry in the journal. 
It may or may not take a date parameter and the text used.
If no date parameter is passed the entry will be saved as todays entry.
The date parameter can be nothing, a DD/MM/YYYY date, a DD-MM-YYYY date, "yesterday" or "today".
Usage:
journal add "Today I waste 15 minutes thinking in an entry example, :("
journal add 14/2/2022 "Oh my god, I totally forgot to write about how normal 14/2/2022 was!"
journal add today "I'm not thirsty at all today."`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 || len(args) > 2 {
			log.Fatal("Error! journal add must receive 1 or 2 parameters, try journal add --help to see more.\n")
		}
		addEntry(args)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

// addEntry process the console entry and call the repository.AddEntry to save the data.
func addEntry(args []string) {
	var date string
	message := ""
	if len(args) == 1 {
		date = d.DateParse("")
		message = args[0]
	} else {
		date = d.DateParse(args[0])
		message = args[1]
	}
	repository.AddEntry(date, message)
}
