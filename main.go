package main

import (
	"emulator/emulator"
	"flag"
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

	newEmulator.LoadAndExecute(*rom)
}
