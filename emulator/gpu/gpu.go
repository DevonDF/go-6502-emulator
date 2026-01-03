package gpu

import (
	"fmt"
	"log/slog"

	"github.com/DevonDF/go-6502-emulator/emulator/memory"
)

const (
	screenWidth  = 40 // 40 bytes/pixels width
	screenHeight = 25 // 25 bytes/pixel height

	screenRAMStart  = 0x4000 // where in memory the screen ram starts, that is the characters to display
	screenRAMLength = 1000   // how long the screen ram is in bytes

	colourRAMStart  = 0xD800 // where in memory the screen colour RAM starts, that is the fg colour
	colourRAMLength = 1000   // how long the colour ram is in bytes

	backgroundColourMemoryLocation = 0xD021 // the memory location of the background colour to use
)

type GPU struct {
	memory *memory.Memory // the memory to read from for interaction with the gpu.
	logger *slog.Logger
}

// NewGPU creates and returns a new GPU attached to the given memory unit.
func NewGPU(memory *memory.Memory, logger *slog.Logger) *GPU {
	return &GPU{
		memory: memory,
		logger: logger,
	}
}

// Display displays the contents of the video RAM onto the screen.
func (gpu *GPU) Display() {
	toPrint := ""
	toPrint += "\033[2J\033[H" // reset the terminal

	backgroundColourCode, err := gpu.memory.ReadByte(backgroundColourMemoryLocation)
	if err != nil {
		gpu.logger.Error("failed to read the background colour RAM", "error", err)
		return
	}

	if backgroundColourCode != 0x0 {
		backgroundAnsiCode, found := PetsciiToAnsiBackground[backgroundColourCode]
		if !found {
			gpu.logger.Error("undefined colour set as background colour", "colourCode", backgroundColourCode)
			return
		}

		toPrint += backgroundAnsiCode // set the background colour
	}

	var addrOffset uint16
	for addrOffset = 0; addrOffset < screenRAMLength; addrOffset++ {

		// find any foreground colour code from colour ram
		colourCode, err := gpu.memory.ReadByte(colourRAMStart + addrOffset)
		if err != nil {
			gpu.logger.Error("failed to read the colour code RAM", "error", err)
			return
		}

		if colourCode != 0x0 {
			colourAnsiCode, found := PetsciiToAnsiColoursForeground[colourCode]
			if !found {
				gpu.logger.Error("undefined colour set as foreground colour", "colourCode", colourCode)
				return
			}

			toPrint += colourAnsiCode // set the foreground colour
		}

		// get the screen character from memory and find matching petscii
		characterCode, err := gpu.memory.ReadByte(screenRAMStart + addrOffset)
		if err != nil {
			gpu.logger.Error("failed to read the character code RAM", "error", err)
			return
		}

		petsciiChar, found := PetsciiToASCII[characterCode]
		if !found {
			petsciiChar = " "
		}

		toPrint += " " + petsciiChar + " "

		if addrOffset > 0 && (addrOffset+1)%screenWidth == 0 {
			toPrint += "\n"
		}

	}

	toPrint += AnsiReset // reset ANSI state

	fmt.Print(toPrint)

}
