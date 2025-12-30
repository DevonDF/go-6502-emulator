LDA #&01
AND #&01
BEQ .success
.fail
ADC #&02
STA &00
BRK
.success
ADC #&01
STA &00
BRK