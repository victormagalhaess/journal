package crypto

import (
	"fmt"
	"hash/fnv"
	"time"
)

type Time struct {
	Time time.Time
}

func GenerateHash(time Time) (string, error) {
	if time.Time.IsZero() {
		return "", fmt.Errorf("invalid time")
	}
	timestamp := fmt.Sprint(time.Time.Unix())
	hash := fnv.New32a()
	hash.Write([]byte(timestamp))
	return fmt.Sprint(hash.Sum32()), nil
}
