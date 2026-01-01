; simple program to compute 2 to the power of 4

; init
LDA #&01 ; start accumulator with 1
STA $00 ; store 1 in 0x0000
LDX #&00 ; start X with 0

; main loop
loop:
    INX ; increment X register
    ADC $00 ; add value of 0x0000 to ac
    STA $00 ; store back in 0x0000
    CPX #$04 ; check if counter is 4
    BEQ finish
    JMP loop

finish:
    BRK