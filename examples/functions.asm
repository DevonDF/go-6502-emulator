; simple code to demonstrate functions within 6502
; functions typically pass arguments via the X/Y registers and zeropage

; init a list of numbers to sum
LDA #$02
STA $80
LDA #$0A
STA $81
LDA #$04
STA $82

; construct arguments to function sum
LDX #$80 ; first argument is the start of the array
LDY #$03 ; second argument is the size of the array
JSR sum ; sum(0x0080, 3)
STA $00
BRK

; sum sums the list (of length Y) of integers given at zeropage address X 
sum:
    LDA #$00 ; null accumulator
sumLoop:
    ADC $00,X
    INX
    DEY
    CPY #$00
    BNE sumLoop
    RTS

