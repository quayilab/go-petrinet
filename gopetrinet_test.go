package gopetrinet

import (
	"testing"
)

var (
	expNet = &Net{
		Element: Element{
			id:      "testing-net",
			label:   "Testing Net",
			desc:    "description of Testing Net",
			typ:     ElementNet,
			net:     nil,
			enabled: true,
		},
		nodes:       []INode{},
		states:      []IState{},
		transitions: []ITransition{},
	}

	expState01 = &State{
		Node: Node{
			Element: Element{
				id:      "testing-state-01",
				label:   "Testing State #01",
				desc:    "description of Testing State #01",
				typ:     ElementNode,
				net:     expNet,
				enabled: true,
			},
			kind:     NodeState,
			inputs:   []INode{},
			outputs:  []INode{},
			interior: nil,
		},
		capacity: 10,
		stateID:  1000,
		tokens:   []IToken{},
	}
	expState02 = &State{
		Node: Node{
			Element: Element{
				id:      "testing-state-02",
				label:   "Testing State #02",
				desc:    "description of Testing State #02",
				typ:     ElementNode,
				net:     expNet,
				enabled: true,
			},
			kind:     NodeState,
			inputs:   []INode{},
			outputs:  []INode{},
			interior: nil,
		},
		capacity: 10,
		stateID:  1001,
		tokens:   []IToken{},
	}
	expState03 = &State{
		Node: Node{
			Element: Element{
				id:      "testing-state-03",
				label:   "Testing State #03",
				desc:    "description of Testing State #03",
				typ:     ElementNode,
				net:     expNet,
				enabled: true,
			},
			kind:     NodeState,
			inputs:   []INode{},
			outputs:  []INode{},
			interior: nil,
		},
		capacity: 10,
		stateID:  1002,
		tokens:   []IToken{},
	}

	expTransition01 = &Transition{
		Node: Node{
			Element: Element{
				id:      "testing-transition-01",
				label:   "Testing Transition #01",
				desc:    "description of Testing Transition #01",
				typ:     ElementNode,
				net:     expNet,
				enabled: true,
			},
			kind:     NodeTransition,
			inputs:   []INode{},
			outputs:  []INode{},
			interior: nil,
		},
		activationTreshold: map[string]int{},
		inputTokens:        map[int][]IToken{},
		outputTokens:       map[int][]IToken{},
		inputStates:        map[int]IState{},
		outputStates:       map[int]IState{},
		inputArcTypes:      map[int]int{},
		onExecute:          nil,
	}
	expTransition02 = &Transition{
		Node: Node{
			Element: Element{
				id:      "testing-transition-02",
				label:   "Testing Transition #02",
				desc:    "description of Testing Transition #02",
				typ:     ElementNode,
				net:     expNet,
				enabled: true,
			},
			inputs:   []INode{},
			outputs:  []INode{},
			interior: nil,
		},
		activationTreshold: map[string]int{},
		inputTokens:        map[int][]IToken{},
		outputTokens:       map[int][]IToken{},
		inputStates:        map[int]IState{},
		outputStates:       map[int]IState{},
		inputArcTypes:      map[int]int{},
		onExecute:          nil,
	}
)

func initConnetStateTransition() {
	expNet.nodes = append(expNet.nodes, &expState01.Node, &expState02.Node, &expTransition01.Node, &expTransition02.Node)
	expNet.states = append(expNet.states, expState01, expState02)
	expNet.transitions = append(expNet.transitions, expTransition01, expTransition02)

	expState01.outputs = append(expState01.outputs, &expTransition01.Node, &expTransition02.Node)
	expState02.outputs = append(expState02.outputs, &expTransition01.Node, &expTransition02.Node)

	expTransition01.inputs = append(expTransition01.inputs, &expState01.Node, &expState02.Node)
	expTransition01.inputStates[1000] = expState01
	expTransition01.inputStates[1001] = expState02
	expTransition01.inputArcTypes[1000] = ElementTypeArcNormal
	expTransition01.inputArcTypes[1001] = ElementTypeArcNormal
	expTransition01.activationTreshold[1000] = 1
	expTransition01.activationTreshold[1001] = 1

	expTransition02.inputs = append(expTransition02.inputs, &expState01.Node, &expState02.Node)
	expTransition02.inputStates[1000] = expState01
	expTransition02.inputStates[1001] = expState02
	expTransition02.inputArcTypes[1000] = ElementTypeArcNormal
	expTransition02.inputArcTypes[1001] = ElementTypeArcInhibit
	expTransition02.activationTreshold[1000] = 1
	expTransition02.activationTreshold[1001] = 1
}

func initConnetTransitionState() {
	expNet.nodes = append(expNet.nodes, &expState03.Node, &expTransition01.Node, &expTransition02.Node)
	expNet.states = append(expNet.states, expState03)
	expNet.transitions = append(expNet.transitions, expTransition01, expTransition02)

	expState03.inputs = append(expState03.inputs, &expTransition01.Node, &expTransition02.Node)

	expTransition01.outputs = append(expTransition01.outputs, &expState03.Node)
	expTransition01.outputStates[1002] = expState03
	expTransition01.inputArcTypes[1002] = ElementTypeArcNormal

	expTransition02.outputs = append(expTransition02.inputs, &expState03.Node)
	expTransition02.outputStates[1002] = expState03
	expTransition02.inputArcTypes[1002] = ElementTypeArcNormal
}

func TestIncorrectElements(t *testing.T) {
	s := &State{}
	*s = *expState01
	s.Node.Element.id = "different-id"
	got := NewState(expNet, "testing-state-01", "Testing State #01", "description of Testing State #01", 1000, 10)
	if result, reason := s.IsIdentic(got); result {
		t.Errorf("result is incorrect: %s\n  want : %v\n   got : %v\n", reason, *s, *got.(*State))
	}
	s = &State{}
	*s = *expState01
	s.Node.Element.typ = ElementTypeNet
	got = NewState(expNet, "testing-state-01", "Testing State #01", "description of Testing State #01", 1000, 10)
	if result, reason := s.IsIdentic(got); result {
		t.Errorf("result is incorrect: %s\n  want : %v\n   got : %v\n", reason, *s, *got.(*State))
	}
	s = &State{}
	*s = *expState01
	s.Node.Element.label = "different-label"
	got = NewState(expNet, "testing-state-01", "Testing State #01", "description of Testing State #01", 1000, 10)
	if result, reason := s.IsIdentic(got); result {
		t.Errorf("result is incorrect: %s\n  want : %v\n   got : %v\n", reason, *s, *got.(*State))
	}
	s = &State{}
	*s = *expState01
	s.Node.Element.desc = "different-desc"
	got = NewState(expNet, "testing-state-01", "Testing State #01", "description of Testing State #01", 1000, 10)
	if result, reason := s.IsIdentic(got); result {
		t.Errorf("result is incorrect: %s\n  want : %v\n   got : %v\n", reason, *s, *got.(*State))
	}
	s = &State{}
	*s = *expState01
	s.Node.Element.Enable()
	got = NewState(expNet, "testing-state-01", "Testing State #01", "description of Testing State #01", 1000, 10)
	if result, reason := s.IsIdentic(got); !result {
		t.Errorf("result is incorrect: %s\n  want : %v\n   got : %v\n", reason, *s, *got.(*State))
	}
	s.Node.Element.Disable()
	got = NewState(expNet, "testing-state-01", "Testing State #01", "description of Testing State #01", 1000, 10)
	if result, reason := s.IsIdentic(got); result {
		t.Errorf("result is incorrect: %s\n  want : %v\n   got : %v\n", reason, *s, *got.(*State))
	}
}

func TestNode(t *testing.T) {
	initConnetStateTransition()
	s := &State{}
	*s = *expState01
	//s.Node.interior = &Net{}
	s.Node.outputs = append(s.Node.outputs, &Node{})
	got := NewState(expNet, "testing-state-01", "Testing State #01", "description of Testing State #01", 1000, 10)
	if result, reason := s.IsIdentic(got); !result {
		t.Errorf("result is incorrect: %s\n  want : %v\n   got : %v\n", reason, *s, *got.(*State))
	}
	/*s= nil

	s= &State{}
	*s = *expState01
	s.Node.outputs[0].(*Node).Element.id = "some-different-id"
	got = NewState(expNet, "testing-state-01", "Testing State #01", "description of Testing State #01", 1000, 10)
	if result, reason := s.IsIdentic(got); result {
		t.Errorf("result is incorrect: %s\n  want : %v\n   got : %v\n", reason, *s, *got.(*State))
	}*/
}

func TestCreateNet(t *testing.T) {
	want := expNet
	got := NewNet("testing-net", "Testing Net", "description of Testing Net")
	if result, reason := want.IsIdentic(got); !result {
		t.Errorf("result is incorrect: %s\n  want : %v\n   got : %v\n", reason, *want, *got.(*Net))
	}
}

func TestCreateState(t *testing.T) {
	want := expState01
	got := NewState(expNet, "testing-state-01", "Testing State #01", "description of Testing State #01", 1000, 10)
	if result, reason := want.IsIdentic(got); !result {
		t.Errorf("result is incorrect: %s\n  want : %v\n   got : %v\n", reason, *want, *got.(*State))
	}
}

func TestCreateTransition(t *testing.T) {
	want := expTransition01
	got := NewTransition(expNet, "testing-transition-01", "Testing Transition #01", "description of Testing Transition #01", nil)
	if result, reason := want.IsIdentic(got); !result {
		t.Errorf("result is incorrect: %s\n  want : %v\n   got : %v\n", reason, *want, *got.(*Transition))
	}
}

func TestConnectStateTransition(t *testing.T) {
	initConnetStateTransition()
	n := NewNet("testing-net", "Testing Net", "description of Testing Net")
	s1 := NewState(n, "testing-state-01", "Testing State #01", "description of Testing State #01", 1000, 10)
	s2 := NewState(n, "testing-state-02", "Testing State #02", "description of Testing State #02", 1001, 10)
	t1 := NewTransition(n, "testing-transition-01", "Testing Transition #01", "description of Testing Transition #01", nil)
	t2 := NewTransition(n, "testing-transition-02", "Testing Transition #02", "description of Testing Transition #02", nil)
	n.ConnectStateTransition(s1, t1, ElementTypeArcNormal, 1)
	n.ConnectStateTransition(s2, t1, ElementTypeArcNormal, 1)
	n.ConnectStateTransition(s1, t2, ElementTypeArcNormal, 1)
	n.ConnectStateTransition(s2, t2, ElementTypeArcInhibit, 1)

	if result, reason := expNet.IsIdentic(n); !result {
		t.Errorf("result is incorrect: %s\n  want : %v\n   got : %v\n", reason, *expNet, *n.(*Net))
	}

	if result, reason := expState01.IsIdentic(s1); !result {
		t.Errorf("result is incorrect: %s\n  want : %v\n   got : %v\n", reason, *expState01, *s1.(*State))
	}

	if result, reason := expState02.IsIdentic(s2); !result {
		t.Errorf("result is incorrect: %s\n  want : %v\n   got : %v\n", reason, *expState02, *s2.(*State))
	}

	if result, reason := expTransition01.IsIdentic(t1); !result {
		t.Errorf("result is incorrect: %s\n  want : %v\n   got : %v\n", reason, *expTransition01, *t1.(*Transition))
	}

	if result, reason := expTransition02.IsIdentic(t2); !result {
		t.Errorf("result is incorrect: %s\n  want : %v\n   got : %v\n", reason, *expTransition02, *t2.(*Transition))
	}
}

func TestConnectTransitionState(t *testing.T) {
	initConnetTransitionState()
	n := NewNet("testing-net", "Testing Net", "description of Testing Net")
	s3 := NewState(n, "testing-state-03", "Testing State #03", "description of Testing State #03", 1002, 10)
	t1 := NewTransition(n, "testing-transition-01", "Testing Transition #01", "description of Testing Transition #01", nil)
	t2 := NewTransition(n, "testing-transition-02", "Testing Transition #02", "description of Testing Transition #02", nil)
	n.ConnectTransitionState(s3, t1)
	n.ConnectTransitionState(s3, t2)

	if result, reason := expNet.IsIdentic(n); !result {
		t.Errorf("result is incorrect: %s\n  want : %v\n   got : %v\n", reason, *expNet, *n.(*Net))
	}

	if result, reason := expState03.IsIdentic(s3); !result {
		t.Errorf("result is incorrect: %s\n  want : %v\n   got : %v\n", reason, *expState01, *s3.(*State))
	}

	if result, reason := expTransition01.IsIdentic(t1); !result {
		t.Errorf("result is incorrect: %s\n  want : %v\n   got : %v\n", reason, *expTransition01, *t1.(*Transition))
	}

	if result, reason := expTransition02.IsIdentic(t2); !result {
		t.Errorf("result is incorrect: %s\n  want : %v\n   got : %v\n", reason, *expTransition02, *t2.(*Transition))
	}
}
