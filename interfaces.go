package gopetrinet

// IElement :
type IElement interface {
	// internal attributes
	ID(values ...string) (result string)
	Label(values ...string) (result string)
	Desc(values ...string) (result string)
	Type(values ...int) (result int)
	Net(values ...INet) (result INet)
	Enabled(values ...bool) (result bool)
	// comparison
	IdenticWith(element IElement) (identic bool, reason string)
}

// INode :
type INode interface {
	// internal attributes
	NodeType() (result int)
	Interior(values ...INet) (result INet)
	// inputs
	Inputs(values ...int) (result []INode)
	InputCount() int
	InputAdd(values ...INode) (err error)
	InputRemove(values ...INode) (err error)
	InputClear()
	// output
	Outputs(values ...INode) (result []INode)
	OutputCount() int
	OutputAdd(values ...INode) (err error)
	OutputRemove(values ...INode) (err error)
	OutputClear()
	// comparison
	IdenticWith(node INode) (identic bool, reason string)
}

// IState :
type IState interface {
	// internal attributes
	Capacity(values ...int) (result int)
	StateID(values ...int) (result int)
	StorageMode(values ...int) (result int)
	TokenTypes(tokenTypes ...string) (result []string)
	// token
	//Token(indexes ...int) (result []IToken)
	TokenCount(tokenType ...string) int
	TokenPeek(count int) (result []IToken)
	TokenFetch(count int) (result []IToken)
	//TokenCopies() (result []IToken)
	TokenAdd(values ...IToken)
	// comparison
	IdenticWith(state IState) (identic bool, reason string)
	// operations
	Ready(transition ITransition) (result bool)
}

// ITransition :
type ITransition interface {
	// comparison
	IdenticWith(transition ITransition) (result bool, reason string)
	// net configuration
	ConnectInput(state IState)
	ConnectOutput(state IState)
	// operations
	ActivationTreshold(tokenType string) (result int)
	StateReady(state IState) (result bool)
	Ready() (result bool)
	Execute()
	DistributeTokens()
}

// IToken :
type IToken interface {
	Type() (result string)
	Data(data ...interface{}) (result interface{})
	Values(values ...string) (result []interface{})
}

// INet :
type INet interface {
	// nodes
	NodeCount() (result int)
	Nodes() (result []INode)
	Node(index int) (result INode)
	// states
	StateCount() (result int)
	States() (result []IState)
	State(index int) (result IState)
	StateAdd(id string, label string, description string, stateid int, capacity int) (result IState)
	StateAddToken(state IState, tokens ...IToken)
	// transitions
	TransitionCount() (result int)
	Transitions() (result []ITransition)
	Transition(index int) (result ITransition)
	TransitionAdd(id string, label string, description string, onExec OnExec, onTestState OnTestState) (result ITransition)
	TransitionReady(ITransition) bool
	// net configuration
	ConnectStateTransition(state IState, transition ITransition) (err error)
	ConnectTransitionState(transition ITransition, state IState) (err error)
	ConnectNodes(node1, node2 INode) (err error)
	TryConnectNodes(node1, node2 INode) (err error)
	DisconnectNodes(node1, node2 INode) (err error)
	TryDisconnectNodes(node1, node2 INode) (err error)

	// comparison
	IdenticWith(net INet) (identic bool, reason string)
	// operation
	Run()
	Step()
	Pause()
	Running() (result bool)
	Ready() (result bool)
	DeadLock() (result bool)
	Execute(transition ITransition) (result bool, err error)
}
