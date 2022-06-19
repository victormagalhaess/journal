package crypto

import (
	"fmt"
	"hash/fnv"
	"time"
)

func GenerateHash(nowTime time.Time) string {
	timestamp := fmt.Sprint(nowTime.Unix())
	hash := fnv.New32a()
	hash.Write([]byte(timestamp))
	return fmt.Sprint(hash.Sum32())
}
