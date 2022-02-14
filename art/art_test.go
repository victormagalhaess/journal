package art_test

import (
	"journal/art"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBuildAsciiArt_WHEN_GetValidHash_THEN_returns_string(t *testing.T) {
	//arrange -> Given that we try to generate ascii art from a valid hash
	input := "3766189190"
	expected := "/%##*O8*8@"

	//act -> When we generate the ascii art
	result := art.BuildAsciiArt(input)

	//assert -> Then we expect the result to be the ascii art
	if diff := cmp.Diff(expected, result); diff != "" {
		t.Fatalf("Unexpected result (-want +got):\n%s", diff)
	}
}

func TestBuildAsciiArt_WHEN_GetEmptyHas_THEN_returns_emptyString(t *testing.T) {
	//arrange -> Given that we try to generate ascii art from an empty hash
	input := ""
	expected := ""

	//act -> When we generate the ascii art
	result := art.BuildAsciiArt(input)

	//assert -> Then we expect the result to be the ascii art
	if diff := cmp.Diff(expected, result); diff != "" {
		t.Fatalf("Unexpected result (-want +got):\n%s", diff)
	}
}

func TestBuildAsciiArt_WHEN_GetInvalidHash_THEN_returns_emptyString(t *testing.T) {
	//arrange -> Given that we try to generate ascii art from an invalid hash
	input := "invalid"
	expected := ""

	//act -> When we generate the ascii art
	result := art.BuildAsciiArt(input)

	//assert -> Then we expect the result to be the ascii art
	if diff := cmp.Diff(expected, result); diff != "" {
		t.Fatalf("Unexpected result (-want +got):\n%s", diff)
	}
}
