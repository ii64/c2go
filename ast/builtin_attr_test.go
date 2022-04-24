package ast

import (
	"testing"
)

func TestBuiltinAttr(t *testing.T) {
	nodes := map[string]Node{
		"0x1503d00 <<invalid sloc>> Implicit 879": &BuiltinAttr{
			Addr:      0x1503d00,
			Pos:       NewPositionFromString("<invalid sloc>"),
			Implicit:  true,
			Inherited: false,
			Unknown1:  879,
		},
		"0x23c4e08 <<invalid sloc>> Inherited Implicit 823": &BuiltinAttr{
			Addr:      0x23c4e08,
			Pos:       NewPositionFromString("<invalid sloc>"),
			Implicit:  true,
			Inherited: true,
			Unknown1:  823,
		},
	}

	runNodeTests(t, nodes)
}
