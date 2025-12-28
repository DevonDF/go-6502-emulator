package main

import (
	"emulator/emulator"
	"flag"
	"fmt"
)

func main() {
	debug := flag.Bool("debug", false, "enable debug mode")
	rom := flag.String("rom", "", "ROM file to load and execute in emulator")
	flag.Parse()

	if *rom == "" {
		flag.Usage()
		return
	}

	newEmulator := emulator.NewEmulator(
		emulator.EmulatorConfiguration{
			Debug: *debug,
		})

	err := newEmulator.LoadROM(*rom)
	if err != nil {
		fmt.Printf("failed to load ROM into memory: %w", err)
		return
	}

	newEmulator.StartEmulator()
}
