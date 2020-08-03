package gopetrinet

const (
	// ElementNet :
	ElementNet = iota
	// ElementNode :
	ElementNode
	// ElementArc :
	ElementArc
	// ElementToken :
	ElementToken
)

const (
	// NodeState :
	NodeState = iota
	// NodeTransition :
	NodeTransition
	// NodeOmni :
	NodeOmni
)

const (
	// ArcNormal :
	ArcNormal = iota
	// ArcInhibit :
	ArcInhibit
)

// ElementIntf :
type ElementIntf interface {
	GetID() string
	GetLabel() string
	GetDesc() string
	GetType() int
	GetNet() NetIntf
	GetStatus() bool
	Enable()
	Disable()
	IsIdentic(ElementIntf) (bool, string)
}

// NodeIntf :
type NodeIntf interface {
	GetInputs() []NodeIntf
	GetOutputs() []NodeIntf
	GetInputCount() int
	GetOutputCount() int
	AddInput(...NodeIntf)
	AddOutput(...NodeIntf)
	GetInterior() NetIntf
	SetInterior(NetIntf)
	GetNodeType() int
	IsIdentic(n NodeIntf) (bool, string)
}

// StateIntf :
type StateIntf interface {
	GetCapacity() int
	SetCapacity(int)
	GetStateID() int
	SetStateID(int)
	GetTokenCount() int
	GetToken(int) []TokenIntf
	GetTokenCopies() []TokenIntf
	AddToken(...TokenIntf)
	IsReady(TransitionIntf) bool
	IsIdentic(StateIntf) (bool, string)
}

// TransitionIntf :
type TransitionIntf interface {
	GetActivationTreshold(int) int
	IsReady() bool
	IsStateReady(StateIntf) bool
	Execute()
	DistributeTokens()
	IsIdentic(s TransitionIntf) (bool, string)
	ConnectInput(s StateIntf, arctype, activationTreshold int)
	ConnectOutput(s StateIntf)
}

// TokenIntf :
type TokenIntf interface {
	GetData() interface{}
	SetData(d interface{})
}

// NetIntf :
type NetIntf interface {
	GetNodeCount() int
	GetNodes() []NodeIntf
	GetStateCount() int
	GetStates() []StateIntf
	GetTransitionCount() int
	GetTransitions() []TransitionIntf
	GetNode(int) NodeIntf
	GetState(int) StateIntf
	GetTransition(int) TransitionIntf
	Run()
	Step()
	Pause()
	IsReady() bool
	IsDeadLock() bool
	IsIdentic(s NetIntf) (bool, string)
	IsTransitionReady(TransitionIntf) bool
	Execute(TransitionIntf) (bool, error)
	AddState(string, string, string, int, int) StateIntf
	AddTransition(string, string, string, OnExec) TransitionIntf
	ConnectStateTransition(StateIntf, TransitionIntf, int, int)
	ConnectTransitionState(StateIntf, TransitionIntf)
	AddTokenToState(StateIntf, ...TokenIntf)
}
