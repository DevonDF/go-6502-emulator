package memory

import "log/slog"

type Memory struct {
	memory []byte
	logger *slog.Logger
}

// NewMemory creates and returns a new Memory object.
func NewMemory(logger *slog.Logger) *Memory {
	return &Memory{
		memory: make([]byte, 65536),
		logger: logger,
	}
}
