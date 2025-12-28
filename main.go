package main

import (
	"emulator/emulator"
	"flag"
	"fmt"
)

func main() {
	verbose := flag.Bool("v", false, "enable verbose logs")
	rom := flag.String("rom", "", "ROM file to load and execute in emulator")
	flag.Parse()

	if *rom == "" {
		flag.Usage()
		return
	}

	newEmulator := emulator.NewEmulator(
		emulator.EmulatorConfiguration{
			Verbose: *verbose,
		})

	err := newEmulator.LoadROM(*rom)
	if err != nil {
		fmt.Printf("failed to load ROM into memory: %w", err)
		return
	}

	newEmulator.StartEmulator()
}
