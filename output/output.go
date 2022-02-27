package output

import (
	"journal/log"
	"journal/repository"
)

func OutputEntries(entries []repository.Entry) {
	log.Success("Your entries are:\n\n")
	for _, entry := range entries {
		log.Printf("Date: %s\n", entry.Date)
		log.Printf("  Hash: %s\n", entry.Text.Hash)
		log.Printf("  Message: %s\n\n", entry.Text.Value)
	}
}
