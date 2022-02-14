package crypto_test

import (
	"journal/crypto"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestGenerateHash_WHEN_GetValidTime_THEN_returnsHash(t *testing.T) {
	//arrange -> Given that we try to generate a hash from a valid time
	nowTime, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	expected := "3766189190"

	//act -> When we generate the hash
	result, err := crypto.GenerateHash(crypto.Time{Time: nowTime})

	//assert -> Then we expect the result to be the hash
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
	if diff := cmp.Diff(expected, result); diff != "" {
		t.Fatalf("Unexpected result (-want +got):\n%s", diff)
	}
}

func TestGenerateHash_WHEN_GetInvalidTime_THEN_returnsErr(t *testing.T) {
	//arrange -> Given that we try to generate a hash from an invalid time
	expectedErr := "invalid time"

	//act -> When we generate the hash
	_, err := crypto.GenerateHash(crypto.Time{Time: time.Time{}})

	//assert -> Then we expect the result to be the hash
	if err == nil {
		t.Fatalf("Expected error: %s", expectedErr)
	}
	if diff := cmp.Diff(expectedErr, err.Error()); diff != "" {
		t.Fatalf("Unexpected result (-want +got):\n%s", diff)
	}
}
