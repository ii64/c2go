package ast

import "testing"

func TestInt(t *testing.T) {
	nodes := map[string]Node{
		"33554432": &Int{
			Value: 33554432,
		},
	}

	runNodeTests(t, nodes)
}
