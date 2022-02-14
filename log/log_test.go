package log_test

import (
	"bytes"
	"journal/log"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPrint_WHEN_SendingContentToPrint_THEN_ShouldSendTheContentToTheOutputBuffer(t *testing.T) {
	//arrange -> Given that we try to print a message
	oldIsWindows := log.IsWindows
	log.IsWindows = false
	defer func() { log.IsWindows = oldIsWindows }()
	output := "\033[97mtesting\033[0m"
	oldOut := log.Out
	log.Out = new(bytes.Buffer)
	defer func() { log.Out = oldOut }()

	//act -> When we print the message
	log.Print("testing")

	//assert -> Then we expect the message to be printed
	if diff := cmp.Diff(output, log.Out.(*bytes.Buffer).String()); diff != "" {
		t.Fatalf("Unexpected result (-want +got):\n%s", diff)
	}
}

func TestPrintf_WHEN_SendingContentToPrint_THEN_ShouldSendTheContentToTheOutputBuffer(t *testing.T) {
	//arrange -> Given that we try to print a message with a format string
	oldIsWindows := log.IsWindows
	log.IsWindows = false
	defer func() { log.IsWindows = oldIsWindows }()
	output := "\033[97mtesting123\033[0m"
	oldOut := log.Out
	log.Out = new(bytes.Buffer)
	defer func() { log.Out = oldOut }()
	//act -> When we print the message
	log.Printf("testing%v", 123)

	//assert -> Then we expect the message to be printed
	if diff := cmp.Diff(output, log.Out.(*bytes.Buffer).String()); diff != "" {
		t.Fatalf("Unexpected result (-want +got):\n%s", diff)
	}
}

func TestError_WHEN_SendingContentToPrint_THEN_ShouldSendTheContentToTheOutputBuffer(t *testing.T) {
	//arrange -> Given that we try to print an error message
	oldIsWindows := log.IsWindows
	log.IsWindows = false
	defer func() { log.IsWindows = oldIsWindows }()
	output := "\033[31mtesting\033[0m"
	oldOut := log.Out
	log.Out = new(bytes.Buffer)
	defer func() { log.Out = oldOut }()

	//act -> When we print the message
	log.Error("testing")

	//assert -> Then we expect the message to be printed
	if diff := cmp.Diff(output, log.Out.(*bytes.Buffer).String()); diff != "" {
		t.Fatalf("Unexpected result (-want +got):\n%s", diff)
	}
}

func TestErrorF_WHEN_SendingContentToPrint_THEN_ShouldSendTheContentToTheOutputBuffer(t *testing.T) {
	//arrange -> Given that we try to print an error message with a format string
	oldIsWindows := log.IsWindows
	log.IsWindows = false
	defer func() { log.IsWindows = oldIsWindows }()
	output := "\033[31mtesting123\033[0m"
	oldOut := log.Out
	log.Out = new(bytes.Buffer)
	defer func() { log.Out = oldOut }()

	//act -> When we print the message
	log.ErrorF("testing%v", 123)

	//assert -> Then we expect the message to be printed
	if diff := cmp.Diff(output, log.Out.(*bytes.Buffer).String()); diff != "" {
		t.Fatalf("Unexpected result (-want +got):\n%s", diff)
	}
}

func TestWarning_WHEN_SendingContentToPrint_THEN_ShouldSendTheContentToTheOutputBuffer(t *testing.T) {
	//arrange -> Given that we try to print a warning message
	oldIsWindows := log.IsWindows
	log.IsWindows = false
	defer func() { log.IsWindows = oldIsWindows }()
	output := "\033[33mtesting\033[0m"
	oldOut := log.Out
	log.Out = new(bytes.Buffer)
	defer func() { log.Out = oldOut }()

	//act -> When we print the message
	log.Warning("testing")

	//assert -> Then we expect the message to be printed
	if diff := cmp.Diff(output, log.Out.(*bytes.Buffer).String()); diff != "" {
		t.Fatalf("Unexpected result (-want +got):\n%s", diff)
	}
}

func TestWarningF_WHEN_SendingContentToPrint_THEN_ShouldSendTheContentToTheOutputBuffer(t *testing.T) {
	//arrange -> Given that we try to print a warning message with a format string
	oldIsWindows := log.IsWindows
	log.IsWindows = false
	defer func() { log.IsWindows = oldIsWindows }()
	output := "\033[33mtesting123\033[0m"
	oldOut := log.Out
	log.Out = new(bytes.Buffer)
	defer func() { log.Out = oldOut }()

	//act -> When we print the message
	log.Warningf("testing%v", 123)

	// assert -> Then we expect the message to be printed
	if diff := cmp.Diff(output, log.Out.(*bytes.Buffer).String()); diff != "" {
		t.Fatalf("Unexpected result (-want +got):\n%s", diff)
	}
}

func TestSuccess_WHEN_SendingContentToPrint_THEN_ShouldSendTheContentToTheOutputBuffer(t *testing.T) {
	//arrange -> Given that we try to print a success message
	oldIsWindows := log.IsWindows
	log.IsWindows = false
	defer func() { log.IsWindows = oldIsWindows }()
	output := "\033[32mtesting\033[0m"
	oldOut := log.Out
	log.Out = new(bytes.Buffer)
	defer func() { log.Out = oldOut }()

	//act -> When we print the message
	log.Success("testing")

	//assert -> Then we expect the message to be printed
	if diff := cmp.Diff(output, log.Out.(*bytes.Buffer).String()); diff != "" {
		t.Fatalf("Unexpected result (-want +got):\n%s", diff)
	}
}

func TestSuccessF_WHEN_SendingContentToPrint_THEN_ShouldSendTheContentToTheOutputBuffer(t *testing.T) {
	//arrange -> Given that we try to print a success message with a format string
	oldIsWindows := log.IsWindows
	log.IsWindows = false
	defer func() { log.IsWindows = oldIsWindows }()
	output := "\033[32mtesting123\033[0m"
	oldOut := log.Out
	log.Out = new(bytes.Buffer)
	defer func() { log.Out = oldOut }()

	//act -> When we print the message
	log.SuccessF("testing%v", 123)

	//assert -> Then we expect the message to be printed
	if diff := cmp.Diff(output, log.Out.(*bytes.Buffer).String()); diff != "" {
		t.Fatalf("Unexpected result (-want +got):\n%s", diff)
	}
}

func TestFatal_WHEN_SendingContentToPrint_THEN_ShouldSendTheContentToTheOutputBuffer(t *testing.T) {
	//arrange -> Given that we try to print a fatal message
	oldIsWindows := log.IsWindows
	log.IsWindows = false
	defer func() { log.IsWindows = oldIsWindows }()
	output := "\033[31mtesting\033[0m"
	outCode := 1
	oldOut := log.Out
	log.Out = new(bytes.Buffer)
	defer func() { log.Out = oldOut }()
	code := 0
	oldExit := log.Exit
	log.Exit = func(c int) { code = c }
	defer func() { log.Exit = oldExit }()

	//act -> When we print the message
	log.Fatal("testing")

	//assert -> Then we expect the message to be printed and the program to exit with code 1
	if diff := cmp.Diff(output, log.Out.(*bytes.Buffer).String()); diff != "" {
		t.Fatalf("Unexpected result (-want +got):\n%s", diff)
	}
	if diff := cmp.Diff(outCode, code); diff != "" {
		t.Fatalf("Unexpected result (-want +got):\n%s", diff)
	}
}

func TestFatalF_WHEN_SendingContentToPrint_THEN_ShouldSendTheContentToTheOutputBuffer(t *testing.T) {
	//arrange -> Given that we try to print a fatal message with a format string
	oldIsWindows := log.IsWindows
	log.IsWindows = false
	defer func() { log.IsWindows = oldIsWindows }()
	output := "\033[31mtesting123\033[0m"
	outCode := 1
	oldOut := log.Out
	log.Out = new(bytes.Buffer)
	defer func() { log.Out = oldOut }()
	code := 0
	oldExit := log.Exit
	log.Exit = func(c int) { code = c }
	defer func() { log.Exit = oldExit }()

	//act -> When we print the message
	log.FatalF("testing%v", 123)

	//assert -> Then we expect the message to be printed and the program to exit with code 1
	if diff := cmp.Diff(output, log.Out.(*bytes.Buffer).String()); diff != "" {
		t.Fatalf("Unexpected result (-want +got):\n%s", diff)
	}
	if diff := cmp.Diff(outCode, code); diff != "" {
		t.Fatalf("Unexpected result (-want +got):\n%s", diff)
	}
}

func TestPrint_WHEN_SendingContentToPrintAndIsWindows_THEN_ShouldSendTheContentToTheOutputBuffer(t *testing.T) {
	//arrange -> Given that we try to print a message on windows
	oldIsWindows := log.IsWindows
	log.IsWindows = true
	defer func() { log.IsWindows = oldIsWindows }()
	output := "testing"
	oldOut := log.Out
	log.Out = new(bytes.Buffer)
	defer func() { log.Out = oldOut }()

	//act -> When we print the message
	log.Print("testing")

	//assert -> Then we expect the message to be printed without any unix color codes
	if diff := cmp.Diff(output, log.Out.(*bytes.Buffer).String()); diff != "" {
		t.Fatalf("Unexpected result (-want +got):\n%s", diff)
	}
}

func TestPrintF_WHEN_SendingContentToPrintAndIsWindows_THEN_ShouldSendTheContentToTheOutputBuffer(t *testing.T) {
	//arrange -> Given that we try to print a message on windows with a format string
	oldIsWindows := log.IsWindows
	log.IsWindows = true
	defer func() { log.IsWindows = oldIsWindows }()
	output := "testing123"
	oldOut := log.Out
	log.Out = new(bytes.Buffer)
	defer func() { log.Out = oldOut }()

	//act -> When we print the message
	log.Printf("testing%v", 123)

	//assert -> Then we expect the message to be printed without any unix color codes
	if diff := cmp.Diff(output, log.Out.(*bytes.Buffer).String()); diff != "" {
		t.Fatalf("Unexpected result (-want +got):\n%s", diff)
	}
}

func TestError_WHEN_SendingContentToPrintAndIsWindows_THEN_ShouldSendTheContentToTheOutputBuffer(t *testing.T) {
	//arrange -> Given that we try to print an info message on windows
	oldIsWindows := log.IsWindows
	log.IsWindows = true
	defer func() { log.IsWindows = oldIsWindows }()
	output := "testing"
	oldOut := log.Out
	log.Out = new(bytes.Buffer)
	defer func() { log.Out = oldOut }()

	//act -> When we print the message
	log.Error("testing")

	//assert -> Then we expect the message to be printed without any unix color codes
	if diff := cmp.Diff(output, log.Out.(*bytes.Buffer).String()); diff != "" {
		t.Fatalf("Unexpected result (-want +got):\n%s", diff)
	}
}

func TestErrorF_WHEN_SendingContentToPrintAndIsWindows_THEN_ShouldSendTheContentToTheOutputBuffer(t *testing.T) {
	//arrange -> Given that we try to print an info message on windows with a format string
	oldIsWindows := log.IsWindows
	log.IsWindows = true
	defer func() { log.IsWindows = oldIsWindows }()
	output := "testing123"
	oldOut := log.Out
	log.Out = new(bytes.Buffer)
	defer func() { log.Out = oldOut }()

	//act -> When we print the message
	log.ErrorF("testing%v", 123)

	//assert -> Then we expect the message to be printed without any unix color codes
	if diff := cmp.Diff(output, log.Out.(*bytes.Buffer).String()); diff != "" {
		t.Fatalf("Unexpected result (-want +got):\n%s", diff)
	}
}

func TestWarning_WHEN_SendingContentToPrintAndIsWindows_THEN_ShouldSendTheContentToTheOutputBuffer(t *testing.T) {
	//arrange -> Given that we try to print a warning message
	oldIsWindows := log.IsWindows
	log.IsWindows = true
	defer func() { log.IsWindows = oldIsWindows }()
	output := "testing"
	oldOut := log.Out
	log.Out = new(bytes.Buffer)
	defer func() { log.Out = oldOut }()

	//act -> When we print the message
	log.Warning("testing")

	//assert -> Then we expect the message to be printed without any unix color codes
	if diff := cmp.Diff(output, log.Out.(*bytes.Buffer).String()); diff != "" {
		t.Fatalf("Unexpected result (-want +got):\n%s", diff)
	}
}

func TestWarningF_WHEN_SendingContentToPrintAndIsWindows_THEN_ShouldSendTheContentToTheOutputBuffer(t *testing.T) {
	//arrange -> Given that we try to print a warning message with a format string
	oldIsWindows := log.IsWindows
	log.IsWindows = true
	defer func() { log.IsWindows = oldIsWindows }()
	output := "testing123"
	oldOut := log.Out
	log.Out = new(bytes.Buffer)
	defer func() { log.Out = oldOut }()

	//act -> When we print the message
	log.Warningf("testing%v", 123)

	//assert -> Then we expect the message to be printed without any unix color codes
	if diff := cmp.Diff(output, log.Out.(*bytes.Buffer).String()); diff != "" {
		t.Fatalf("Unexpected result (-want +got):\n%s", diff)
	}
}

func TestSuccess_WHEN_SendingContentToPrintAndIsWindows_THEN_ShouldSendTheContentToTheOutputBuffer(t *testing.T) {
	//arrange -> Given that we try to print a success message on windows
	oldIsWindows := log.IsWindows
	log.IsWindows = true
	defer func() { log.IsWindows = oldIsWindows }()
	output := "testing"
	oldOut := log.Out
	log.Out = new(bytes.Buffer)
	defer func() { log.Out = oldOut }()

	//act -> When we print the message
	log.Success("testing")

	//assert -> Then we expect the message to be printed without any unix color codes
	if diff := cmp.Diff(output, log.Out.(*bytes.Buffer).String()); diff != "" {
		t.Fatalf("Unexpected result (-want +got):\n%s", diff)
	}
}

func TestSuccessF_WHEN_SendingContentToPrintAndIsWindows_THEN_ShouldSendTheContentToTheOutputBuffer(t *testing.T) {
	//arrange -> Given that we try to print a success message on windows with a format string
	oldIsWindows := log.IsWindows
	log.IsWindows = true
	defer func() { log.IsWindows = oldIsWindows }()
	output := "testing123"
	oldOut := log.Out
	log.Out = new(bytes.Buffer)
	defer func() { log.Out = oldOut }()

	//act -> When we print the message
	log.SuccessF("testing%v", 123)

	//assert -> Then we expect the message to be printed without any unix color codes
	if diff := cmp.Diff(output, log.Out.(*bytes.Buffer).String()); diff != "" {
		t.Fatalf("Unexpected result (-want +got):\n%s", diff)
	}
}

func TestFatal_WHEN_SendingContentToPrintAndIsWindows_THEN_ShouldSendTheContentToTheOutputBuffer(t *testing.T) {
	//arrange -> Given that we try to print a fatal message
	oldIsWindows := log.IsWindows
	log.IsWindows = true
	defer func() { log.IsWindows = oldIsWindows }()
	output := "testing"
	outCode := 1
	oldOut := log.Out
	log.Out = new(bytes.Buffer)
	defer func() { log.Out = oldOut }()
	code := 0
	oldExit := log.Exit
	log.Exit = func(c int) { code = c }
	defer func() { log.Exit = oldExit }()

	//act -> When we print the message
	log.Fatal("testing")

	//assert -> Then we expect the message to be printed without any unix color code and the program to exit with code 1
	if diff := cmp.Diff(output, log.Out.(*bytes.Buffer).String()); diff != "" {
		t.Fatalf("Unexpected result (-want +got):\n%s", diff)
	}
	if diff := cmp.Diff(outCode, code); diff != "" {
		t.Fatalf("Unexpected result (-want +got):\n%s", diff)
	}
}

func TestFatalF_WHEN_SendingContentToPrintAndIsWindows_THEN_ShouldSendTheContentToTheOutputBuffer(t *testing.T) {
	//arrange -> Given that we try to print a fatal message with a format string
	oldIsWindows := log.IsWindows
	log.IsWindows = true
	defer func() { log.IsWindows = oldIsWindows }()
	output := "testing123"
	outCode := 1
	oldOut := log.Out
	log.Out = new(bytes.Buffer)
	defer func() { log.Out = oldOut }()
	code := 0
	oldExit := log.Exit
	log.Exit = func(c int) { code = c }
	defer func() { log.Exit = oldExit }()

	//act -> When we print the message
	log.FatalF("testing%v", 123)

	//assert -> Then we expect the message to be printed without any unix color code and the program to exit with code 1
	if diff := cmp.Diff(output, log.Out.(*bytes.Buffer).String()); diff != "" {
		t.Fatalf("Unexpected result (-want +got):\n%s", diff)
	}
	if diff := cmp.Diff(outCode, code); diff != "" {
		t.Fatalf("Unexpected result (-want +got):\n%s", diff)
	}
}
