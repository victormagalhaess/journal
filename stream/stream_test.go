package stream_test

import (
	"journal/log"
	"journal/repository"
	"journal/stream"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestOutputEntry_WHEN_WithValidEntries_THEN_shouldLogTheEntries(t *testing.T) {
	//arrange -> Given that we try to output entries
	entries := []repository.Entry{
		{
			Date: "2020-01-01",
			Text: repository.Text{
				Value: "Hello World",
				Hash:  "3766189190",
			},
		},
		{
			Date: "2020-01-02",
			Text: repository.Text{
				Value: "Hello World",
				Hash:  "3766189190",
			},
		},
	}

	//act -> When we output the entries
	stream.OutputEntries(entries)

	//assert -> Then we expect the result to be the entries
	expected := "\x1b[97mYour entries are:\n\nDate: 2020-01-01\n  Hash: 3766189190\n  Message: Hello World\n\nDate: 2020-01-02\n  Hash: 3766189190\n  Message: Hello World\n\n\x1b[0m"
	if diff := cmp.Diff(expected, log.GetBuffer()); diff != "" {
		t.Fatalf("Unexpected result (-want +got):\n%s", diff)
	}
}

func TestOutputEntry_WHEN_WithEmptyEntries_THEN_shouldLogEmptyEntries(t *testing.T) {
	//arrange -> Given that we try to output empty entries
	entries := []repository.Entry{}

	//act -> When we output the empty entries
	stream.OutputEntries(entries)

	//assert -> Then we expect the result to be the entries
	expected := "\x1b[97mYour entries are:\n\n\x1b[0m"
	if diff := cmp.Diff(expected, log.GetBuffer()); diff != "" {
		t.Fatalf("Unexpected result (-want +got):\n%s", diff)
	}
}

func TestIsYes_WHEN_promptedWithY_THEN_returnsTrue(t *testing.T) {
	//arrange -> Given that we try to check if the user input is yes
	input := "y"

	//act -> When we check if the user input is yes
	result := stream.IsYes(input)

	//assert -> Then we expect the result to be true
	if !result {
		t.Fatalf("Expected true")
	}
}

func TestIsYes_WHEN_promptedWithYes_THEN_returnsTrue(t *testing.T) {
	//arrange -> Given that we try to check if the user input is yes
	input := "yes"

	//act -> When we check if the user input is yes
	result := stream.IsYes(input)

	//assert -> Then we expect the result to be true
	if !result {
		t.Fatalf("Expected true")
	}
}

func TestIsYes_WHEN_promptedWithAnythingElse_THEN_returnsFalse(t *testing.T) {
	//arrange -> Given that we try to check if the user input is yes
	input := "GIBBERISH"

	//act -> When we check if the user input is yes
	result := stream.IsYes(input)

	//assert -> Then we expect the result to be false
	if result {
		t.Fatalf("Expected false")
	}
}

func TestIsNo_WHEN_promptedWithN_THEN_returnsTrue(t *testing.T) {
	//arrange -> Given that we try to check if the user input is no
	input := "n"

	//act -> When we check if the user input is no
	result := stream.IsNo(input)

	//assert -> Then we expect the result to be true
	if !result {
		t.Fatalf("Expected true")
	}
}

func TestIsNo_WHEN_promptedWithNo_THEN_returnsTrue(t *testing.T) {
	//arrange -> Given that we try to check if the user input is no
	input := "no"

	//act -> When we check if the user input is no
	result := stream.IsNo(input)

	//assert -> Then we expect the result to be true
	if !result {
		t.Fatalf("Expected true")
	}
}

func TestIsNo_WHEN_promptedWithAnythingElse_THEN_returnsTrue(t *testing.T) {
	//arrange -> Given that we try to check if the user input is no
	input := "GIBBERISH"

	//act -> When we check if the user input is no
	result := stream.IsNo(input)

	//assert -> Then we expect the result to be false
	if !result {
		t.Fatalf("Expected false")
	}
}

func TestAskPermission_WHEN_promptedWithY_THEN_returnsTrue(t *testing.T) {
	//arrange -> Given that we try to ask the user for permission
	input := "y\n"

	//act -> When we ask the user for permission
	result := stream.AskPermission("", strings.NewReader(input))

	//assert -> Then we expect the result to be true
	if !result {
		t.Fatalf("Expected true")
	}
}

func TestAskPermission_WHEN_promptedWithYes_THEN_returnsTrue(t *testing.T) {
	//arrange -> Given that we try to ask the user for permission
	input := "yes\n"

	//act -> When we ask the user for permission
	result := stream.AskPermission("", strings.NewReader(input))

	//assert -> Then we expect the result to be true
	if !result {
		t.Fatalf("Expected true")
	}
}

func TestAskPermission_WHEN_promptedWithN_THEN_returnsFalse(t *testing.T) {
	//arrange -> Given that we try to ask the user for permission
	input := "n\n"

	//act -> When we ask the user for permission
	result := stream.AskPermission("", strings.NewReader(input))

	//assert -> Then we expect the result to be false
	if result {
		t.Fatalf("Expected false")
	}
}

func TestAskPermission_WHEN_promptedWithNo_THEN_returnsFalse(t *testing.T) {
	//arrange -> Given that we try to ask the user for permission
	input := "no\n"

	//act -> When we ask the user for permission
	result := stream.AskPermission("", strings.NewReader(input))

	//assert -> Then we expect the result to be false
	if result {
		t.Fatalf("Expected false")
	}
}

func TestAskPermission_WHEN_promptedWithAnythingElse_THEN_returnsFalse(t *testing.T) {
	//arrange -> Given that we try to ask the user for permission
	input := "GIBBERISH\n"

	//act -> When we ask the user for permission
	result := stream.AskPermission("", strings.NewReader(input))

	//assert -> Then we expect the result to be false
	if result {
		t.Fatalf("Expected false")
	}
}
