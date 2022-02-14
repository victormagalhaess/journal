package repository

import (
	"io/ioutil"
	"journal/flatfile"
)

type Text struct {
	Value     string `yaml:"value"`
	Hash      string `yaml:"hash"`
	UniqueArt string `yaml:"uniqueArt"`
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

func clean() error {
	file, err := flatfile.Open()
	defer flatfile.Close(file)
	if err != nil {
		return err
	}
	if err := file.Truncate(0); err != nil {
		return err
	}
	return nil
}
