package gopetrinet

import (
	"fmt"
)

// OnExec :
type OnExec func(tkns map[int][]TokenIntf, outTkns *map[int][]TokenIntf)

// Transition :
type Transition struct {
	Node
	activationTreshold map[int]int
	inputTokens        map[int][]TokenIntf
	outputTokens       map[int][]TokenIntf
	inputStates        map[int]StateIntf
	outputStates       map[int]StateIntf
	inputArcTypes      map[int]int
	onExecute          OnExec
}

// IsIdentic :
func (t *Transition) IsIdentic(t1 TransitionIntf) (result bool, reason string) {
	if result, reason = t.Node.IsIdentic(&t1.(*Transition).Node); !result {
		return
	}

	for tknType, count := range t.activationTreshold {
		if count != t1.GetActivationTreshold(tknType) {
			reason = fmt.Sprintf("activation treshold #%d not equal", tknType)
			break
		}
	}
	return
}

// GetActivationTreshold :
func (t *Transition) GetActivationTreshold(stateID int) (result int) {
	result = t.activationTreshold[stateID]
	return
}

// IsStateReady :
func (t *Transition) IsStateReady(s StateIntf) (result bool) {
	result = s.GetTokenCount() > t.activationTreshold[s.GetStateID()]
	return
}

// Execute :
func (t *Transition) Execute() {
	if !t.IsReady() {
		return
	}
	for stateID, state := range t.inputStates {
		t.inputTokens[stateID] = state.GetToken(t.activationTreshold[stateID])
	}
	t.onExecute(t.inputTokens, &t.outputTokens)
	t.inputTokens = map[int][]TokenIntf{}
	t.DistributeTokens()
}

// DistributeTokens :
func (t *Transition) DistributeTokens() {
	for stateID, tkns := range t.outputTokens {
		state := t.outputStates[stateID]
		state.AddToken(tkns...)
	}
	t.outputTokens = map[int][]TokenIntf{}
}

// IsReady :
func (t *Transition) IsReady() (result bool) {
	result = true
	for _, n := range t.inputs {
		result = result && t.IsStateReady(n.(StateIntf))
	}
	result = result && (t.onExecute != nil)
	return
}

// ConnectInput :
func (t *Transition) ConnectInput(s StateIntf, arctype, activationTreshold int) {
	t.Node.AddInput(&s.(*State).Node)
	s.(*State).Node.AddOutput(&t.Node)
	t.inputStates[s.GetStateID()] = s
	t.inputArcTypes[s.GetStateID()] = arctype
	t.activationTreshold[s.GetStateID()] = activationTreshold
}

// ConnectOutput :
func (t *Transition) ConnectOutput(s StateIntf) {
	t.Node.AddOutput(&s.(*State).Node)
	t.outputStates[s.GetStateID()] = s
}

// NewTransition :
func NewTransition(net NetIntf, id, label, desc string, onexec OnExec) (result TransitionIntf) {
	result = &Transition{
		Node:               *NewNode(net, id, label, desc, ElementTypeNodeTransition).(*Node),
		activationTreshold: map[int]int{},
		inputStates:        map[int]StateIntf{},
		outputStates:       map[int]StateIntf{},
		inputTokens:        map[int][]TokenIntf{},
		outputTokens:       map[int][]TokenIntf{},
		inputArcTypes:      map[int]int{},
		onExecute:          onexec,
	}
	return
}
