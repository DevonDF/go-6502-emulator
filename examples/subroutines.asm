main:
    LDA #$01
    JSR routine
    STA $00
    BRK

routine:
    ADC #$01
    RTS