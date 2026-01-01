main:
    LDA #$01 ; load one into the accumulator
    JSR routine ; jump to the subroutine
    STA $00 ; store the value at $0000
    BRK

routine:
    ADC #$01 ; add one to the accumulator
    RTS ; return