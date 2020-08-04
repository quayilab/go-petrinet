package gopetrinet

import (
	"fmt"
)

// State :
type State struct {
	Node
	capacity int
	stateID  int
	tokens   []TokenIntf
}

// GetCapacity :
func (s *State) GetCapacity() (result int) {
	result = s.capacity
	return
}

// SetCapacity :
func (s *State) SetCapacity(v int) {
	s.capacity = v
}

// GetStateID :
func (s *State) GetStateID() (result int) {
	result = s.stateID
	return
}

// SetStateID :
func (s *State) SetStateID(v int) {
	s.stateID = v
}

// GetTokenCount :
func (s *State) GetTokenCount() (result int) {
	result = len(s.tokens)
	return
}

// GetToken :
func (s *State) GetToken(c int) (result []TokenIntf) {
	result = s.tokens[:c]
	s.tokens = s.tokens[c:]
	return
}

// GetTokenCopies :
func (s *State) GetTokenCopies() (result []TokenIntf) {
	result = s.tokens[:]
	return
}

// AddToken :
func (s *State) AddToken(t ...TokenIntf) {
	s.tokens = append(s.tokens, t...)
}

// IsReady :
func (s *State) IsReady(t TransitionIntf) (result bool) {
	result = t.IsStateReady(s)
	return
}

// IsIdentic :
func (s *State) IsIdentic(s1 StateIntf) (result bool, reason string) {
	if result, reason = s.Node.IsIdentic(&s1.(*State).Node); !result {
		return
	}

	if s1.GetCapacity() != s.capacity {
		reason = "token count not equal"
	} else if s1.GetTokenCount() != len(s.tokens) {
		reason = "token count not equal"
	} else {
		s1tokens := s1.GetTokenCopies()
		for i, t := range s.tokens {
			t0 := t
			t1 := s1tokens[i]
			if t1.(ElementIntf).GetID() != t0.(ElementIntf).GetID() {
				reason = fmt.Sprintf("token #%d not equal", i)
				return
			}
		}
	}
	return
}

// NewState :
func NewState(net NetIntf, id, label, desc string, stateid, capacity int) (result StateIntf) {
	result = &State{
		Node:     *NewNode(net, id, label, desc, ElementTypeNodeState).(*Node),
		capacity: capacity,
		stateID:  stateid,
		tokens:   []TokenIntf{},
	}
	return
}
