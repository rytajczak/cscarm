    MOVW    R4, #0x0000
    MOVT    R4, #0x3f20
    ADD     R2, R4, #0x08
    LDR     R3, [R2]
    ORR     R3, R3, #0b001000
    STR     R3, [R2]

blink_loop:
    ADD     R3, R4, #0x1c           @ blink_loop
    MOVW    R2, #0x0000
    MOVT    R2, #0x0020
    STR     R2, [R3]
    BL      delay                   @ branch forward 5 instructions

    ADD     R3, R4, #0x28
    MOVW    R2, #0x0000
    MOVT    R2, #0x0020
    STR     R2, [R3]
    BL      delay                   @ branch forward 0 instructions

    B       blink_loop              @ branch back 12 instructions

delay:
    MOVW    R5, #0xFFFF             @ delay
    MOVT    R5, #0x000F
delay_loop:
    SUBS    R5, R5, #1              @ delay_loop
    BPL     delay_loop              @ branch back 3 instructions
    BX      LR