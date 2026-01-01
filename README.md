# 6502 Emulator

An **unfinished** emulator and accompanying assembler for executing binaries written for the 6502 chipset. 

All information & documentation regarding the instruction-set can be found here: https://www.masswerk.at/6502/6502_instruction_set.html

## Building

The assembler can be built and executed using Golang as normal

```shell
go build ./cmd/assembler
```

The emulator can be built and executed using Golang as normal

```shell
go build ./cmd/emulator
```

## Usage

```
Usage of assembler:
  -input string
        assembly file input
  -loadAddr string
        hex address of where the ROM will be loaded (default "0x8000")
  -output string
        compiled bytecode output
```

The assembler expects an assembly file of assembly instructions, the address of where this binary will be loaded into memory (for jumps & subroutines within the 6502 architecture), and an output path for the assembled bytecode.

```
Usage of emulator:
  -debug
        enable debug mode
  -rom string
        ROM file to load and execute in emulator
```

The emulator expects you to pass in a ROM/IMG file that will be mapped directly into the ROM pages within its emulated memory ($0x8000 onwards). The PC will start at $0x8000. Therefore, binaries should begin with instructions, and have any data appended at the bottom.

## Example Usage

Assemble an assembly program `subroutines.asm` and run it on the emulator in debug mode:
```shell
go run ./cmd/assembler -input ./examples/subroutines.asm -output ./examples/subroutines.bin -loadAddr 0x8000
go run ./cmd/emulator -rom ./examples/subroutines.bin -debug
```
