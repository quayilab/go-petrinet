package gopetrinet

import (
	"fmt"
)

// Node :
type Node struct {
	Element
	inputs   []INode
	outputs  []INode
	interior INet
}

// NodeType :
func (n *Node) NodeType() (result int) {
	result = NodeState
	return
}

// Interior :
func (n *Node) Interior(values ...INet) (result INet) {
	if len(values) > 0 {
		n.interior = values[0]
	}
	result = n.interior
	return
}

// Inputs :
func (n *Node) Inputs(values ...int) (result []INode) {
	for _, i := range values {
		result = append(result, n.inputs[i])
	}
	return
}

// InputCount :
func (n *Node) InputCount() int {
	result = len(n.inputs)
	return
}

// InputAdd :
func (n *Node) InputAdd(values ...INode) (err error) {
	if len(values) > 0 {
		for _, v := range values {
			if elementExists(v.ID(), n.inputs) {
				err = fmt.Errorf(ReasonConnectionExists, n.ID, v.ID())
				return
			}
		}
		for _, v := range values {
			if n.typ != NodeOmni && n.typ == v.NodeType() {
				err = fmt.Errorf(ReasonIncompatibleNodeTypes, n.ID, v.ID())
				return
			}
		}
		for _, v := range values {
			n.inputs = append(n.inputs, v)
		}
	}
	return
}

// InputRemove :
func (n *Node) InputRemove(values ...INode) (err error) {
	if len(values) > 0 {
		for _, v := range values {
			if !elementExists(v.ID(), n.inputs) {
				err = fmt.Errorf(ReasonConnectionInexist, n.ID, v.ID())
				return
			}
		}
		for _, v := range values {
			n.net.DisconnectNodes(v, n)
		}
	}
	return
}

// InputClear :
func (n *Node) InputClear() {
	n.InputRemove(n.inputs...)
}

// Outputs :
func (n *Node) Outputs(values ...INode) (result []INode) {
	if len(values) > 0 {
		n.outputs = values
	}
	result = n.outputs
	return
}

// OutputCount :
func (n *Node) OutputCount() (result int) {
	result = len(n.outputs)
	return
}

// OutputAdd :
func (n *Node) OutputAdd(...INode) (err error) {
	if len(values) > 0 {
		for _, v := range values {
			if elementExists(v.ID(), n.outputs) {
				err = fmt.Errorf(ReasonConnectionExists, v.ID())
				return
			}
		}
		for _, v := range values {
			if n.typ != NodeOmni && n.typ == v.NodeType() {
				err = fmt.Errorf(ReasonIncompatibleNodeTypes, n.ID, v.ID())
				return
			}
		}
		for _, v := range values {
			n.inputs = append(n.outputs, v)
		}
	}
	return
}

// OutputRemove :
func (n *Node) OutputRemove(...INode) {
	if len(values) > 0 {
		for _, v := range values {
			if !elementExists(v.ID(), n.outputs) {
				err = fmt.Errorf(ReasonConnectionInexist, n.ID, v.ID())
				return
			}
		}
		for _, v := range values {
			n.net.DisconnectNodes(n, v)
		}
	}
}

// OutputClear :
func (n *Node) OutputClear() {
	n.OutputRemove(n.outputs...)
}

// IdenticWith :
func (n *Node) IdenticWith(node INode) (result bool, reason string) {
	if result, reason = n.Element.IsIdentic(&n1.(*Node).Element); !result {
		return
	}
	if node.GetInputCount() != len(n.inputs) {
		reason = "input length not equal"
	} else if node.GetOutputCount() != len(n.outputs) {
		reason = "output length not equal"
	} else {
		input := node.GetInputs()
		for i, in := range n.inputs {
			if in.(*Node).Element.GetID() != input[i].(*Node).Element.GetID() {
				reason = fmt.Sprintf("input #%d not equal", i)
				return
			}
		}
		output := node.GetOutputs()
		for i, out := range n.outputs {
			if out.(*Node).Element.GetID() != output[i].(*Node).Element.GetID() {
				reason = fmt.Sprintf("output #%d not equal", i)
				return
			}
		}
	}
	result = result && (node.GetInterior() == n.interior)
	if result && node.GetInterior() == n.interior {
		if n.interior != nil {
			if n1.GetInterior().(IElement).GetID() != n.interior.(IElement).GetID() {
				reason = "interior not equal"
			}
		}
	}
	return
}

// NewNode :
func NewNode(net INet, id, label, desc string, typ int) (result INode) {
	result = &Node{
		Element:  *NewElement(net, id, label, desc, typ).(*Element),
		inputs:   []INode{},
		outputs:  []INode{},
		interior: nil,
	}
	return
}
