package repository

import (
	"log"

	"journal/crypto"

	"gopkg.in/yaml.v2"
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
		log.Fatalf("error: %v", err)
		return
	}

	append(string(d))
}
