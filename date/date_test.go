package date_test

import (
	d "journal/date"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestIsDate_WHEN_ReceivedDateSlashFormat_THEN_RetursTrue(t *testing.T) {
	//arrange -> Given that we try to check if a DD/MM/YYYY string is a date
	inputDate := "01/01/1982"
	expected := true

	//act -> When we check if it is a date
	result := d.IsDate(inputDate)

	//assert -> Then we expect the result to be true
	if diff := cmp.Diff(expected, result); diff != "" {
		t.Fatalf("Unexpected result (-want +got):\n%s", diff)
	}
}

func TestIsDate_WHEN_ReceivedDateSlashFormatWithOutOfBoundsDate_THEN_RetursFalse(t *testing.T) {
	//arrange -> Given that we try to check if a DD/MM/YYYY string with out of bounds numbers is a date
	inputDate := "34/16/1982"
	expected := false

	//act -> When we check if it is a date
	result := d.IsDate(inputDate)

	//assert -> Then we expect the result to be false
	if diff := cmp.Diff(expected, result); diff != "" {
		t.Fatalf("Unexpected result (-want +got):\n%s", diff)
	}
}

func TestIsDate_WHEN_ReceivedDateStrokeFormat_THEN_RetursTrue(t *testing.T) {
	//arrange -> Given that we try to check if a DD-MM-YYYY string is a date
	inputDate := "16-01-2017"
	expected := true

	//act -> When we check if it is a date
	result := d.IsDate(inputDate)

	//assert -> Then we expect the result to be true
	if diff := cmp.Diff(expected, result); diff != "" {
		t.Fatalf("Unexpected result (-want +got):\n%s", diff)
	}
}

func TestIsDate_WHEN_ReceivedDateStrokeFormatWithOutOfBoundsDate_THEN_RetursFalse(t *testing.T) {
	//arrange -> Given that we try to check if a DD/MM/YYYY string is a date
	inputDate := "32-16-2017"
	expected := false

	//act -> When we check if it is a date
	result := d.IsDate(inputDate)

	//assert -> Then we expect the result to be true
	if diff := cmp.Diff(expected, result); diff != "" {
		t.Fatalf("Unexpected result (-want +got):\n%s", diff)
	}
}

func TestIsDate_WHEN_ReceivedDateValidTextStringFormatTODAY_THEN_RetursTrue(t *testing.T) {
	//arrange -> Given that we try to check if text string is a date
	inputDate := "today"
	expected := true

	//act -> When we check if it is a date
	result := d.IsDate(inputDate)

	//assert -> Then we expect the result to be true
	if diff := cmp.Diff(expected, result); diff != "" {
		t.Fatalf("Unexpected result (-want +got):\n%s", diff)
	}
}

func TestIsDate_WHEN_ReceivedDateValidTextStringFormatYESTERDAY_THEN_RetursTrue(t *testing.T) {
	//arrange -> Given that we try to check if text string is a date
	inputDate := "yesterday"
	expected := true

	//act -> When we check if it is a date
	result := d.IsDate(inputDate)

	//assert -> Then we expect the result to be true
	if diff := cmp.Diff(expected, result); diff != "" {
		t.Fatalf("Unexpected result (-want +got):\n%s", diff)
	}
}

func TestIsDate_WHEN_ReceivedDateInvalidTextStringFormat_THEN_RetursFalse(t *testing.T) {
	//arrange -> Given that we try to check if an invalid text string is a date
	inputDate := "GIBBERISH"
	expected := false

	//act -> When we check if it is a date
	result := d.IsDate(inputDate)

	//assert -> Then we expect the result to be true
	if diff := cmp.Diff(expected, result); diff != "" {
		t.Fatalf("Unexpected result (-want +got):\n%s", diff)
	}
}

var now = time.Now()
var today, _ = time.Parse(d.SlashDDMMYY, now.Format(d.SlashDDMMYY))
var yesterday, _ = time.Parse(d.SlashDDMMYY, now.AddDate(0, 0, -1).Format(d.SlashDDMMYY))

func TestDateParser_WHEN_ReceivedEmptyString_THEN_returnsToday(t *testing.T) {
	//arrange -> Given that we try to parse an empty string
	inputDate := ""
	expected := today.Format(d.SlashDDMMYY)

	//act -> When we parse the date
	result := d.DateParse(inputDate)

	//assert -> Then we expect the result to be today
	if diff := cmp.Diff(expected, result); diff != "" {
		t.Fatalf("Unexpected result (-want +got):\n%s", diff)
	}
}

func TestDateParser_WHEN_ReceivedToday_THEN_returnsToday(t *testing.T) {
	//arrange -> Given that we try to parse today
	inputDate := "today"
	expected := today.Format(d.SlashDDMMYY)

	//act -> When we parse the date
	result := d.DateParse(inputDate)

	//assert -> Then we expect the result to be today
	if diff := cmp.Diff(expected, result); diff != "" {
		t.Fatalf("Unexpected result (-want +got):\n%s", diff)
	}
}

func TestDateParser_WHEN_ReceivedYesterday_THEN_returnsYesterday(t *testing.T) {
	//arrange -> Given that we try to parse yesterday
	inputDate := "yesterday"
	expected := yesterday.Format(d.SlashDDMMYY)

	//act -> When we parse the date
	result := d.DateParse(inputDate)

	//assert -> Then we expect the result to be yesterday
	if diff := cmp.Diff(expected, result); diff != "" {
		t.Fatalf("Unexpected result (-want +got):\n%s", diff)
	}
}

func TestDateParser_WHEN_ReceivedSlashDDMMYYYY_THEN_returnsSlashDDMMYYYY(t *testing.T) {
	//arrange -> Given that we try to parse a slash DD/MM/YYYY date
	inputDate := "02/02/2020"
	expected := "02/02/2020"

	//act -> When we parse the date
	result := d.DateParse(inputDate)

	//assert -> Then we expect the result to be the same as the input
	if diff := cmp.Diff(expected, result); diff != "" {
		t.Fatalf("Unexpected result (-want +got):\n%s", diff)
	}
}

func TestDateParser_WHEN_ReceivedDashDDMMYYYY_THEN_returnsSlashDDMMYYYY(t *testing.T) {
	//arrange -> Given that we try to parse a dash DD-MM-YYYY date
	inputDate := "02-02-2020"
	expected := "02/02/2020"

	//act -> When we parse the date
	result := d.DateParse(inputDate)

	//assert -> Then we expect the result to be the same as the input
	if diff := cmp.Diff(expected, result); diff != "" {
		t.Fatalf("Unexpected result (-want +got):\n%s", diff)
	}
}

func TestDateParser_WHEN_ReceivedRandomText_THEN_returnsToday(t *testing.T) {
	//arrange -> Given that we try to parse a random text
	inputDate := "GIBBERISH"
	expected := today.Format(d.SlashDDMMYY)

	//act -> When we parse the date
	result := d.DateParse(inputDate)

	//assert -> Then we expect the result to be today
	if diff := cmp.Diff(expected, result); diff != "" {
		t.Fatalf("Unexpected result (-want +got):\n%s", diff)
	}
}
