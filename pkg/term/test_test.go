package term

import (
	"errors"
	"io"
	"strings"
	"testing"
)

func TestTestTerminal(t *testing.T) {
	in := strings.NewReader("input")
	out := io.Discard
	errOut := io.Discard
	sizeErr := errors.New("size unavailable")

	terminal := TestTerminal{
		InReader:         in,
		OutWriter:        out,
		ErrOutWriter:     errOut,
		TerminalOutput:   true,
		ColorEnabled:     true,
		Color256Enabled:  true,
		TrueColorEnabled: true,
		TerminalWidth:    80,
		TerminalHeight:   24,
		SizeError:        sizeErr,
		TerminalTheme:    "dark",
	}

	if terminal.In() != in {
		t.Fatal("expected configured input reader")
	}
	if terminal.Out() != out {
		t.Fatal("expected configured output writer")
	}
	if terminal.ErrOut() != errOut {
		t.Fatal("expected configured error writer")
	}
	if !terminal.IsTerminalOutput() || !terminal.IsColorEnabled() ||
		!terminal.Is256ColorSupported() || !terminal.IsTrueColorSupported() {
		t.Fatal("expected configured terminal capabilities")
	}
	if width, height, err := terminal.Size(); width != 80 || height != 24 || !errors.Is(err, sizeErr) {
		t.Fatalf("unexpected terminal size: %d x %d, %v", width, height, err)
	}
	if terminal.Theme() != "dark" {
		t.Fatalf("expected dark theme, got %q", terminal.Theme())
	}
}
