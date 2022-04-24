package ast

import "github.com/elliotchance/c2go/util"

type Int struct {
	Value      int
	ChildNodes []Node
}

func parseInt(line string) *Int {
	// hacks for address parsing
	line = "0x0 " + line
	groups := groupsFromRegex(
		`(?P<value>\d+)`,
		line,
	)
	return &Int{Value: util.Atoi(groups["value"])}
}

var _ Node = (*Int)(nil)

// AddChild adds a new child node. Child nodes can then be accessed with the
// Children attribute.
func (n *Int) AddChild(node Node) {
	n.ChildNodes = append(n.ChildNodes, node)
}

// Address returns the numeric address of the node. See the documentation for
// the Address type for more information.
func (n *Int) Address() Address {
	return ParseAddress("0")
}

// Children returns the child nodes. If this node does not have any children or
// this node does not support children it will always return an empty slice.
func (n *Int) Children() []Node {
	return n.ChildNodes
}

// Position returns the position in the original source code.
func (n *Int) Position() Position {
	return NewPositionFromString("<invalid sloc>")
}
