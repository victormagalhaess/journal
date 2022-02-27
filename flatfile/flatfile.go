package flatfile

import (
	"errors"
	"os"
)

const (
	path = "journal.yaml"
	// the flatfile db must be readable for the user and owner, refer to: https://en.wikipedia.org/wiki/Chmod#Numerical_permissions
	fileMode    = 0644
	permissions = os.O_APPEND | os.O_CREATE | os.O_RDWR
)

func IsReady() (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}

func Open() (*os.File, error) {
	file, err := os.OpenFile(path, permissions, fileMode)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func Close(file *os.File) error {
	if err := file.Close(); err != nil {
		return err
	}
	return nil
}
