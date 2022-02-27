package repository

import (
	"journal/log"
	"strings"

	"journal/crypto"

	"gopkg.in/yaml.v2"
)

const (
	Separator = "# end\n"
)

func AddEntry(date string, text string) {
	entry := &Entry{
		Date: date,
		Text: Text{
			Value: text,
			Hash:  crypto.GenerateHash(),
		},
	}
	d, err := yaml.Marshal(&entry)
	if err != nil {
		log.Fatal("Error: An error occurred saving the entry!\n")
		return
	}
	data := string(d) + Separator
	add(data)
	log.Success("Entry added with success!\n")
}

func ReadEntries(date string) []Entry {
	var (
		entries = make([]Entry, 0)
		entry   = &Entry{}
	)
	file, err := read()
	if err != nil {
		log.Fatal("Error: An error occurred reading the flatfile!\n")
	}
	unparsedEntries := strings.Split(file, Separator)
	if len(unparsedEntries) == 0 {
		log.Warning("No entries found!\n")
		return nil
	}
	unparsedEntries = unparsedEntries[:len(unparsedEntries)-1]
	for _, item := range unparsedEntries {
		err = yaml.Unmarshal([]byte(item), &entry)
		if err != nil {
			log.Fatal("Error: An error occurred unparsing the entries!\n")
		}
		if entry.Date == date {
			entries = append(entries, *entry)
		}
	}

	return entries
}
