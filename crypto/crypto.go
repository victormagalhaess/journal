package crypto

import (
	"crypto/sha256"
	"fmt"
	"time"
)

func GenerateHash() string {

	timestamp := time.Now().Unix()
	return fmt.Sprintf("%x", sha256.Sum256([]byte(string(timestamp))))[:16]
}
