package gopetrinet

const (
	// ElementNet :
	ElementNet = iota
	// ElementNode :
	ElementNode
	// ElementArc :
	ElementArc
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
	// ArcBidirectional :
	ArcBidirectional
)

const (
	// StorageChanel :
	StorageChanel = iota
	// StorageStack :
	StorageStack
	// StorageMultiset :
	StorageMultiset
)

const (
	// EventStateCreate :
	EventStateCreate = `creating state "%s"`
	// EventStateAdd :
	EventStateAdd = `adding state "%s"`
	// EventStateReadyTest :
	EventStateReadyTest = `testing state "%s"`
	// EventTransitionCreate :
	EventTransitionCreate = `creating transition "%s"`
	// EventTransitionAdd :
	EventTransitionAdd = `adding transition "%s"`
	// EventTransitionReadyTest :
	EventTransitionReadyTest = `testing transition "%s"`
	// EventTokenCreate :
	EventTokenCreate = `creating token "%s"`
	// EventTokenAdd :
	EventTokenAdd = `adding token "%s" to state "%"`
	// EventTokenDataAssign :
	EventTokenDataAssign = `assigning data to token "%s"`
	// EventConnectNodes :
	EventConnectNodes = `connecting nodes "%s" to "%s"`
	// EventDisconnectNodes :
	EventDisconnectNodes = `disconnecting node "%s" from "%s"`
	// EventConnectStateTransition :
	EventConnectStateTransition = `connecting state "%s" to transition "%s"`
	// EventConnectTransitionState :
	EventConnectTransitionState = `connecting transition "%s" to state "%s"`

	// ResultSucceeded :
	ResultSucceeded = "succeeded"
	// ResultFailed :
	ResultFailed = "failed"
	// ResultValue :
	ResultValue = "finished with result: %v"

	// ReasonObjectExists :
	ReasonObjectExists = `object "%s" exists`
	// ReasonConnectionExists :
	ReasonConnectionExists = `connection between node "%s" and "%s" exists`
	// ReasonConnectionInexist :
	ReasonConnectionInexist = `connection between node "%s" and "%s" not exist`
	// ReasonIncompatibleNodeTypes :
	ReasonIncompatibleNodeTypes = `incompatible node types of nodes "%s" and "%s"`
	// ReasonStateNotReady :
	ReasonStateNotReady = `state "%s" not ready`
	// ReasonTransitionNotReady :
	ReasonTransitionNotReady = `transition "%s" not ready`
	// ReasonNetDeadlock :
	ReasonNetDeadlock = `net "%s" deadlock`
	// ReasonNetNotRunning :
	ReasonNetNotRunning = `net "%s" not running`
	// ReasonNetRunning :
	ReasonNetRunning = `net "%s" running`
)

// ElementTypeStr :
var ElementTypeStr = map[int]string{
	ElementNet:  "net",
	ElementNode: "node",
	ElementArc:  "arc",
}

// NodeTypeStr :
var NodeTypeStr = map[int]string{
	NodeState:      "node-state",
	NodeTransition: "node-transition",
	NodeOmni:       "node-omni",
}

// ArcTypeStr :
var ArcTypeStr = map[int]string{
	ArcNormal:        "arc-normal",
	ArcInhibit:       "arc-inhibitor",
	ArcBidirectional: "arc-bidirectional",
}

// NodeIndex :
type NodeIndex = int

func elementExists(id string, collections []IElement) (result bool) {
	for _, e := range collections {
		if e.(IElement).ID() == id {
			result = true
			return
		}
	}
	return
}

func nodeExists(id string, collections []INode) (result bool) {
	for _, e := range collections {
		if e.(IElement).ID() == id {
			result = true
			return
		}
	}
	return
}
