package compiler

type PinModeType int

const (
	INPUT PinModeType = iota
	OUTPUT
)

type PinWriteType int

const (
	HIGH PinWriteType = iota
	LOW
)
