	movw    sp, #0
	movt    sp, #0
    movw    r4, #0x0000
    movt    r4, #0x3f20
    add     r2, r4, #0x08
    ldr     r3, [r2]
    orr     r3, r3, #0b001000
    str     r3, [r2]

blink_loop:
    add     r3, r4, #0x1c           @ blink loop
    movw    r2, #0x0000
    movt    r2, #0x0020
    str     r2, [r3]
    bl      delay                   @ 5

    add     r3, r4, #0x28
    movw    r2, #0x0000
    movt    r2, #0x0020
    str     r2, [r3]
    bl      delay                   @ 0

    b       blink_loop              @ -12

delay:
    movw    r5, #0xFFFF             @ delay
    movt    r5, #0x000F
delay_loop:
    subs    r5, r5, #1              @ delay_loop
    bpl     delay_loop              @ -3
    bx      lr
