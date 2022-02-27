package crypto

import (
	"fmt"
	"hash/fnv"
	"time"
)

func GenerateHash() string {
	timestamp := fmt.Sprint(time.Now().Unix())
	hash := fnv.New32a()
	hash.Write([]byte(timestamp))
	return fmt.Sprint(hash.Sum32())
}
