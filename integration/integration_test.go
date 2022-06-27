package integration_test

import (
	"journal/cmd"
	"journal/flatfile"
	"journal/log"
	"journal/repository"
	"os"
	"strings"
	"testing"
)

var (
	entryMessage1 string = "Today I waste 15 minutes thinking in an entry example, :("
	entryMessage2 string = "Today is Luis birthday, I'm so happy!"
)

func deleteFile() {
	e := os.Remove("journal.yaml")
	if e != nil {
		log.Print(e.Error())
	}
}

func addSampleEntry1() {
	message := []string{"16-10-2000", entryMessage1}
	cmd.AddEntry(message)
}

func addSampleEntry2() {
	message := []string{"17-03-2001", entryMessage2}
	cmd.AddEntry(message)
}

func TestFileCreation(t *testing.T) {
	deleteFile()
	addSampleEntry1()
	if isReady, _ := flatfile.IsReady(); !isReady {
		t.Error("The flatfile database should be ready")
	}
	deleteFile()
}

func TestCheckAddingEntries(t *testing.T) {
	deleteFile()
	for i := 0; i < 5; i++ {
		addSampleEntry1()
	}

	entries := repository.ReadEntries("")
	if len(entries) != 5 {
		t.Error("The number of entries should be 5")
	}

	if entries[0].Text.Value != "Today I waste 15 minutes thinking in an entry example, :(" {
		t.Error("The first entry should be the one we added")
	}
	deleteFile()
}

func TestCheckDeleteEntries(t *testing.T) {
	deleteFile()
	for i := 0; i < 5; i++ {
		addSampleEntry1()
	}
	addSampleEntry2()
	command := []string{"16-10-2000"}
	cmd.DeleteEntry(command)
	entries := repository.ReadEntries("")
	if len(entries) != 1 {
		t.Error("The number of entries should be 0")
	}
	if entries[0].Text.Value != "Today is Luis birthday, I'm so happy!" {
		t.Error("The first entry should be the one we added")
	}
	deleteFile()
}

func TestCheckDumpEntries(t *testing.T) {
	deleteFile()
	addSampleEntry1()
	addSampleEntry2()
	cmd.DumpEntries()
	outputContainsExpected := strings.Contains(log.GetBuffer(), entryMessage1) && strings.Contains(log.GetBuffer(), entryMessage2)
	if !outputContainsExpected {
		t.Error("The output should contain the entries we added")
	}
	deleteFile()
}

func TestCheckReadEntries(t *testing.T) {
	deleteFile()
	addSampleEntry1()
	addSampleEntry2()
	var command1 = []string{"16-10-2000"}
	var command2 = []string{"17-03-2001"}
	cmd.ReadEntries(command1)
	if !strings.Contains(log.GetBuffer(), entryMessage1) {
		t.Error("The output should contain the entry we added")
	}

	cmd.ReadEntries(command2)
	if !strings.Contains(log.GetBuffer(), entryMessage2) {
		t.Error("The output should contain the entry we added")
	}
	deleteFile()
}
