package crypto_test

import (
	"journal/crypto"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestGenerateHash(t *testing.T) {
	nowTime, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	expected := "3766189190"
	result := crypto.GenerateHash(nowTime)
	if diff := cmp.Diff(expected, result); diff != "" {
		t.Fatalf("Unexpected result (-want +got):\n%s", diff)
	}
}
