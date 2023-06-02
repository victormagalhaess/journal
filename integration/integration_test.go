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

func SetupTest() func() {
	//setup
	if isReady, _ := flatfile.IsReady(); isReady {
		deleteFile()
	}
	//teardown
	return func() { deleteFile() }
}

func TestFileCreation(t *testing.T) {
	defer SetupTest()()
	addSampleEntry1()
	if isReady, _ := flatfile.IsReady(); !isReady {
		t.Error("The flatfile database should be ready")
	}
}

func TestCheckAddingEntries(t *testing.T) {
	//arrange
	defer SetupTest()()
	for i := 0; i < 5; i++ {
		addSampleEntry1()
	}
	//act
	entries := repository.ReadEntries("")

	//assert
	if len(entries) != 5 {
		t.Error("The number of entries should be 5")
	}

	if entries[0].Text.Value != "Today I waste 15 minutes thinking in an entry example, :(" {
		t.Error("The first entry should be the one we added")
	}
}

func TestCheckDeleteEntries(t *testing.T) {
	//arrange
	defer SetupTest()()
	for i := 0; i < 5; i++ {
		addSampleEntry1()
	}
	addSampleEntry2()
	command := []string{"16-10-2000"}
	cmd.DeleteEntry(command)

	//act
	entries := repository.ReadEntries("")

	//assert
	if len(entries) != 1 {
		t.Error("The number of entries should be 0")
	}
	if entries[0].Text.Value != "Today is Luis birthday, I'm so happy!" {
		t.Error("The first entry should be the one we added")
	}
}

func TestCheckDumpEntries(t *testing.T) {
	//arrange
	defer SetupTest()()
	addSampleEntry1()
	addSampleEntry2()

	//act
	cmd.DumpEntries()
	outputContainsExpected := strings.Contains(log.GetBuffer(), entryMessage1) && strings.Contains(log.GetBuffer(), entryMessage2)

	//assert
	if !outputContainsExpected {
		t.Error("The output should contain the entries we added")
	}
}

func TestCheckReadEntries(t *testing.T) {
	//arrange
	defer SetupTest()()
	addSampleEntry1()
	addSampleEntry2()
	var command1 = []string{"16-10-2000"}

	//act
	cmd.ReadEntries(command1)

	//assert
	if !strings.Contains(log.GetBuffer(), entryMessage1) {
		t.Error("The output should contain the entry we added")
	}

}
