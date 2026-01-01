main:
    LDA #$09
    JMP test
    BRK
    BRK
    BRK
    BRK

test:
    ADC #$01
    STA $01