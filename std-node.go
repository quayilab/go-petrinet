package gopetrinet

import (
	"fmt"
)

type Node struct {
	Element
	inputs   []NodeIntf
	outputs  []NodeIntf
	interior NetIntf
}

// GetNodeType :
func (n *Node) GetNodeType() (result int) {
	result = NodeState
	return
}

// GetInputCount :
func (n *Node) GetInputCount() (result int) {
	result = len(n.inputs)
	return
}

// GetOutputCount :
func (n *Node) GetOutputCount() (result int) {
	result = len(n.outputs)
	return
}

// GetInputs :
func (n *Node) GetInputs() (result []NodeIntf) {
	result = n.inputs
	return
}

// GetOutputs :
func (n *Node) GetOutputs() (result []NodeIntf) {
	result = n.outputs
	return
}

// AddInput :
func (n *Node) AddInput(n1 ...NodeIntf) {
	n.inputs = append(n.inputs, n1...)
}

// AddOutput :
func (n *Node) AddOutput(n1 ...NodeIntf) {
	n.outputs = append(n.outputs, n1...)
}

// GetInterior :
func (n *Node) GetInterior() (result NetIntf) {
	result = n.interior
	return
}

// SetInterior :
func (n *Node) SetInterior(net NetIntf) {
	n.interior = net
}

// IsIdentic :
func (n *Node) IsIdentic(n1 NodeIntf) (result bool, reason string) {
	if result, reason = n.Element.IsIdentic(&n1.(*Node).Element); !result {
		return
	}
	if n1.GetInputCount() != len(n.inputs) {
		reason = "input lenght not equal"
	} else if n1.GetOutputCount() != len(n.outputs) {
		reason = "output lenght not equal"
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
			if n1.GetInterior().(ElementIntf).GetID() != n.interior.(ElementIntf).GetID() {
				reason = "interior not equal"
			}
		}
	}
	return
}

func NewNode(net NetIntf, id, label, desc string, typ int) (result NodeIntf) {
	result = &Node{
		Element:  *NewElement(net, id, label, desc, typ).(*Element),
		inputs:   []NodeIntf{},
		outputs:  []NodeIntf{},
		interior: nil,
	}
	return
}
