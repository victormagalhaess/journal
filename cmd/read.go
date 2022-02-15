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
	"fmt"

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
	Example:
	journal read 
	journal read 14/2/2022
	journal read today`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("read called")
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