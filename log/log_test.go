package log_test

import (
	"bytes"
	"journal/log"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPrint(t *testing.T) {
	oldIsWindows := log.IsWindows
	log.IsWindows = false
	defer func() { log.IsWindows = oldIsWindows }()
	output := "\033[97mtesting\033[0m"
	oldOut := log.Out
	log.Out = new(bytes.Buffer)
	defer func() { log.Out = oldOut }()
	log.Print("testing")
	if diff := cmp.Diff(output, log.Out.(*bytes.Buffer).String()); diff != "" {
		t.Fatalf("Unexpected result (-want +got):\n%s", diff)
	}
}

func TestPrintf(t *testing.T) {
	oldIsWindows := log.IsWindows
	log.IsWindows = false
	defer func() { log.IsWindows = oldIsWindows }()
	output := "\033[97mtesting123\033[0m"
	oldOut := log.Out
	log.Out = new(bytes.Buffer)
	defer func() { log.Out = oldOut }()
	log.Printf("testing%v", 123)
	if diff := cmp.Diff(output, log.Out.(*bytes.Buffer).String()); diff != "" {
		t.Fatalf("Unexpected result (-want +got):\n%s", diff)
	}
}

func TestError(t *testing.T) {
	oldIsWindows := log.IsWindows
	log.IsWindows = false
	defer func() { log.IsWindows = oldIsWindows }()
	output := "\033[31mtesting\033[0m"
	oldOut := log.Out
	log.Out = new(bytes.Buffer)
	defer func() { log.Out = oldOut }()
	log.Error("testing")
	if diff := cmp.Diff(output, log.Out.(*bytes.Buffer).String()); diff != "" {
		t.Fatalf("Unexpected result (-want +got):\n%s", diff)
	}
}

func TestErrorF(t *testing.T) {
	oldIsWindows := log.IsWindows
	log.IsWindows = false
	defer func() { log.IsWindows = oldIsWindows }()
	output := "\033[31mtesting123\033[0m"
	oldOut := log.Out
	log.Out = new(bytes.Buffer)
	defer func() { log.Out = oldOut }()
	log.ErrorF("testing%v", 123)
	if diff := cmp.Diff(output, log.Out.(*bytes.Buffer).String()); diff != "" {
		t.Fatalf("Unexpected result (-want +got):\n%s", diff)
	}
}

func TestWarning(t *testing.T) {
	oldIsWindows := log.IsWindows
	log.IsWindows = false
	defer func() { log.IsWindows = oldIsWindows }()
	output := "\033[33mtesting\033[0m"
	oldOut := log.Out
	log.Out = new(bytes.Buffer)
	defer func() { log.Out = oldOut }()
	log.Warning("testing")
	if diff := cmp.Diff(output, log.Out.(*bytes.Buffer).String()); diff != "" {
		t.Fatalf("Unexpected result (-want +got):\n%s", diff)
	}
}

func TestWarningF(t *testing.T) {
	oldIsWindows := log.IsWindows
	log.IsWindows = false
	defer func() { log.IsWindows = oldIsWindows }()
	output := "\033[33mtesting123\033[0m"
	oldOut := log.Out
	log.Out = new(bytes.Buffer)
	defer func() { log.Out = oldOut }()
	log.Warningf("testing%v", 123)
	if diff := cmp.Diff(output, log.Out.(*bytes.Buffer).String()); diff != "" {
		t.Fatalf("Unexpected result (-want +got):\n%s", diff)
	}
}

func TestSuccess(t *testing.T) {
	oldIsWindows := log.IsWindows
	log.IsWindows = false
	defer func() { log.IsWindows = oldIsWindows }()
	output := "\033[32mtesting\033[0m"
	oldOut := log.Out
	log.Out = new(bytes.Buffer)
	defer func() { log.Out = oldOut }()
	log.Success("testing")
	if diff := cmp.Diff(output, log.Out.(*bytes.Buffer).String()); diff != "" {
		t.Fatalf("Unexpected result (-want +got):\n%s", diff)
	}
}

func TestSuccessF(t *testing.T) {
	oldIsWindows := log.IsWindows
	log.IsWindows = false
	defer func() { log.IsWindows = oldIsWindows }()
	output := "\033[32mtesting123\033[0m"
	oldOut := log.Out
	log.Out = new(bytes.Buffer)
	defer func() { log.Out = oldOut }()
	log.SuccessF("testing%v", 123)
	if diff := cmp.Diff(output, log.Out.(*bytes.Buffer).String()); diff != "" {
		t.Fatalf("Unexpected result (-want +got):\n%s", diff)
	}
}

func TestFatal(t *testing.T) {
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
	log.Fatal("testing")
	if diff := cmp.Diff(output, log.Out.(*bytes.Buffer).String()); diff != "" {
		t.Fatalf("Unexpected result (-want +got):\n%s", diff)
	}
	if diff := cmp.Diff(outCode, code); diff != "" {
		t.Fatalf("Unexpected result (-want +got):\n%s", diff)
	}
}

func TestFatalF(t *testing.T) {
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
	log.FatalF("testing%v", 123)
	if diff := cmp.Diff(output, log.Out.(*bytes.Buffer).String()); diff != "" {
		t.Fatalf("Unexpected result (-want +got):\n%s", diff)
	}
	if diff := cmp.Diff(outCode, code); diff != "" {
		t.Fatalf("Unexpected result (-want +got):\n%s", diff)
	}
}

func TestPrintWindows(t *testing.T) {
	oldIsWindows := log.IsWindows
	log.IsWindows = true
	defer func() { log.IsWindows = oldIsWindows }()
	output := "\x1b[97mtesting"
	oldOut := log.Out
	log.Out = new(bytes.Buffer)
	defer func() { log.Out = oldOut }()
	log.Print("testing")
	if diff := cmp.Diff(output, log.Out.(*bytes.Buffer).String()); diff != "" {
		t.Fatalf("Unexpected result (-want +got):\n%s", diff)
	}
}
