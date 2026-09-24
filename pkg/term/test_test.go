package term

import (
	"bytes"
	"strings"
	"testing"
)

func TestTestTerminal(t *testing.T) {
	input := strings.NewReader("input")
	output := &bytes.Buffer{}
	errOutput := &bytes.Buffer{}
	terminal := TestTerminal{
		InReader:           input,
		OutWriter:          output,
		ErrOutWriter:       errOutput,
		IsTTY:              true,
		ColorEnabled:       true,
		Is256ColorEnabled:  true,
		IsTrueColorEnabled: true,
		Width:              120,
		Height:             40,
		ThemeValue:         "dark",
	}

	if terminal.In() != input {
		t.Fatal("expected configured input reader")
	}
	if terminal.Out() != output {
		t.Fatal("expected configured output writer")
	}
	if terminal.ErrOut() != errOutput {
		t.Fatal("expected configured error writer")
	}
	if !terminal.IsTerminalOutput() || !terminal.IsColorEnabled() {
		t.Fatal("expected configured terminal and color capabilities")
	}
	if !terminal.Is256ColorSupported() || !terminal.IsTrueColorSupported() {
		t.Fatal("expected configured color depth capabilities")
	}
	if width, height, err := terminal.Size(); err != nil || width != 120 || height != 40 {
		t.Fatalf("expected configured size, got %dx%d (%v)", width, height, err)
	}
	if theme := terminal.Theme(); theme != "dark" {
		t.Fatalf("expected dark theme, got %q", theme)
	}
}

func TestTestTerminalZeroValue(t *testing.T) {
	terminal := TestTerminal{}
	if terminal.IsTerminalOutput() || terminal.IsColorEnabled() {
		t.Fatal("expected zero-value terminal capabilities to be disabled")
	}
	if width, height, err := terminal.Size(); err == nil || width != -1 || height != -1 {
		t.Fatalf("expected unconfigured size error, got %dx%d (%v)", width, height, err)
	}
	if theme := terminal.Theme(); theme != "none" {
		t.Fatalf("expected no theme, got %q", theme)
	}
}
