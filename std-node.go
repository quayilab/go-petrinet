package gopetrinet

import (
	"fmt"
)

// Node :
type Node struct {
	Element
	kind     int
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
func (n *Node) InputCount() (result int) {
	result = len(n.inputs)
	return
}

// InputAdd :
func (n *Node) InputAdd(values ...INode) (err error) {
	if len(values) > 0 {
		for _, v := range values {
			if nodeExists(v.(IElement).ID(), n.inputs) {
				err = fmt.Errorf(ReasonConnectionExists, n.ID, v.(IElement).ID())
				return
			}
		}
		for _, v := range values {
			if n.kind != NodeOmni && n.kind == v.NodeType() {
				err = fmt.Errorf(ReasonIncompatibleNodeTypes, n.ID, v.(IElement).ID())
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
			if !nodeExists(v.(IElement).ID(), n.inputs) {
				err = fmt.Errorf(ReasonConnectionInexist, n.ID, v.(IElement).ID())
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
func (n *Node) OutputAdd(values ...INode) (err error) {
	if len(values) > 0 {
		for _, v := range values {
			if nodeExists(v.(IElement).ID(), n.outputs) {
				err = fmt.Errorf(ReasonConnectionExists, v.(IElement).ID())
				return
			}
		}
		for _, v := range values {
			if n.kind != NodeOmni && n.kind == v.NodeType() {
				err = fmt.Errorf(ReasonIncompatibleNodeTypes, n.ID, v.(IElement).ID())
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
func (n *Node) OutputRemove(values ...INode) (err error) {
	if len(values) > 0 {
		for _, v := range values {
			if !nodeExists(v.(IElement).ID(), n.outputs) {
				err = fmt.Errorf(ReasonConnectionInexist, n.ID, v.(IElement).ID())
				return
			}
		}
		for _, v := range values {
			n.net.DisconnectNodes(n, v)
		}
	}
	return
}

// OutputClear :
func (n *Node) OutputClear() {
	n.OutputRemove(n.outputs...)
}

// IdenticWith :
func (n *Node) IdenticWith(node INode) (result bool, reason string) {
	if result, reason = n.IdenticWith(node); !result {
		return
	}
	if node.InputCount() != len(n.inputs) {
		reason = "input length not equal"
	} else if node.OutputCount() != len(n.outputs) {
		reason = "output length not equal"
	} else {
		input := node.Inputs()
		for i, in := range n.inputs {
			if in.(IElement).ID() != input[i].(IElement).ID() {
				reason = fmt.Sprintf("input #%d not equal", i)
				return
			}
		}
		output := node.Outputs()
		for i, out := range n.outputs {
			if out.(IElement).ID() != output[i].(IElement).ID() {
				reason = fmt.Sprintf("output #%d not equal", i)
				return
			}
		}
	}
	result = result && (node.Interior() == n.interior)
	if result && node.Interior() == n.interior {
		if n.interior != nil {
			if node.Interior().(IElement).ID() != n.interior.(IElement).ID() {
				reason = "interior not equal"
			}
		}
	}
	return
}

// NewNode :
func NewNode(net INet, id, label, desc string, typ, kind int) (result INode) {
	result = &Node{
		Element:  *NewElement(net, id, label, desc, typ).(*Element),
		kind:     kind,
		inputs:   []INode{},
		outputs:  []INode{},
		interior: nil,
	}
	return
}
