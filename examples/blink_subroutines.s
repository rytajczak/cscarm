.global _start

.text
_start:
    movw    r4, #0x0000
    movt    r4, #0x3f20
    add     r2, r4, #0x08
    ldr     r3, [r2]
    orr     r3, r3, #0b001000
    str     r3, [r2]
    movw    r2, #0x0000
    movt    r2, #0x0020

loop:
    add     r3, r4, #0x1c
    str     r2, [r3]
    bl      delay                   @ branch forward 5 instructions

    add     r3, r4, #0x28
    str     r2, [r3]
    bl      delay                   @ branch forward 0 instructions

    b       loop                    @ branch back 12 instructions

delay:
    movw    r5, #0x4240
    movt    r5, #0x000f
delay_loop:
    subs    r5, r5, #1
    bge     delay_loop              @ branch back 3 instructions
    bx      lr
