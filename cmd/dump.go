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

	"github.com/spf13/cobra"
)

// dumpCmd represents the dump command
var dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "dump - dump outputs all entries of your journal",
	Long: `The dump command is used to output all entries registered on your journal.
It does not receive any parameter.
Usage:
	journal dump`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			log.Fatal("Error: journal dump must not receive any parameter!\n")
		}
		DumpEntries()
	},
}

func init() {
	rootCmd.AddCommand(dumpCmd)
}

func DumpEntries() {
	entries := repository.ReadEntries("")
	stream.OutputEntries(entries)
}
