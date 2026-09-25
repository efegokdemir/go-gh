package term

import "io"

// Terminal describes the terminal properties used by extensions.
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

// TestTerminal is a configurable Terminal implementation for tests.
type TestTerminal struct {
	InReader         io.Reader
	OutWriter        io.Writer
	ErrOutWriter     io.Writer
	TerminalOutput   bool
	ColorEnabled     bool
	Color256Enabled  bool
	TrueColorEnabled bool
	TerminalWidth    int
	TerminalHeight   int
	SizeError        error
	TerminalTheme    string
}

var _ Terminal = Term{}
var _ Terminal = TestTerminal{}

func (t TestTerminal) In() io.Reader {
	return t.InReader
}

func (t TestTerminal) Out() io.Writer {
	return t.OutWriter
}

func (t TestTerminal) ErrOut() io.Writer {
	return t.ErrOutWriter
}

func (t TestTerminal) IsTerminalOutput() bool {
	return t.TerminalOutput
}

func (t TestTerminal) IsColorEnabled() bool {
	return t.ColorEnabled
}

func (t TestTerminal) Is256ColorSupported() bool {
	return t.Color256Enabled
}

func (t TestTerminal) IsTrueColorSupported() bool {
	return t.TrueColorEnabled
}

func (t TestTerminal) Size() (int, int, error) {
	return t.TerminalWidth, t.TerminalHeight, t.SizeError
}

func (t TestTerminal) Theme() string {
	return t.TerminalTheme
}
