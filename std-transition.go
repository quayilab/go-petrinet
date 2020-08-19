package gopetrinet

import (
	"fmt"
)

// OnExec :
type OnExec func(tkns map[int][]IToken, outTkns *map[int][]IToken)

// OnTestState :
type OnTestState func(state IState) (result bool)

// Transition :
type Transition struct {
	Node
	activationTreshold map[string]int
	inputTokens        map[int][]IToken
	outputTokens       map[int][]IToken
	inputStates        map[int]IState
	outputStates       map[int]IState
	inputArcTypes      map[int]int
	onExecute          OnExec
	onTestState        OnTestState
}

// IdenticWith :
func (t *Transition) IdenticWith(transition ITransition) (identic bool, reason string) {
	if identic, reason = t.Node.IdenticWith(transition.(INode)); !identic {
		return
	}
	for tknType, count := range t.activationTreshold {
		if count != transition.ActivationTreshold(tknType) {
			reason = fmt.Sprintf("activation treshold #%d not equal", tknType)
			break
		}
	}
	return
}

// ConnectInput :
func (t *Transition) ConnectInput(state IState) {
	t.net.ConnectStateTransition(state, t)
}

// ConnectOutput :
func (t *Transition) ConnectOutput(state IState) {
	t.net.ConnectTransitionState(t, state)
}

// ActivationTreshold :
func (t *Transition) ActivationTreshold(tokenType string) (result int) {
	result = t.activationTreshold[tokenType]
	return result
}

// StateReady :
func (t *Transition) StateReady(state IState) (result bool) {
	if t.onTestState != nil {
		result = t.onTestState(state)
		return
	}
	switch state.StorageMode() {
	case StorageMultiset:
		for ttype, thold := range t.activationTreshold {
			if state.TokenCount(ttype) >= thold {
				result = true
				break
			}
		}
		break
	case StorageChanel:

		break
	case StorageStack:
		break
	}
	return
}

// Ready :
func (t *Transition) Ready() (result bool) {
	result = true
	for _, n := range t.inputs {
		result = result && t.StateReady(n.(IState))
	}
	result = result && (t.onExecute != nil)
	return
}

// Execute :
func (t *Transition) Execute() {
	if !t.Ready() {
		return
	}
	for stateID, state := range t.inputStates {
		state.
			t.inputTokens[stateID] = state.TokenFetch(t.activationTreshold[stateID])
	}
	t.onExecute(t.inputTokens, &t.outputTokens)
	t.inputTokens = map[int][]IToken{}
	t.DistributeTokens()
}

// DistributeTokens :
func (t *Transition) DistributeTokens() {
	for stateID, tkns := range t.outputTokens {
		state := t.outputStates[stateID]
		state.AddToken(tkns...)
	}
	t.outputTokens = map[int][]IToken{}
}

// NewTransition :
func NewTransition(net INet, id, label, desc string, onexec OnExec, onteststate OnTestState) (result ITransition) {
	result = &Transition{
		Node:               *NewNode(net, id, label, desc, ElementTypeNodeTransition, NodeTransition).(*Node),
		activationTreshold: map[int]int{},
		inputStates:        map[int]IState{},
		outputStates:       map[int]IState{},
		inputTokens:        map[int][]IToken{},
		outputTokens:       map[int][]IToken{},
		inputArcTypes:      map[int]int{},
		onExecute:          onexec,
		onTestState:        onteststate,
	}
	return
}
