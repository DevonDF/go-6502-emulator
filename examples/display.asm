; a quick program to display 'EMULATOR' in the middle of the screen
; throughout this program, functions are called where the X register points to the arguments
; and the Y register points to where to store result
; e.g. LDX #$00, LDY #$01, JSR add8
; this calls the func add8 which reads the arguments from $00 and stores the result in $01

; definitions
define char_a #$41
define char_b #$42
define char_c #$43
define char_d #$44
define char_e #$45
define char_f #$46
define char_g #$47
define char_h #$48
define char_i #$49
define char_j #$4A
define char_k #$4B
define char_l #$4C
define char_m #$4D
define char_n #$4E
define char_o #$4F
define char_p #$50
define char_q #$51
define char_r #$52
define char_s #$53
define char_t #$54
define char_u #$55
define char_v #$56
define char_w #$57
define char_x #$58
define char_y #$59
define char_z #$5A

define capitalAddition #$20 ; what to add to capitalise a character

define char_0 #$30
define char_1 #$31
define char_2 #$32
define char_3 #$33
define char_4 #$34
define char_5 #$35
define char_6 #$36
define char_7 #$37
define char_8 #$38
define char_9 #$39

define black #$90
define red #$1C
define green #$1E
define yellow #$9E
define blue #$1F
define purple #$9C
define cyan #$9F
define white #$05

define screenRAMAddrH #$40
define screenRAMAddrL #$00
define colourRAMAddrH #$D8
define colourRAMAddrL #$00

define backgroundColourAddr $D021

; main program
JMP main

; swap swaps X & Y register
swap:
    PHA ; push accumulator to stack to avoid disrupting accumulator
    PHP ; push SR to stack to avoid disrupting SR
    STX $FF ; store X in $FF temporarily
    TYA ; transfer Y to A
    TAX ; transfer A to X
    LDY $FF ; transfer $FF (cached X) into Y
    PLP ; pull SR from stack to avoid disrupting SR
    PLA ; pull accumulator from stack to avoid disrupting accumulator
    RTS


; add16 adds two uint16 values
add16:
    CLC ; clear carry bit
    LDA $00,X ; load val1_low
    INX ; increment X by one to get val1_high
    INX ; increment X by one to get val2_low
    ADC $00,X ; add val2_low
    DEX ; decrement X by one to get val1_high
    JSR swap ; swap X & Y to get the result pointer
    STA $00,X ; store in result_low
    INX ; increment X by one to get result_high
    JSR swap ; swap X & Y to get the val1_high pointer back
    LDA $00,X ; load val1_high
    INX ; increment X by one to get val2_low
    INX ; increment X by one to get val2_high
    ADC $00,X ; add val2_high
    JSR swap ; swap X & Y to get the result pointer
    STA $00,X ; store in result_high
    RTS ; return


main:
    ; create uint16 variable to store pointer to screenRAM index
    LDA screenRAMAddrL
    STA $10
    LDA screenRAMAddrH
    STA $11 ; $10 now holds a uint16 pointer to screenRAM

    ; create uint16 variable to store pointer to colourRAM index
    LDA colourRAMAddrL
    STA $14
    LDA colourRAMAddrH
    STA $15 ; $14 now holds a uint16 pointer to colourRAM

    ; set background display to black
    LDA blue
    STA backgroundColourAddr

    ; add 0x01e0 to the screenRAM pointer & colourRAM pointer
    ; 0x01f0 is 496 - the middle of the screen
    ; when we call add16, we need a pointer to the two arguments in X, and a pointer to return value in Y
    LDA #$f0 ; low
    STA $12
    LDA #$01 ; high
    STA $13 ; $13 now holds the value 496
    LDX #$10 ; $10 holds the arguments for add16, two int16s
    LDY #$18 ; store result of addition in $18
    JSR add16 ; add the two int16s given at $X and place them into $Y
    ; $18 now holds ptr to screenRAM + 496

    LDA #$f0 ; low
    STA $16
    LDA #$01 ; high
    STA $17 ; $16 now holds the value 496
    LDX #$14 ; $14 holds the arguments for add16, two int16s
    LDY #$1A ; store result of addition in $1A
    JSR add16 ; add the two int16s given at $X and place them into $Y
    ; $1A now holds ptr to colourRAM + 496

    LDY #$00 ; offset y
    LDA white
    STA ($1A),Y ; set text colour to white
    LDA char_e ; load e into ac
    ADC capitalAddition ; capitalise
    STA ($18),Y ; write to screen

    INY ; increment Y to write to the next byte in the screenRAM & colourRAM
    LDA white
    STA ($1A),Y ; set text colour to white
    LDA char_m ; load m into ac
    ADC capitalAddition ; capitalise
    STA ($18),Y ; write to screen

    INY ; increment Y to write to the next byte in the screenRAM & colourRAM
    LDA white
    STA ($1A),Y ; set text colour to white
    LDA char_u ; load u into ac
    ADC capitalAddition ; capitalise
    STA ($18),Y ; write to screen

    INY ; increment Y to write to the next byte in the screenRAM & colourRAM
    LDA white
    STA ($1A),Y ; set text colour to white
    LDA char_l ; load l into ac
    ADC capitalAddition ; capitalise
    STA ($18),Y ; write to screen

    INY ; increment Y to write to the next byte in the screenRAM & colourRAM
    LDA white
    STA ($1A),Y ; set text colour to white
    LDA char_a ; load a into ac
    ADC capitalAddition ; capitalise
    STA ($18),Y ; write to screen

    INY ; increment Y to write to the next byte in the screenRAM & colourRAM
    LDA white
    STA ($1A),Y ; set text colour to white
    LDA char_t ; load t into ac
    ADC capitalAddition ; capitalise
    STA ($18),Y ; write to screen

    INY ; increment Y to write to the next byte in the screenRAM & colourRAM
    LDA white
    STA ($1A),Y ; set text colour to white
    LDA char_o ; load o into ac
    ADC capitalAddition ; capitalise
    STA ($18),Y ; write to screen

    INY ; increment Y to write to the next byte in the screenRAM & colourRAM
    LDA white
    STA ($1A),Y ; set text colour to white
    LDA char_r ; load r into ac
    ADC capitalAddition ; capitalise
    STA ($18),Y ; write to screen

    BRK