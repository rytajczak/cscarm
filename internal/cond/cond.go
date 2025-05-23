package cond

const (
	EQ uint32 = iota // Equal | Z == 1
	NE               // Not equal | Z == 0
	CS               // Carry set | C == 1
	CC               // Carry clear | C == 0
	MI               // Minus, negative | N == 1
	PL               // Plus, positive or zero | N == 0
	VS               // Overflow | V == 1
	VC               // No overflow | V == 0
	HI               // Unsigned higher | C == 1 and Z == 0
	LS               // Unsigned lower or same | C == 0 or Z == 1
	GE               // Signed Greater than or equal | N == V
	LT               // Signed less than | N != V
	GT               // Signed greater than | Z == 0 and N == V
	LE               // Signed less than or equal | Z == 1 or N != V
	AL               // Always (unconditional) | Any
)

var CondCodeMap = map[string]uint32{
	"EQ": EQ,
	"NE": NE,
	"CS": CS,
	"CC": CC,
	"MI": MI,
	"PL": PL,
	"VS": VS,
	"VC": VC,
	"HI": HI,
	"LS": LS,
	"GE": GE,
	"LT": LT,
	"GT": GT,
	"LE": LE,
	"AL": AL,
}
