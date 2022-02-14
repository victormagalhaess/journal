package stream

import (
	"bufio"
	"fmt"
	"io"
	"journal/log"
	"journal/repository"
	"strings"
)

func OutputEntries(entries []repository.Entry) {
	out := "Your entries are:\n\n"

	for _, entry := range entries {
		out += fmt.Sprintf("Date: %s\n", entry.Date)
		out += fmt.Sprintf("  Hash: %s\n", entry.Text.Hash)
		out += fmt.Sprintf("  Message: %s\n\n", entry.Text.Value)
	}

	log.Print(out)
}

func IsYes(response string) bool {
	response = strings.ToLower(strings.TrimSpace(response))
	return response == "y" || response == "yes"
}

func IsNo(response string) bool {
	return !IsYes(response)
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
		if IsYes(response) {
			return true
		} else if IsNo(response) {
			return false
		}
	}
}
