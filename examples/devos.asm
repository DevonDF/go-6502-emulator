; DevOS is my terrible Operating System for the 6502 architecture with Commodore 64 display-style graphics

; define OS function space within the zeropage
; $00-$0F are reserved for function calling & returns
define ZP_ARG0 $00 ; arg0 is 2 byte block for arg0 for functions
define ZP_ARG1 $02 ; arg1 is 2 byte block for arg1 for functions
define ZP_ARG2 $04 ; arg2 is 2 byte block for arg2 for functions
define ZP_ARG3 $06 ; arg3 is 2 byte block for arg3 for functions
define ZP_RET $08 ; ret is a 2 byte block for returns for functions
define ZP_TMP $0A ; tmp is a 6 byte block for temporary usage within functions

; bootstrap
JMP main

; utility functions

; putString prints a string held at ZP_ARG0
putString:
    LDY #$00 ; Y is our index variable for the string
    _putStringLoop:
        LDA (ZP_ARG0),Y ; load into A the Y'th element of the string given in ZP_ARG0
        BEQ _putStringRet ; if it's 0x00 i.e. null byte, lets return, we're done here
        STA $4000,Y ; print it to the screen at 0,Y 
        INY ; increment Y
        JMP _putStringLoop ; jump back
    _putStringRet:
        RTS


; main
main:
    ; start by printing DevOS
    LDA #<strDevOS ; low byte of strDevOS
    STA ZP_ARG0

    LDA #>strDevOS ; high byte of strDevOS
    STA ZP_ARG0+1

    JSR putString ; print
    BRK


strDevOS:
    .byte "DevOS", $00