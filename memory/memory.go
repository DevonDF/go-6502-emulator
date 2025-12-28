package memory

type Memory struct {
	memory []byte
}

// NewMemory creates and returns a new Memory object with given size.
func NewMemory(size int) *Memory {
	return &Memory{
		memory: make([]byte, size),
	}
}
