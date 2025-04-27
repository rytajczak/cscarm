	MOVW   R4, 0              ;set pin to output
	MOVT   R4, 0x3F20
	ADD    R2, R4, 0x08
	LDR    R3, (R2)
	ORR    R3, R3, 0x00000008
	STR    R3, (R2)

	ADD    R3, R4, 0x1C       ;GPIO pin output set 0
	MOVW   R2, 0x0000
	MOVT   R2, 0x0020
	STR    R2, (R3)
	MOVW   R5, 0xFFFF
	MOVT   R5, 0x000F
	SUBS   R5, R5, 0x01
	BPL    -3 0xFFFFFD

	ADD    R3, R4, 0x28 		  ;GPIO pin output clear 0 
	MOVW   R2, 0x0000
	MOVT   R2, 0x0020
	STR    R2, (R3)
	MOVW   R5, 0xFFFF
	MOVT   R5, 0x000F
	SUBS   R5, R5, 0x01
	BPL    -3 0xFFFFFD

	B      -18 0xFFFFEE