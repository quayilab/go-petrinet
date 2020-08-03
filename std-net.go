package gopetrinet

import (
	"fmt"
	_ "sync"
)

type Net struct {
	Element
	nodes       []NodeIntf
	states      []StateIntf
	transitions []TransitionIntf
	paused      bool
	stepCount   int
}

// GetNodeCount :
func (n *Net) GetNodeCount() (result int) {
	result = len(n.nodes)
	return
}

// GetStateCount :
func (n *Net) GetStateCount() (result int) {
	result = len(n.states)
	return
}

// GetTransitionCount :
func (n *Net) GetTransitionCount() (result int) {
	result = len(n.transitions)
	return
}

// GetNode :
func (n *Net) GetNode(index int) (result NodeIntf) {
	result = nil
	if index < 0 || index >= len(n.nodes) {
		return
	}
	result = n.nodes[index]
	return
}

// GetNodes :
func (n *Net) GetNodes() (result []NodeIntf) {
	result = n.nodes
	return
}

// GetState :
func (n *Net) GetState(index int) (result StateIntf) {
	result = nil
	if index < 0 || index >= len(n.states) {
		return
	}
	result = n.states[index]
	return
}

// GetStates :
func (n *Net) GetStates() (result []StateIntf) {
	result = n.states
	return
}

// GetTransition :
func (n *Net) GetTransition(index int) (result TransitionIntf) {
	result = nil
	if index < 0 || index >= len(n.transitions) {
		return
	}
	result = n.transitions[index]
	return
}

// GetStates :
func (n *Net) GetTransitions() (result []TransitionIntf) {
	result = n.transitions
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

// IsReady :
func (n *Net) IsReady() (result bool) {

	return
}

// IsDeadLock :
func (n *Net) IsDeadLock() (result bool) {

	return
}

// IsIdentic :
func (n *Net) IsIdentic(n1 NetIntf) (result bool, reason string) {
	if result, reason = n.Element.IsIdentic(&n1.(*Net).Element); !result {
		return
	}
	if n1.GetNodeCount() != len(n.nodes) {
		reason = "nodes length not equal"
	} else if n1.GetStateCount() != len(n.states) {
		reason = "states length not equal"
	} else if n1.GetTransitionCount() != len(n.transitions) {
		reason = "transitions length not equal"
	} else {
		nodes := n1.GetNodes()
		for i, node := range n.nodes {
			if node.(*Node).Element.GetID() != nodes[i].(ElementIntf).GetID() {
				reason = fmt.Sprintf("node #%d not equal", i)
				return
			}
		}
		states := n1.GetStates()
		for i, state := range n.states {
			if state.(*State).Node.Element.GetID() != states[i].(ElementIntf).GetID() {
				reason = fmt.Sprintf("state #%d not equal", i)
				return
			}
		}
		transitions := n1.GetTransitions()
		for i, transition := range n.transitions {
			if transition.(*Transition).Node.Element.GetID() != transitions[i].(ElementIntf).GetID() {
				reason = fmt.Sprintf("transition #%d not equal", i)
				return
			}
		}
	}
	return
}

// IsTransitionReady :
func (n *Net) IsTransitionReady(t TransitionIntf) (result bool) {

	return
}

// Execute :
func (n *Net) Execute(t TransitionIntf) (result bool, err error) {

	return
}

// AddState :
func (n *Net) AddState(id, label, desc string, stateid, capacity int) (result StateIntf) {
	result = NewState(n, id, label, desc, stateid, capacity)
	n.nodes = append(n.nodes, &result.(*State).Node)
	n.states = append(n.states, result)
	return
}

// AddTransition :
func (n *Net) AddTransition(id, label, desc string, onexec OnExec) (result TransitionIntf) {
	result = NewTransition(n, id, label, desc, onexec)
	n.nodes = append(n.nodes, &result.(*Transition).Node)
	n.transitions = append(n.transitions, result)
	return
}

// ConnectStateTransition :
func (n *Net) ConnectStateTransition(s StateIntf, t TransitionIntf, arctype, activationTreshold int) {
	t.ConnectInput(s, arctype, activationTreshold)
}

// ConnectTransitionState :
func (n *Net) ConnectTransitionState(s StateIntf, t TransitionIntf) {
	t.ConnectOutput(s)
}

// AddTokenToState :
func (n *Net) AddTokenToState(s StateIntf, tokens ...TokenIntf) {

}

func NewNet(id, label, desc string) (result NetIntf) {
	result = &Net{
		nodes:       []NodeIntf{},
		states:      []StateIntf{},
		transitions: []TransitionIntf{},
	}
	result.(*Net).Element = *NewElement(nil, id, label, desc, ElementTypeNet).(*Element)
	return
}
