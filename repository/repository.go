package repository

import (
	"io/ioutil"
	"journal/flatfile"
)

type Text struct {
	Value string `yaml:"value"`
	Hash  string `yaml:"hash"`
}

type Entry struct {
	Date string `yaml:"date"`
	Text Text
}

func read() (string, error) {
	file, err := flatfile.Open()
	defer flatfile.Close(file)
	if err != nil {
		return "", err
	}
	content, err := ioutil.ReadAll(file)
	return string(content), err
}

func add(entry string) error {
	file, err := flatfile.Open()
	defer flatfile.Close(file)
	if err != nil {
		return err
	}
	if _, err := file.Write([]byte(entry)); err != nil {
		return err
	}
	return nil
}
