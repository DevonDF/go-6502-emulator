package cpu

type CPU struct {
	rPC uint16
	rAC uint8
	rX  uint8
	rY  uint8
	rSR uint8
	rSP uint8
}

// NewCPU creates a new CPU struct and returns it.
func NewCPU() *CPU {
	return &CPU{
		rPC: 0,
		rAC: 0,
		rX:  0,
		rY:  0,
		rSR: 0,
		rSP: 0,
	}
}
