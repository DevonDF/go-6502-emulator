# 6502 Emulator

An **unfinished** emulator for executing binaries written for the 6502 chipset. All information & documentation regarding the instruction-set can be found here: https://www.masswerk.at/6502/6502_instruction_set.html

## Building

The emulator can be built and executed using Golang as normal

```shell
go build -o 6502-emulator
```

## Usage

```
Usage of 6502-emulator.exe:
  -debug
        enable debug mode
  -rom string
        ROM file to load and execute in emulator
```

The emulator expects you to pass in a ROM/IMG file that will be mapped directly into the ROM pages within its emulated memory ($0x8000 onwards). The PC will start at $0x8000. Therefore, binaries should begin with instructions, and have any data appended at the bottom.


## Example Usage

With building:
```shell
./6502-emulator -rom program.bin -debug
```

Without building:
```shell
go run . -rom program.bin -debug
```