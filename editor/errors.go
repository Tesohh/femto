package editor

import (
	"fmt"
	"log/slog"
)

type FemtoError struct {
	Message  string
	LogLevel slog.Level
}

func (f FemtoError) Error() string {
	return f.Message
}

func (f FemtoError) Context(msg string) FemtoError {
	return FemtoError{
		Message:  fmt.Sprintf("%s (%s)", f.Error(), msg),
		LogLevel: f.LogLevel,
	}
}

var (
	ErrNoKeyAssociated = FemtoError{
		Message:  "no key associated with pressed key",
		LogLevel: slog.LevelWarn,
	}
	ErrNoCommandFound = FemtoError{
		Message:  "no command found",
		LogLevel: slog.LevelError,
	}
	ErrKeyUnparsable = FemtoError{
		Message:  "Cannot parse key",
		LogLevel: slog.LevelError,
	}
)
