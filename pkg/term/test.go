package term

import (
	"errors"
	"io"
)

// Terminal describes the terminal capabilities exposed by Term.
//
// It is useful for code that accepts terminal dependencies to use this
// interface so callers can provide TestTerminal in unit tests.
type Terminal interface {
	In() io.Reader
	Out() io.Writer
	ErrOut() io.Writer
	IsTerminalOutput() bool
	IsColorEnabled() bool
	Is256ColorSupported() bool
	IsTrueColorSupported() bool
	Size() (int, int, error)
	Theme() string
}

// TestTerminal is a configurable terminal implementation for tests.
//
// Its zero value is deterministic: streams are nil, terminal capabilities are
// disabled, and Size reports that no size was configured.
type TestTerminal struct {
	InReader           io.Reader
	OutWriter          io.Writer
	ErrOutWriter       io.Writer
	IsTTY              bool
	ColorEnabled       bool
	Is256ColorEnabled  bool
	IsTrueColorEnabled bool
	Width              int
	Height             int
	ThemeValue         string
}

var _ Terminal = Term{}
var _ Terminal = TestTerminal{}

// In is the reader configured for standard input.
func (t TestTerminal) In() io.Reader {
	return t.InReader
}

// Out is the writer configured for standard output.
func (t TestTerminal) Out() io.Writer {
	return t.OutWriter
}

// ErrOut is the writer configured for standard error.
func (t TestTerminal) ErrOut() io.Writer {
	return t.ErrOutWriter
}

// IsTerminalOutput reports whether output is configured as a terminal.
func (t TestTerminal) IsTerminalOutput() bool {
	return t.IsTTY
}

// IsColorEnabled reports whether ANSI color is enabled.
func (t TestTerminal) IsColorEnabled() bool {
	return t.ColorEnabled
}

// Is256ColorSupported reports whether ANSI 256 color is supported.
func (t TestTerminal) Is256ColorSupported() bool {
	return t.Is256ColorEnabled
}

// IsTrueColorSupported reports whether ANSI true color is supported.
func (t TestTerminal) IsTrueColorSupported() bool {
	return t.IsTrueColorEnabled
}

// Size returns the configured terminal dimensions.
func (t TestTerminal) Size() (int, int, error) {
	if t.Width <= 0 || t.Height <= 0 {
		return -1, -1, errors.New("terminal size is not configured")
	}
	return t.Width, t.Height, nil
}

// Theme returns the configured terminal theme, or "none" when no theme is set.
func (t TestTerminal) Theme() string {
	if t.ThemeValue == "" {
		return "none"
	}
	return t.ThemeValue
}
