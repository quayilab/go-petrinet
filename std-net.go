package gopetrinet

import (
	"fmt"
	// "sync"
)

// Net struct
type Net struct {
	Element
	nodes       []INode
	states      []IState
	transitions []ITransition
	paused      bool
	stepCount   int
}

// NodeCount :
func (n *Net) NodeCount() (result int) {
	result = len(n.nodes)
	return
}

// Nodes :
func (n *Net) Nodes() (result []INode) {
	result = n.nodes
	return
}

// Node :
func (n *Net) Node(index int) (result INode) {
	result = nil
	if index < 0 || index >= len(n.nodes) {
		return
	}
	result = n.nodes[index]
	return
}

// StateCount :
func (n *Net) StateCount() (result int) {
	result = len(n.states)
	return
}

// States :
func (n *Net) States() (result []IState) {
	result = n.states
	return
}

// State :
func (n *Net) State(index int) (result IState) {
	result = nil
	if index < 0 || index >= len(n.states) {
		return
	}
	result = n.states[index]
	return
}

// StateAdd :
func (n *Net) StateAdd(id string, label string, description string, stateid int, capacity int) (result IState) {
	result = NewState(n, id, label, description, stateid, capacity)
	return
}

// StateAddToken :
func (n *Net) StateAddToken(state IState, tokens ...IToken) {
	state.TokenAdd(tokens...)
}

// TransitionCount :
func (n *Net) TransitionCount() (result int) {
	result = len(n.transitions)
	return
}

// Transitions :
func (n *Net) Transitions() (result []ITransition) {
	result = n.transitions
	return
}

// Transition :
func (n *Net) Transition(index int) (result ITransition) {
	result = nil
	if index < 0 || index >= len(n.transitions) {
		return
	}
	result = n.transitions[index]
	return
}

// TransitionAdd :
func (n *Net) TransitionAdd(id string, label string, description string, onExec OnExec, onTestState OnTestState) (result ITransition) {
	result = NewTransition(n, id, label, description, onExec, onTestState)
	return
}

// TransitionReady :
func (n *Net) TransitionReady(transition ITransition) (result bool) {
	result = transition.Ready()
	return
}

// ConnectStateTransition :
func (n *Net) ConnectStateTransition(state IState, transition ITransition) (err error) {
	err = state.(INode).OutputAdd(transition.(INode))
	return
}

// ConnectTransitionState :
func (n *Net) ConnectTransitionState(state IState, transition ITransition) (err error) {
	err = state.(INode).InputAdd(transition.(INode))
	return
}

// ConnectNodes :
func (n *Net) ConnectNodes(node1, node2 INode) (err error) {
	err = node1.OutputAdd(node2)
	return
}

// TryConnectNodes :
func (n *Net) TryConnectNodes(node1, node2 INode) (err error) {
	return
}

// DisconnectNodes :
func (n *Net) DisconnectNodes(node1, node2 INode) (err error) {
	node1.OutputRemove(node2)
	return
}

// TryDisconnectNodes :
func (n *Net) TryDisconnectNodes(node1, node2 INode) (err error) {
	return
}

// IdenticWith :
func (n *Net) IdenticWith(net INet) (identic bool, reason string) {
	if identic, reason = n.Element.IdenticWith(&net.(*Net).Element); !identic {
		return
	}
	if net.NodeCount() != len(n.nodes) {
		reason = "nodes length not equal"
	} else if net.StateCount() != len(n.states) {
		reason = "states length not equal"
	} else if net.TransitionCount() != len(n.transitions) {
		reason = "transitions length not equal"
	} else {
		nodes := net.Nodes()
		for i, node := range n.nodes {
			if node.(*Node).Element.ID() != nodes[i].(IElement).ID() {
				reason = fmt.Sprintf("node #%d not equal", i)
				return
			}
		}
		states := net.States()
		for i, state := range n.states {
			if state.(*State).Node.Element.ID() != states[i].(IElement).ID() {
				reason = fmt.Sprintf("state #%d not equal", i)
				return
			}
		}
		transitions := net.Transitions()
		for i, transition := range n.transitions {
			if transition.(*Transition).Node.Element.ID() != transitions[i].(IElement).ID() {
				reason = fmt.Sprintf("transition #%d not equal", i)
				return
			}
		}
	}
	return
}

// Run :
func (n *Net) Run() {

}

// Step :
func (n *Net) Step() {

}

// Pause :
func (n *Net) Pause() {

}

// Running :
func (n *Net) Running() (result bool) {

}

// Ready :
func (n *Net) Ready() (result bool) {

}

// DeadLock :
func (n *Net) DeadLock() (result bool) {

}

// Execute :
func (n *Net) Execute(transition ITransition) (result bool, err error) {

}

// NewNet :
func NewNet(id, label, desc string) (result INet) {
	result = &Net{
		nodes:       []INode{},
		states:      []IState{},
		transitions: []ITransition{},
	}
	result.(*Net).Element = *NewElement(nil, id, label, desc, ElementTypeNet).(*Element)
	return
}
