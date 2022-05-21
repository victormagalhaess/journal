package date_test

import (
	d "journal/date"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

var testsIsDate = []struct {
	inputDate string
	expected  bool
}{
	{
		inputDate: "NONSENSE",
		expected:  false,
	},
	{
		inputDate: "01/01/0001",
		expected:  true,
	},
	{
		inputDate: "16/01/2017",
		expected:  true,
	},
	{
		inputDate: "01-01-0001",
		expected:  true,
	},
	{
		inputDate: "16-01-2017",
		expected:  true,
	},
	{
		inputDate: "01-16-2017",
		expected:  false,
	},
	{
		inputDate: "01/16/2017",
		expected:  false,
	},
	{
		inputDate: "today",
		expected:  true,
	},
	{
		inputDate: "yesterday",
		expected:  true,
	},
}

func TestIsDate(t *testing.T) {
	for _, test := range testsIsDate {
		t.Log(test.inputDate)
		result := d.IsDate(test.inputDate)
		if diff := cmp.Diff(test.expected, result); diff != "" {
			t.Fatalf("Unexpected result (-want +got):\n%s", diff)
		}
	}
}

var now = time.Now()
var today, _ = time.Parse(d.SlashDDMMYY, now.Format(d.SlashDDMMYY))
var yesterday, _ = time.Parse(d.SlashDDMMYY, now.AddDate(0, 0, -1).Format(d.SlashDDMMYY))
var testsDateParser = []struct {
	unparsed string
	expected string
}{
	{
		unparsed: "",
		expected: today.Format(d.SlashDDMMYY),
	},
	{
		unparsed: "NONSENSE",
		expected: today.Format(d.SlashDDMMYY),
	},
	{
		unparsed: "today",
		expected: today.Format(d.SlashDDMMYY),
	},
	{
		unparsed: "yesterday",
		expected: yesterday.Format(d.SlashDDMMYY),
	},
	{
		unparsed: "02/02/2020",
		expected: "02/02/2020",
	},
	{
		unparsed: "02-02-2020",
		expected: "02/02/2020",
	},
}

func TestDateParser(t *testing.T) {
	for _, test := range testsDateParser {
		t.Log(test.unparsed)
		result := d.DateParse(test.unparsed)
		if diff := cmp.Diff(test.expected, result); diff != "" {
			t.Fatalf("Unexpected result (-want +got):\n%s", diff)
		}
	}
}
