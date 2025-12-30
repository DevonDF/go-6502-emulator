LDA #&0B
.loop
ADC #&FF
TAX
STA &00,X
BNE .loop
BRK