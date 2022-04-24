package ast

import "github.com/elliotchance/c2go/util"

// BuiltinAttr is a type of attribute that is ....
type BuiltinAttr struct {
	Addr      Address
	Pos       Position
	Implicit  bool
	Inherited bool
	Unknown1  int

	ChildNodes []Node
}

var _ Node = (*BuiltinAttr)(nil)

func parseBuiltinAttr(line string) *BuiltinAttr {
	groups := groupsFromRegex(
		`<(?P<position>.*)>
		(?P<inherited> Inherited)?(?P<implicit> Implicit)?
		 (?P<unknown1>\d+)`,
		line,
	)
	return &BuiltinAttr{
		Addr:      ParseAddress(groups["address"]),
		Pos:       NewPositionFromString(groups["position"]),
		Implicit:  len(groups["implicit"]) > 0,
		Inherited: len(groups["inherited"]) > 0,
		Unknown1:  util.Atoi(groups["unknown1"]),
	}
}

// AddChild adds a new child node. Child nodes can then be accessed with the
// Children attribute.
func (n *BuiltinAttr) AddChild(node Node) {
	n.ChildNodes = append(n.ChildNodes, node)
}

// Address returns the numeric address of the node. See the documentation for
// the Address type for more information.
func (n *BuiltinAttr) Address() Address {
	return n.Addr
}

func (n *BuiltinAttr) Children() []Node {
	return n.ChildNodes
}

// Position returns the position in the original source code.
func (n *BuiltinAttr) Position() Position {
	return n.Pos
}
