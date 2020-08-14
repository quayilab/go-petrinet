package gopetrinet

import (
	"fmt"
)

// Node :
type Node struct {
	Element
	inputs   []*INode
	outputs  []*INode
	interior *INet
}

// NodeType :
func (n *Node) NodeType() (result int) {
	result = NodeState
	return
}

// Interior :
func (n *Node) Interior(values ...*INet) (result *INet) {
	if len(values) > 0 {
		n.interior = values[0]
	}
	result = n.interior
	return
}

// Inputs :
func (n *Node) Inputs(values ...*INode) (result []*INode) {
	if len(values) > 0 {
		n.inputs = values
	}
	result = n.inputs
	return
}

// InputCount :
func (n *Node) InputCount() int {
	result = len(n.inputs)
	return
}

// InputAdd :
func (n *Node) InputAdd(values ...INode) {
	if len(values) > 0 {
		for _, v := range values {
			if n.typ == NodeState && (v.NodeType() == NodeTransition ||
				v.NodeType() == NodeOmni) {
				n.net.ConnectStateTransition(v, n)
			} else if n.typ == NodeTransition && (v.NodeType() == NodeState ||
				v.NodeType() == NodeOmni) {
				n.net.ConnectTransitionState(v, n)
			}
		}
	}
}

// InputRemove :
func (n *Node) InputRemove(values ...INode) {
	if len(values) > 0 {
		for _, v := range values {
			n.net.DisconnectNodes(v, n)
		}
	}
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
func (n *Node) OutputAdd(...INode) {
	if len(values) > 0 {
		for _, v := range values {
			if n.typ == NodeState && (v.NodeType() == NodeTransition ||
				v.NodeType() == NodeOmni) {
				n.net.ConnectStateTransition(n, v)
			} else if n.typ == NodeTransition && (v.NodeType() == NodeState ||
				v.NodeType() == NodeOmni) {
				n.net.ConnectTransitionState(n, v)
			}
		}
	}
}

// OutputRemove :
func (n *Node) OutputRemove(...INode) {
	if len(values) > 0 {
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
func (n *Node) IdenticWith(n1 INode) (result bool, reason string) {
	if result, reason = n.Element.IsIdentic(&n1.(*Node).Element); !result {
		return
	}
	if n1.GetInputCount() != len(n.inputs) {
		reason = "input length not equal"
	} else if n1.GetOutputCount() != len(n.outputs) {
		reason = "output length not equal"
	} else {
		input := n1.GetInputs()
		for i, in := range n.inputs {
			if in.(*Node).Element.GetID() != input[i].(*Node).Element.GetID() {
				reason = fmt.Sprintf("input #%d not equal", i)
				return
			}
		}
		output := n1.GetOutputs()
		for i, out := range n.outputs {
			if out.(*Node).Element.GetID() != output[i].(*Node).Element.GetID() {
				reason = fmt.Sprintf("output #%d not equal", i)
				return
			}
		}
	}
	result = result && (n1.GetInterior() == n.interior)
	if result && n1.GetInterior() == n.interior {
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
