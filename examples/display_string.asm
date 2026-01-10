; simple program to display a string

JMP main

main:
    LDX #$00
loop:
    LDA testMemory,X
    BEQ finish
    STA $4000,X
    INX
    JMP loop
finish:
    BRK

testMemory:
    .byte "Hello World!", 0x00