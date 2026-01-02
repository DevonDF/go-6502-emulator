package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/DevonDF/go-6502-emulator/emulator"
)

func main() {
	rom := flag.String("rom", "", "ROM file to load and execute in emulator")
	logPath := flag.String("logPath", "emu.log", "path to the log file for this emulator")
	debug := flag.Bool("debug", false, "use the debugger")
	flag.Parse()

	if *rom == "" {
		flag.Usage()
		return
	}

	logFile, err := os.OpenFile(*logPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("failed to open log file %s: %v", *logPath, err)
		return
	}
	defer logFile.Close()

	newEmulator := emulator.NewEmulator(
		emulator.EmulatorConfiguration{
			Debug:   *debug,
			Logfile: *logFile,
		})

	err = newEmulator.LoadROM(*rom)
	if err != nil {
		fmt.Printf("failed to load ROM into memory: %v", err)
		return
	}
	newEmulator.StartEmulator()
}
