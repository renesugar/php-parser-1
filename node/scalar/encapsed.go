package scalar

import (
	"github.com/z7zmey/php-parser/node"
)

type Encapsed struct {
	name       string
	attributes map[string]interface{}
	position   *node.Position
	parts      []node.Node
}

func NewEncapsed(parts []node.Node) node.Node {
	return Encapsed{
		"Encapsed",
		map[string]interface{}{},
		nil,
		parts,
	}
}

func (n Encapsed) Name() string {
	return "Encapsed"
}

func (n Encapsed) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Encapsed) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Encapsed) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
}

func (n Encapsed) Position() *node.Position {
	return n.position
}

func (n Encapsed) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Encapsed) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.parts != nil {
		vv := v.GetChildrenVisitor("parts")
		for _, nn := range n.parts {
			nn.Walk(vv)
		}
	}
}
