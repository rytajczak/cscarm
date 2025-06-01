	movw 	sp, #0x0000             @ initialize stack pointer to 0
	movt 	sp, #0x0000	
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
	stmea 	sp!, {r0-r12}           @ push registers 0 - 12 onto the stack
	movw 	r12, #0x0900            @ move 4 million into r12
	movt 	r12, #0x3d
	str 	r12, [sp], #4           @ push the value at r12 onto the stack
    bl      delay
	ldmea 	sp!, {r0-r12}           @ restore registers 0 - 12 from the stack

    add     r3, r4, #0x28
    str     r2, [r3]
	stmea 	sp!, {r0-r12}
	movw 	r12, #0x4240            @ move 1 million into r12 this time
	movt 	r12, #0xf
	str 	r12, [sp], #4
    bl      delay
	ldmea 	sp!, {r0-r12}

    b       loop

delay:
	ldr 	r5, [sp, #-4]!          @ load the top value from stack into r5
delay_loop:
    subs    r5, r5, #1
    bge     delay_loop
	bx      lr

