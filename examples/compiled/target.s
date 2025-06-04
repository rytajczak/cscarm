	movw    sp, #0x0000
	movt    sp, #0x0000

_start:
    movw    r4, #0x0000
    movt    r4, #0x3f20
    add     r2, r4, #0x08
    ldr     r3, [r2]
    orr     r3, r3, #0b001000
    str     r3, [r2]
    movw    r2, #0x0000
    movt    r2, #0x0020
    movw    r6, #7
    str     r6, [sp], #4

loop:
    bl      blink
    movw    r5, #0x4240
    movt    r5, #0xf
    str     r5, [sp], #4
    bl      delay
    subs    r6, r6, #1
    bne     loop

    movw    r5, #0x0900
    movt    r5, #0x3d
    str     r5, [sp], #4
    bl      delay

    ldr     r6, [sp, #-4]
    b       loop

blink:
    str     lr, [sp], #4
    add     r3, r4, #0x1c
    str     r2, [r3]
    movw    r5, #0x4240
    movt    r5, #0xf
    str     r5, [sp], #4
    bl      delay
    add     r3, r4, #0x28
    str     r2, [r3]
    ldr     lr, [sp, #-4]!
    bx      lr

delay:
    ldr     r5, [sp, #-4]!
delay_loop:
    subs    r5, r5, #1
    bge     delay_loop
    bx      lr
