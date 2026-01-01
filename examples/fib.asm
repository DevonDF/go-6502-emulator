; simple quick fibonacci program

define last $00 ; variable last
define secondLast $01 ; variable secondLast

define index #$0A ; calculate the 10th fib number

LDA #$01 ; start the fib sequence with 1
STA last ; start the fib sequence with 1

LDY #$01 ; use Y as a counter

loop:
    LDA last ; load last into accumulator
    TAX ; save last into register X
    ADC secondLast ; add the value of secondLast to accumulator
    STA last ; store the current fib number into last
    TXA ; restore the previous last value
    STA secondLast ; store the previous last value into secondLast
    INY ; increment Y
    CPY index ; check if we have calculated enough fib numbers
    BEQ finish
    JMP loop

finish:
    STA $00 ; store the result at $0000
    BRK
