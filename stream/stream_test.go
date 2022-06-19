package stream_test

import (
	"journal/log"
	"journal/repository"
	"journal/stream"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestOutputEntries(t *testing.T) {
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
	stream.OutputEntries(entries)
	expected := "\x1b[97mYour entries are:\n\nDate: 2020-01-01\n  Hash: 3766189190\n  Message: Hello World\n\nDate: 2020-01-02\n  Hash: 3766189190\n  Message: Hello World\n\n\x1b[0m"
	if diff := cmp.Diff(expected, log.GetBuffer()); diff != "" {
		t.Fatalf("Unexpected result (-want +got):\n%s", diff)
	}
}

func TestIsYes(t *testing.T) {
	if !stream.IsYes("y") {
		t.Fatal("Expected true")
	}
	if !stream.IsYes("yes") {
		t.Fatal("Expected true")
	}
	if stream.IsYes("n") {
		t.Fatal("Expected false")
	}
	if stream.IsYes("no") {
		t.Fatal("Expected false")
	}
}

func TestIsNo(t *testing.T) {
	if !stream.IsNo("n") {
		t.Fatal("Expected true")
	}
	if !stream.IsNo("no") {
		t.Fatal("Expected true")
	}
	if stream.IsNo("y") {
		t.Fatal("Expected false")
	}
	if stream.IsNo("yes") {
		t.Fatal("Expected false")
	}
}

func TestAskPermission(t *testing.T) {
	if !stream.AskPermission("", strings.NewReader("y\n")) {
		t.Fatal("Expected true")
	}
	if !stream.AskPermission("", strings.NewReader("yes\n")) {
		t.Fatal("Expected true")
	}
	if stream.AskPermission("", strings.NewReader("n\n")) {
		t.Fatal("Expected false")
	}
	if stream.AskPermission("", strings.NewReader("no\n")) {
		t.Fatal("Expected false")
	}
}
