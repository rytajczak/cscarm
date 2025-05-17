    MOVW    R4, #0x0000
    MOVT    R4, #0x3F20
    ADD     R2, R4, #0x08
    LDR     R3, [R2]
    ORR     R3, R3, #0b001000
    STR     R3, [R2]
    MOVW    R2, #0x0000
    MOVT    R2, #0x0020

blink_loop:
    ADD     R3, R4, #0x1C
    STR     R2, [R3]
    BL      delay

    ADD     R3, R4, #0x28
    STR     R2, [R3]
    BL      delay

    B       blink_loop

delay:
    MOVW    R5, #0xFFFF
    MOVT    R5, #0x000F
delay_loop:
    SUBS    R5, R5, #1
    BPL     delay_loop
    BX      LR