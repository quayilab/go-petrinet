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
	activationTreshold map[int]int
	inputTokens        map[int][]IToken
	outputTokens       map[int][]IToken
	inputStates        map[int]IState
	outputStates       map[int]IState
	inputArcTypes      map[int]int
	onExecute          OnExec
	onTestState        OnTestState
}

// IdenticWith :
func (t *Transition) IdenticWith(transition ITransition) (result bool, reason string) {
	if result, reason = t.Node.IsIdentic(&transition.(*Transition).Node); !result {
		return
	}
	for tknType, count := range t.activationTreshold {
		if count != transition.GetActivationTreshold(tknType) {
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

// StateReady :
func (t *Transition) StateReady(state IState) (result bool) {
	if t.OnTestState(state) != nil {
		result = t.OnTestState(state)
		return
	}
	result = t.activationTreshold[state.StateID] <= state.TokenCount
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
	if !t.IsReady() {
		return
	}
	for stateID, state := range t.inputStates {
		t.inputTokens[stateID] = state.GetToken(t.activationTreshold[stateID])
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
		Node:               *NewNode(net, id, label, desc, ElementTypeNodeTransition).(*Node),
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
