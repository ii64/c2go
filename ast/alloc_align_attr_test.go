package ast

import "testing"

func TestAllocAlignAttr(t *testing.T) {
	nodes := map[string]Node{
		"0x18d3d00 <col:14> Implicit 1": &AllocAlignAttr{
			Addr:      0x18d3d00,
			Pos:       NewPositionFromString("col:14"),
			Implicit:  true,
			Inherited: false,
			Unknown1:  1,
		},
	}

	runNodeTests(t, nodes)
}
