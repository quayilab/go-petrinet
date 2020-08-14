package gopetrinet

// IElement :
type IElement interface {
	// internal attributes
	ID(values ...string) string
	Label(values ...string) string
	Desc(values ...string) string
	Type(values ...int) int
	Net(values ...INet) INet
	Enabled(values ...bool) bool
	// comparison
	IdenticWith(element IElement) (bool, string)
}

// INode :
type INode interface {
	// internal attributes
	NodeType() int
	Interior(values ...INet) INet
	// inputs
	Inputs(values ...int) []INode
	InputCount() int
	InputAdd(values ...INode)
	InputRemove(values ...INode)
	InputClear()
	// output
	Outputs(values ...INode) []INode
	OutputCount() int
	OutputAdd(values ...INode)
	OutputRemove(values ...INode)
	OutputClear()
	// comparison
	IdenticWith(node INode) (bool, string)
}

// IState :
type IState interface {
	// internal attributes
	Capacity(values ...int) int
	StateID(values ...int) int
	// token
	Token(index int) []IToken
	TokenCount() int
	TokenCopies() []IToken
	TokenAdd(values ...IToken)
	// comparison
	IdenticWith(state IState) (bool, string)
	// operations
	Ready(transition ITransition) bool
}

// ITransition :
type ITransition interface {
	// comparison
	IdenticWith(ITransition) (bool, string)
	// net configuration
	ConnectInput(IState)
	ConnectOutput(IState)
	// operations
	StateReady(IState) bool
	Ready() bool
	Execute()
	DistributeTokens()
}

// IToken :
type IToken interface {
	Data(...interface{}) interface{}
}

// INet :
type INet interface {
	// nodes
	NodeCount() int
	Nodes() []INode
	Node(int) INode
	// states
	StateCount() int
	States() []IState
	State(int) IState
	StateAdd(string, string, string, int, int) IState
	StateAddToken(IState, ...IToken)
	// transitions
	TransitionCount() int
	Transitions() []ITransition
	Transition(int) ITransition
	TransitionAdd(string, string, string, OnExec) ITransition
	TransitionReady(ITransition) bool
	// net configuration
	ConnectStateTransition(IState, ITransition) error
	ConnectTransitionState(IState, ITransition) error
	ConnectNodes(node1, node2 INode) error
	TryConnectNodes(node1, node2 INode) error
	DisconnectNodes(node1, node2 INode) error
	TryDisconnectNodes(node1, node2 INode) error

	// comparison
	IdenticWith(s INet) (bool, string)
	// operation
	Run()
	Step()
	Pause()
	Running() bool
	Ready() bool
	DeadLock() bool
	Execute(ITransition) (bool, error)
}
