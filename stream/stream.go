package stream

import (
	"bufio"
	"io"
	"journal/log"
	"journal/repository"
	"strings"
)

func OutputEntries(entries []repository.Entry) {
	log.Success("Your entries are:\n\n")
	for _, entry := range entries {
		log.Printf("Date: %s\n", entry.Date)
		log.Printf("  Hash: %s\n", entry.Text.Hash)
		log.Printf("  Message: %s\n\n", entry.Text.Value)
	}
}

func AskPermission(message string, in io.Reader) bool {
	reader := bufio.NewReader(in)
	for {
		log.Warningf("%s [y/n]: ", message)
		response, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("Error: An error occurred while reading your answer.")
		}
		response = strings.ToLower(strings.TrimSpace(response))
		if response == "y" || response == "yes" {
			return true
		} else if response == "n" || response == "no" {
			return false
		}
	}
}
