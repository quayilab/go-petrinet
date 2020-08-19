package gopetrinet

import (
	"fmt"
)

// State :
type State struct {
	Node
	capacity    int
	stateID     int
	tokens      []IToken
	tokenTypes  map[string]int
	storageMode int
	channel     chan IToken
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

// StorageMode :
func (s *State) StorageMode(values ...int) (result int) {
	if len(values) > 0 {
		s.storageMode = values[0]
	}
	result = s.storageMode
	return
}

// TokenTypes :
func (s *State) TokenTypes(tokenTypes ...string) (result []string) {
	if len(tokenTypes) > 0 {
		for _, ttype := range tokenTypes {
			v := s.tokenTypes[ttype]
			s.tokenTypes[ttype] = v
		}
	}
	for k := range s.tokenTypes {
		result = append(result, k)
	}
	return
}

// Token :
/*func (s *State) Token(indexes ...int) (result []IToken) {
	if len(indexes) > 0 {
		for _, i := range indexes {
			result = append(result, s.tokens[i])
		}
	} else {
		result = s.tokens
	}
	return
}*/

// TokenCount :
func (s *State) TokenCount(tokenType ...string) (result int) {
	if len(tokenType) > 0 {
		result = s.tokenTypes[tokenType[0]]
	} else {
		result = len(s.tokens)
	}
	return
}

// TokenPeek :
func (s *State) TokenPeek(count int) (result []IToken) {
	switch s.storageMode {
	case StorageChanel, StorageStack:
		result = s.tokens[:count]
		break
	case StorageMultiset:
		break
	}
	return
}

// TokenFetch :
func (s *State) TokenFetch(count int) (result []IToken) {
	switch s.storageMode {
	case StorageChanel:
		s.tokens = s.tokens[count:]
		break
	case StorageStack:
		s.tokens = s.tokens[:len(s.tokens)-count]
		break
	case StorageMultiset:
		break
	}
	for i := 0; i < count; i++ {
		result = append(result, <-s.channel)
	}
	return
}

// TokenCopies :
/*func (s *State) TokenCopies() (result []IToken) {
	result = s.tokens[:]
	return
}*/

// TokenAdd :
func (s *State) TokenAdd(values ...IToken) {
	switch s.storageMode {
	case StorageChanel:
		s.tokens = append(s.tokens, values...)
		break
	case StorageStack:
		for i := len(values)/2 - 1; i >= 0; i-- {
			opp := len(values) - 1 - i
			values[i], values[opp] = values[opp], values[i]
		}
		s.tokens = append(values, s.tokens...)
		break
	case StorageMultiset:
		break
	}
	for tt, tc := range s.tokenTypes {
		for _, t := range s.tokens {
			if t.Type() == tt {
				s.tokenTypes[tt] = tc + 1
			}
		}
	}
}

// IdenticWith :
func (s *State) IdenticWith(state IState) (identic bool, reason string) {
	if identic, reason = s.Node.IdenticWith(&state.(*State).Node); !identic {
		return
	}

	if state.Capacity() != s.capacity {
		reason = "token count not equal"
	} else if state.TokenCount() != len(s.tokens) {
		reason = "token count not equal"
	} else {
		s1tokens := state.TokenPeek(state.TokenCount())
		for i, t := range s.tokens {
			t0 := t
			t1 := s1tokens[i]
			if t1.(IElement).ID() != t0.(IElement).ID() {
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
		Node:        *NewNode(net, id, label, desc, ElementNode, NodeState).(*Node),
		capacity:    capacity,
		stateID:     stateid,
		storageMode: StorageChanel,
		tokens:      []IToken{},
		tokenTypes:  map[string]int{},
		channel:     make(chan IToken, capacity),
	}
	return
}
