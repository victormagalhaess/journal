package repository

import (
	"journal/art"
	"journal/log"
	"strings"
	"time"

	"journal/crypto"

	"gopkg.in/yaml.v2"
)

const (
	separator          = "# end\n"
	successAddingEntry = "Entry added with success!\n"
	errorAddingEntry   = "Error: An error occurred while adding the entry!\n"
	errorReadingFile   = "Error: An error occurred while reading the file!\n"
	errorUnparsing     = "Error: An error occurred while unparsing the entries!\n"
	errorCleaningFile  = "Error: An error occurred while cleaning the file!\n"
	noEntriesFound     = "No entries found!\n"
)

func NewEntry(date string, text string) {
	hash, err := crypto.GenerateHash(crypto.Time{Time: time.Now()})
	if err != nil {
		log.Fatal(errorAddingEntry)
		return
	}

	uniqueArt := art.BuildAsciiArt(hash)

	entry := &Entry{
		Date: date,
		Text: Text{
			Value:     text,
			Hash:      hash,
			UniqueArt: uniqueArt,
		},
	}
	addEntry(*entry)
	log.Success(successAddingEntry)
}

func addEntry(entry Entry) {
	d, err := yaml.Marshal(&entry)
	if err != nil {
		log.Fatal(errorAddingEntry)
		return
	}
	data := string(d) + separator
	add(data)
}

func verifyError(err error, message string) {
	if err != nil {
		log.Fatal(message)
	}
}

func parseEntries(unparsedEntries []string) []Entry {
	var (
		entries = make([]Entry, 0)
		entry   = &Entry{}
	)
	for _, item := range unparsedEntries {
		err := yaml.Unmarshal([]byte(item), &entry)
		verifyError(err, errorUnparsing)
		entries = append(entries, *entry)
	}
	return entries
}

func ReadEntries(date string) []Entry {
	file, err := read()
	verifyError(err, errorReadingFile)
	unparsedEntries := strings.Split(file, separator)
	if len(unparsedEntries) == 0 {
		log.Warning(noEntriesFound)
		return nil
	}
	unparsedEntries = unparsedEntries[:len(unparsedEntries)-1]
	return parseEntries(unparsedEntries)
}

func CleanFile() {
	if err := clean(); err != nil {
		log.Fatal(errorCleaningFile)
	}
}

func getProcessedEntries(entries []Entry, date string) []Entry {
	var processedEntries []Entry
	for _, entry := range entries {
		if entry.Date == date {
			continue
		}
		processedEntries = append(processedEntries, entry)
	}
	return processedEntries
}

func addEntries(entries []Entry) {
	for _, entry := range entries {
		addEntry(entry)
	}
}

func DeleteEntry(key string, isDate bool) {
	entries := ReadEntries("")
	processedEntries := getProcessedEntries(entries, key)
	CleanFile()
	addEntries(processedEntries)
}
