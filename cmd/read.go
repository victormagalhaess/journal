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
	"journal/stream"

	"github.com/spf13/cobra"
)

// readCmd represents the read command
var readCmd = &cobra.Command{
	Use:   "read",
	Short: "read - loads an entry from the journal",
	Long: `The read command is used to retrieve an entry from the journal. 
It may or may not take a date parameter.
If no date parameter is passed the today's entry will be read.
The date parameter can be nothing, a DD/MM/YYYY date, "yesterday" or "today".
Usage:
	journal read 
	journal read 14/2/2022
	journal read today`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			log.Fatal("Error! journal add must receive 0 or 1 parameters, try journal read --help to see more.\n")
		}
		ReadEntries(args)
	},
}

func init() {
	rootCmd.AddCommand(readCmd)
}

func ReadEntries(args []string) {
	var date string
	if len(args) == 0 {
		date = d.DateParse("")
	} else {
		date = d.DateParse(args[0])
	}
	entries := repository.ReadEntries(date)
	stream.OutputEntries(entries)
}
