package ast

import "github.com/elliotchance/c2go/util"

type AllocAlignAttr struct {
	Addr      Address
	Pos       Position
	Implicit  bool
	Inherited bool
	Unknown1  int

	ChildNodes []Node
}

var _ Node = (*AllocAlignAttr)(nil)

func parseAllocAlignAttr(line string) *AllocAlignAttr {
	groups := groupsFromRegex(
		`<(?P<position>.*)>
		(?P<inherited> Inherited)?(?P<implicit> Implicit)?
		 (?P<unknown1>\d+)`,
		line,
	)

	return &AllocAlignAttr{
		Addr:      ParseAddress(groups["address"]),
		Pos:       NewPositionFromString(groups["position"]),
		Implicit:  len(groups["implicit"]) > 0,
		Inherited: len(groups["inherited"]) > 0,
		Unknown1:  util.Atoi(groups["unknown1"]),
	}
}

// AddChild adds a new child node. Child nodes can then be accessed with the
// Children attribute.
func (n *AllocAlignAttr) AddChild(node Node) {
	n.ChildNodes = append(n.ChildNodes, node)
}

// Address returns the numeric address of the node. See the documentation for
// the Address type for more information.
func (n *AllocAlignAttr) Address() Address {
	return n.Addr
}

func (n *AllocAlignAttr) Children() []Node {
	return n.ChildNodes
}

// Position returns the position in the original source code.
func (n *AllocAlignAttr) Position() Position {
	return n.Pos
}
