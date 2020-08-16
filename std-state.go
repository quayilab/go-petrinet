package gopetrinet

import (
	"fmt"
)

// State :
type State struct {
	Node
	capacity int
	stateID  int
	tokens   []IToken
}

// Capacity :
func (s *State) Capacity(values ...int) (result int) {
	result = s.capacity
	return
}

// StateID :
func (s *State) StateID(values ...int) (result int) {
	if len(values) > 0 {
		s.stateID = values[0]
	}
	result = s.stateID
	return
}

// Token :
func (s *State) Token(indexes ...int) (result []IToken) {
	if len(indexes) > 0 {
		for _, i := range indexes {
			result = append(result, s.tokens[i])
		}
	} else {
		result = s.tokens
	}
	return
}

// TokenCount :
func (s *State) TokenCount() (result int) {
	result = len(s.tokens)
	return
}

// TokenFetch :
func (s *State) TokenFetch(count int) []IToken {
	result = s.tokens[:count]
	s.tokens = s.tokens[count:]
	return
}

// TokenCopies :
func (s *State) TokenCopies() []IToken {
	result = s.tokens[:]
	return
}

// TokenAdd :
func (s *State) TokenAdd(values ...IToken) {
	s.tokens = append(s.tokens, values...)
}

// IdenticWith :
func (s *State) IdenticWith(state IState) (result bool, reason string) {
	if result, reason = s.Node.IdenticWith(&state.(*State).Node); !result {
		return
	}

	if s1.GetCapacity() != s.capacity {
		reason = "token count not equal"
	} else if state.GetTokenCount() != len(s.tokens) {
		reason = "token count not equal"
	} else {
		s1tokens := state.GetTokenCopies()
		for i, t := range s.tokens {
			t0 := t
			t1 := s1tokens[i]
			if t1.(IElement).GetID() != t0.(IElement).GetID() {
				reason = fmt.Sprintf("token #%d not equal", i)
				return
			}
		}
	}
	return
}

// Ready :
func (s *State) Ready(transition ITransition) (result bool) {
	result = transition.StateReady(s)
	return
}

// NewState :
func NewState(net INet, id, label, desc string, stateid, capacity int) (result IState) {
	result = &State{
		Node:     *NewNode(net, id, label, desc, ElementTypeNodeState).(*Node),
		capacity: capacity,
		stateID:  stateid,
		tokens:   []IToken{},
	}
	return
}
