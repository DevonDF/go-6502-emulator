; simple code to demonstrate functions within 6502
; functions typically pass arguments via the X/Y registers and zeropage

; init a list of numbers to sum
LDA #$02
STA $0200
LDA #$0A
STA $0201
LDA #$04
STA $0202

; construct arguments to function sum
LDA #$00
STA $00
LDA #$02
STA $01 ; first argument is uint16 address to the location of array
LDA #$03
STA $02 ; second argument is uint8 number of elements in array
JSR sum ; sum()

STA $00 ; store result in 0x0000
BRK ; end

; sum(address uint16, length uint8) int8
sum:
    LDA #$00 ; null accumulator
    LDX $02 ; load X with the length of the array
    LDY $00 ; Y will be our index of the array
sumLoop:
    ADC ($00),Y ; add the current array value to the accumulator
    INY ; increment our index Y
    DEX ; decrease our counter X
    CPX #$00 ; when we have read all of the array, we need to return
    BNE sumLoop
    RTS

